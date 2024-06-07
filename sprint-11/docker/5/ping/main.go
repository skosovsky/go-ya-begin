package main

import (
	"fmt"
	"io"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	res, err := http.Get("http://counter:8080/count")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	_, err = io.Copy(w, res.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/ping", ping)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
