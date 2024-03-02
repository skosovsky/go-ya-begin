package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func mainHandle(w http.ResponseWriter, r *http.Request) {
	var out string

	switch r.URL.Path {
	case "/time", "/time/":
		out = time.Now().Format("02.01.2006 15:04:05")
	default:
		out = fmt.Sprintf("Host %s\nPath: %s\nMethod: %s",
			r.Host, r.URL.Path, r.Method)
	}

	_, err := w.Write([]byte(out))
	if err != nil {
		return
	}
	log.Println("Получен запрос")
}

func main() {
	log.Println("Запускаем сервер")
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	log.Println("Завершаем сервер")
}
