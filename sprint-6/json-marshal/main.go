package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}

func main() {
	users := []User{
		{
			ID:       2, //nolint:mnd
			Name:     "Гофер",
			Email:    "gopher@gophermate.com",
			Password: "I4mG0ph3R",
		},
		{ //nolint:exhaustruct
			ID:       1,
			Name:     "Алиса",
			Password: "4L1c3iAnD3x",
		},
		{ //nolint:exhaustruct
			ID:       3, //nolint:mnd
			Email:    "rustocean@rust.org",
			Password: "Rust0C34n1T$m3",
		},
	}

	out, err := json.Marshal(users)
	if err != nil {
		fmt.Printf("ошибка при сериализации в json: %s\n", err.Error())
		return
	}
	fmt.Println(string(out))
}
