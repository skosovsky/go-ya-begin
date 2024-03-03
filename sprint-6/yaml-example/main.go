package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Artist содержит данные об артисте
type Artist struct {
	ID    int      `yaml:"id"`
	Name  string   `yaml:"name"`
	Genre string   `yaml:"genre"`
	Songs []string `yaml:"songs"`
}

func main() {
	artist := Artist{
		ID:    1,
		Name:  "30 seconds to Mars",
		Genre: "rock",
		Songs: []string{
			`The Kill`,
			`A Beautiful Lie`,
			`Attack`,
			`Live Like A Dream`,
		},
	}

	yamlData, err := yaml.Marshal(&artist)
	if err != nil {
		err = fmt.Errorf("filed to marshal yaml: %w", err)
		log.Println(err)
		return
	}

	fmt.Println(yamlData)
	fmt.Println(string(yamlData))

	file, err := os.Create("artist.yaml")
	if err != nil {
		err = fmt.Errorf("filed to create file: %w", err)
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(yamlData)
	if err != nil {
		err = fmt.Errorf("filed to write file: %w", err)
		log.Println(err)
		return
	}

}
