package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")

	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>２回目以降</body></html>")
	} else {
		fmt.Fprintf(w, "<html><body>初訪問</body></html>")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)

	log.Println("start http listening :8888")
	httpServer.Addr = ":8888"
	log.Println(httpServer.ListenAndServe())

}
