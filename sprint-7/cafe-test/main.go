package main

import (
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{ //nolint:gochecknoglobals // it's learning code
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("count missing"))
		if err != nil {
			return
		}
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("wrong count value"))
		if err != nil {
			return
		}
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("wrong city value"))
		if err != nil {
			return
		}
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(answer))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc(`/cafe`, mainHandle)
	err := http.ListenAndServe(":8080", nil) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
