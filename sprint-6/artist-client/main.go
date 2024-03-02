package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	_, err := http.Post(`http://localhost:8080`, `application/json`,
		bytes.NewBufferString(`{"ID": 10, "name": "Garbage", "Genre":"rock", "SongS":["Only Happy When It Rains", "Stupid Girl", "Push It"]}`))
	if err != nil {
		err = fmt.Errorf("failed post json: %w", err)
		log.Println(err)
		return
	}
}
