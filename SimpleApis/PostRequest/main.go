package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm("http://localhost:3000", values)

	if err != nil {
		// Can't make the request
		log.Fatal(err)
	}
	log.Printf("Status: %v", resp.Status)

}
