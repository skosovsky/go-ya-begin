package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Artist содержит данные об артисте.
type Artist struct {
	ID    int      `yaml:"id"`
	Name  string   `yaml:"name"`
	Genre string   `yaml:"genre"`
	Songs []string `yaml:"songs"`
}

func main() {
	yamlFile, err := os.ReadFile("artist.yaml")
	if err != nil {
		err = fmt.Errorf("filed to read file: %w", err)
		log.Println(err)

		return
	}

	var artist Artist

	err = yaml.Unmarshal(yamlFile, &artist)
	if err != nil {
		err = fmt.Errorf("filed to unmarshal: %w", err)
		log.Println(err)

		return
	}

	log.Println("ID:", artist.ID)
	log.Println("Name:", artist.Name)
	log.Println("Genre:", artist.Genre)
	log.Printf("Songs: %v\n", artist.Songs)
}
