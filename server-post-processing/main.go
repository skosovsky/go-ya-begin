package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		_, err := fmt.Fprintf(w, "Email: %s\nName: %s",
			r.PostFormValue("email"), r.PostFormValue("name"))
		if err != nil {
			return
		}
		return
	}
	_, err := io.WriteString(w, "Отправьте POST запрос с параметрами email и name")
	if err != nil {
		return
	}
}

func secondHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != http.MethodGet {
		//w.WriteHeader(http.StatusMethodNotAllowed)
		//fmt.Fprintf(w, pattern, "Сервер поддерживает только GET-запросы")
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", r.Method), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, pattern, "Получен GET-запрос")
}

func exampleHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Сервер не поддерживает "+r.Method,
			http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(r.Method)
	_, err := w.Write([]byte(r.FormValue("name")))
	if err != nil {
		return
	}
}

const pattern = `<!DOCTYPE html>
  <html lang="ru"><head>
  <meta charset="utf-8" />
  <title>Тестовый сервер</title>
  </head>
<body>%s</body></html>`

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainHandle)
	mux.HandleFunc("/second/", secondHandle)
	mux.HandleFunc("/example/", exampleHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
