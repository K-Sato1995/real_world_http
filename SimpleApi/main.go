package main

import (
	"io/ioutil" // Package ioutil implements some I/O utility functions.
	"log"       // Package log implements a simple logging package
	"net/http"  // Package http provides HTTP client and server implementations.
)

func main() {
	resp, _ := GetRequest()

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// Log status
	log.Println(resp.Status)
	// Log an object of http.Response
	log.Println(resp)
	// Converty byte[] to string
	log.Println(string(body))
}

// GetRequest makes a simple get requests
func GetRequest() (*http.Response, error) {
	resp, err := http.Get("https://k-blog0130.herokuapp.com/api/v2/posts")
	// Error handling
	if err != nil {
		return nil, err
	}

	return resp, nil
}
