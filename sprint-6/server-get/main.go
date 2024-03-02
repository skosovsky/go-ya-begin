package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainHandle(w http.ResponseWriter, r *http.Request) {
	var answer string

	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		answer = "Укажите имя заголовка в параметре: http://localhost:8080/?name=User-Agent"
	} else if v := r.Header.Get(name); len(v) > 0 {
		answer = fmt.Sprintf("%s: %s", name, v)
	} else {
		answer = fmt.Sprintf("Заголовок %s не определен", name)
	}
	_, err := io.WriteString(w, answer)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
