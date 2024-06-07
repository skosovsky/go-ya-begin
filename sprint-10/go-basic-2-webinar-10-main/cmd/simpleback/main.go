package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Requested path:", req.URL.Path)
		data, _ := httputil.DumpRequest(req, true)
		io.WriteString(w, "Hello, world!\n")
		w.Write(data)
	}

	http.HandleFunc("/", helloHandler)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
