package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func printAnswer(response *http.Response) {
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Println("Ошибка:", err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	response, err := http.Get("http://localhost:8080/")
	if err != nil {
		log.Println("Ошибка", err)
	}
	printAnswer(response)

	response, err = http.PostForm("http://localhost:8080/", url.Values{
		"email": {"skosovsky@gmail.com"},
		"name":  {"Sergey"},
	})
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	printAnswer(response)

	response, err = http.PostForm("http://localhost:8080/second/", url.Values{})
	if err != nil {
		log.Println("Ошибка", err)
		return
	}

	fmt.Println("Код статуса:", response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Ошибка", err)
		return
	}
	response.Body.Close()

	fmt.Println(string(body))
}
