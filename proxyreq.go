package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func ProxyReq(subdomain string, client *http.Client, ch chan struct{}, wg *sync.WaitGroup, p *string) {

	defer func() { <-ch }()
	defer wg.Done()

	var url string
	if *p != "" {
		path := strings.TrimPrefix(*p, "/")
		url = fmt.Sprintf("%s/%s", subdomain, path)
	} else {
		url = fmt.Sprintf("%s/", subdomain)
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[!] The subdomain %s has failed\n", subdomain)
		return
	}
	defer resp.Body.Close()

}
