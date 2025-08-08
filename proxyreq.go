package main

import (
	"fmt"
	"net/http"
	"sync"
)

func ProxyReq(subdomain string, client *http.Client, ch chan struct{}, wg *sync.WaitGroup) {

	defer func() { <-ch }()
	defer wg.Done()

	url := fmt.Sprintf("%s/", subdomain)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[!] The subdomain %s has failed\n", subdomain)
		return
	}
	defer resp.Body.Close()

}
