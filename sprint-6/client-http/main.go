package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", http.NoBody)
	if err != nil {
		log.Println("Ошибка формирования запроса:", err)
		return
	}

	request.Header.Set("Custom-Header", "John Doe")
	request.Header.Add("Custom-Header", "1234")

	response, err := client.Do(request)
	if err != nil {
		log.Println("Ошибка отправки запроса:", err)
		return
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Println("Ошибка закрытия:", err)
			return
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println(string(body))
}
