package main

import (
	"fmt"
	"net/http"
	"os"
)

func ping(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("ENV1 =", os.Getenv("ENV1"))
	fmt.Println("ENV2 =", os.Getenv("ENV2"))
}

func main() {
	http.HandleFunc("/ping", ping)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
