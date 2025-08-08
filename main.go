package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {

	targetFile := flag.String("l", "", "Subdomains file from your target")
	proxy := flag.String("proxy", "http://127.0.0.1:8080", "Proxy address")
	concurrent := flag.Int("t", 20, "Amount of threads")
	followRedirect := flag.Bool("redirect", false, "Follow the redirects and logs the last request")
	flag.Parse()

	if *targetFile == "" {
		fmt.Println("[!] Enter a subdomains file")
		os.Exit(1)
	}

	client := HttpClient(proxy, followRedirect)

	subdomainsFile, err := os.Open(*targetFile)
	if err != nil {
		fmt.Println("[!] Error reading the subdomains file")
	}

	ch := make(chan struct{}, *concurrent)
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(subdomainsFile)

	for scanner.Scan() {
		subdomain := strings.TrimSpace(scanner.Text())
		subdomain = strings.TrimSuffix(subdomain, "/")
		if !strings.HasPrefix(subdomain, "http://") && !strings.HasPrefix(subdomain, "https://") {
			fmt.Printf("[!] The subdomain %s is invalid. Subdomains must start with http:// or https://\n", subdomain)
			os.Exit(1)
		}

		ch <- struct{}{}
		wg.Add(1)
		go ProxyReq(subdomain, client, ch, &wg)
	}

	wg.Wait()

}
