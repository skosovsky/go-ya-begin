package main

import (
	"bytes"
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
var artists = map[string]Artist{
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

// JSONHandler принимает значение из параметра id, ищет по нему в мапе группу, конвертирует
// данные из переменной band в JSON и выводит их в браузере.
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	var band string

	if r.Method == http.MethodPost {
		var artist Artist
		var buf bytes.Buffer

		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			err = fmt.Errorf("filed to read from body: %w", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(buf.Bytes(), &artist); err != nil {
			err = fmt.Errorf("filed to unmarshal: %w", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		artists[band] = artist
	}

	if r.Method != http.MethodPost {
		band = r.URL.Query().Get("band")
	}

	response, err := json.MarshalIndent(artists[band], " ", "  ")
	if err != nil {
		err = fmt.Errorf("filed to marshal: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		err = fmt.Errorf("filed to write resonse: %w", err)
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", JSONHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		err = fmt.Errorf("failed running server: %w", err)
		log.Println(err)
		return
	}
}
