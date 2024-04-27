package main

import (
	"encoding/json"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}

func main() {
	users := []User{
		{
			ID:       2, //nolint:gomnd // it's learning code
			Name:     "Гофер",
			Email:    "gopher@gophermate.com",
			Password: "I4mG0ph3R",
		},
		{
			ID:       1,
			Name:     "Алиса",
			Email:    "",
			Password: "4L1c3iAnD3x",
		},
		{
			ID:       3, //nolint:gomnd // it's learning code
			Name:     "",
			Email:    "rustocean@rust.org",
			Password: "Rust0C34n1T$m3",
		},
	}

	out, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		log.Printf("ошибка при сериализации в json: %s\n", err.Error())
		return
	}
	log.Println(string(out))
}
