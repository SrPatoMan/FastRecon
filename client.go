package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func HttpClient(proxy *string, redirect *bool) *http.Client {

	transport := &http.Transport{
		MaxIdleConns:          200,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		Proxy: func(r *http.Request) (*url.URL, error) {
			proxyURL, err := url.Parse(*proxy)
			if err != nil {
				fmt.Println("[!] Bad proxy defined")
				os.Exit(1)
			}

			return proxyURL, nil
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	if !*redirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return client
}
