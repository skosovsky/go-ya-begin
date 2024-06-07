package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const filePath = "counter.txt"

func count(w http.ResponseWriter, _ *http.Request) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(content) == "" {
		content = append(content, '0')
	}

	num, _ := strconv.Atoi(string(content))
	num++

	err = os.WriteFile(filePath, []byte(strconv.Itoa(num)), 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("num is %d\n", num)
	if _, err := fmt.Fprintf(w, "%d\n", num); err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/count", count)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
