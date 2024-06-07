package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Artist описывает музыкальную группу.
type Artist struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Genre string   `json:"genre"`
	Songs []string `json:"songs"`
}

// Переменная artists содержит пока один музыкальный коллектив.
var artists = map[string]Artist{ //nolint:gochecknoglobals // it's learning code
	"30 seconds to Mars": {
		ID:    1,
		Name:  "30 seconds to Mars",
		Genre: "rock",
		Songs: []string{
			"The Kill",
			"A Beautiful Lie",
			"Attack",
			"Live Like A Dream",
		},
	},
}

// JSONHandler принимает значение из параметра id, ищет по нему в map группу, конвертирует
// данные из переменной artists в JSON и выводит их в браузере.
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	band := r.URL.Query().Get("id")
	response, err := json.Marshal(artists[band])
	if err != nil {
		err = fmt.Errorf("filed to marshal response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		log.Println(fmt.Errorf("filed to write response: %w", err))
		return
	}
}

func main() {
	http.HandleFunc("/", JSONHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil { //nolint:gosec // it's learning code
		err = fmt.Errorf("failed running server: %w", err)
		log.Println(err)
		return
	}
}
