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
	flag.Parse()

	if *targetFile == "" {
		fmt.Println("[!] Enter a subdomains file")
		os.Exit(1)
	}

	client := HttpClient(proxy)

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
		ch <- struct{}{}
		wg.Add(1)
		// A continuar desde aqui
	}

	wg.Wait()

}
