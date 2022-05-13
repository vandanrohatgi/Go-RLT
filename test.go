package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func make_request(url string, request_type string) (http.Response, error) {
	var response *http.Response
	var err error
	if request_type == "GET" {
		response, err = http.Get(url)
	} else if request_type == "POST" {
		body, _ := json.Marshal(map[string]string{"test": "test"})
		response, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	}
	if err != nil {
		return http.Response{}, err

	} else {
		return *response, nil
	}
}

func main() {
	url := flag.String("url", "", "URL to test")
	requests := flag.Int("requests", 1000, "Number of requests to send")
	request_type := flag.String("request-type", "GET", "Type of request to send")
	flag.Parse()
	fmt.Println(*url, *request_type, *requests)
	if len(*url) == 0 {
		fmt.Println("URL is required!")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Checking availability of %s\n", *url)
	_, err := make_request(*url, *request_type)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Starting rate limit test...")
		fmt.Printf("Sending %d requests to %s\n", *requests, *url)
	}
	var wg sync.WaitGroup
	for i := 0; i < *requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			make_request(*url, *request_type)
		}()
	}
	wg.Wait()
}
