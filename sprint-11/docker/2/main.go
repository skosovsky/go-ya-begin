package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Hello")
	if _, err := fmt.Fprintln(w, "Hello, world!"); err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
