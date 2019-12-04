package main

import (
	"io/ioutil" // Package ioutil implements some I/O utility functions.
	"log"
	"net/http" // Package http provides HTTP client and server implementations.
)

func main() {
	log.Println(GetRequest())
}

// GetRequest makes a simple get requests
func GetRequest() (string, error) {
	resp, err := http.Get("https://k-blog0130.herokuapp.com/api/v2/posts")
	// Error handling
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// Converty byte[] to string
	return string(body), nil
}
