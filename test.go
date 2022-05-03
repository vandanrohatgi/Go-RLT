package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var requests int = 1000

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Kindly provide the URL as input")
		fmt.Println("Usage: go run test.go https://google.com")
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Printf("Checking availability of %s\n", url)
	_, err := make_request(url)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Starting rate limit test...")
		fmt.Printf("Sending %d requests to %s\n", requests, url)
	}
	for i := 0; i < requests; i++ {
		make_request(url)
	}
}

// make basic GET request
func make_request(url string) (http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return http.Response{}, err

	} else {
		return *response, nil
	}
}
