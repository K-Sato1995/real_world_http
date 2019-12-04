package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url" // Package url parses URLs and implements query escaping.

	"../utilities"
)

func main() {
	// Create a query
	values := url.Values{
		"query": {"Hello World"},
	}
	resp, err := http.Get(utilities.EndPoint() + "?" + values.Encode())

	// Error Handling
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(body)

}
