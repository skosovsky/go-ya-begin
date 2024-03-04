package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Artist struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Born  string   `json:"born"`
	Genre string   `json:"genre"`
	Songs []string `json:"songs"`
}

var artists = map[string]Artist{ //nolint:gochecknoglobals //it's lesson
	"1": {
		ID:    "1",
		Name:  "30 Seconds To Mars",
		Born:  "1998",
		Genre: "alternative",
		Songs: []string{
			"The Kill",
			"A Beautiful Lie",
			"Attack",
			"Live Like A Dream",
		},
	},
	"2": {
		ID:    "2",
		Name:  "Garbage",
		Born:  "1994",
		Genre: "alternative",
		Songs: []string{
			"Queer",
			"Shut Your Mouth",
			"Cup of Coffee",
			"Til the Day I Die",
		},
	},
}

func getArtists(w http.ResponseWriter, _ *http.Request) {
	response, err := json.Marshal(artists)
	if err != nil {
		err = fmt.Errorf("filed to marshal response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		err = fmt.Errorf("filed to write response: %w", err)
		log.Println(err)
	}
}

func postArtist(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		err = fmt.Errorf("filed to read buffer: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &artist)
	if err != nil {
		err = fmt.Errorf("filed to unmarshal: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	artists[artist.ID] = artist

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	artist, ok := artists[id]
	if !ok {
		err := errors.New("artist not found")
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	response, err := json.Marshal(artist)
	if err != nil {
		err = fmt.Errorf("filed to marshal response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		err = fmt.Errorf("filed to write response: %w", err)
		log.Println(err)
	}
}

func main() {
	router := chi.NewRouter()

	router.Get("/artists", getArtists)
	router.Post("/artists", postArtist)
	router.Get("/artist/{id}", getArtist)

	if err := http.ListenAndServe(":8080", router); err != nil { //nolint:gosec //it's lesson
		err = fmt.Errorf("filed to runing server: %w", err)
		log.Println(err)
		return
	}
}
