package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	// Create a variable that stores a cookie
	jar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
	}
	// Create a http client that can store a coockie
	client := http.Client{
		Jar: jar,
	}

	// The client receive a coockie on the first access and send the coockie back to the server from the second access.
	for i := 0; i < 2; i++ {
		resp, err := client.Get("http://localhost:3000/coockie")
		if err != nil {
			log.Fatal(err)
		}
		dump, err := httputil.DumpResponse(resp, true)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(dump))
	}

}
