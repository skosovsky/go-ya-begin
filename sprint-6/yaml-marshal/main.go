package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
)

type User struct {
	ID       int    `yaml:"id"`
	Name     string `yaml:"name,omitempty"`
	Email    string `yaml:"email,omitempty"`
	Password string `yaml:"-"`
}

func main() {
	users := []User{
		{
			ID:       2,
			Name:     "Гофер",
			Email:    "gopher@gophermate.com",
			Password: "I4mG0ph3R",
		},
		{
			ID:       1,
			Name:     "Алиса",
			Password: "4L1c3iAnD3x",
		},
		{
			ID:       3,
			Email:    "rustocean@rust.org",
			Password: "Rust0C34n1T$m3",
		},
	}

	out, err := yaml.Marshal(&users)
	if err != nil {
		err = fmt.Errorf("filed to marshal yaml: %w", err)
		log.Println(err)
		return
	}

	fmt.Println(string(out))
}
