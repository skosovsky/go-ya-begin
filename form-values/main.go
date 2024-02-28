package main

import (
	"fmt"
	"net/url"
)

func FormValues(achievement string) url.Values {
	// переделайте код функции с использованием переменной типа url.Values
	// и методов Set() и Add()

	result := url.Values{}

	result.Set("name", "Вася")
	result.Add("nick", "superstar")
	result.Add("achieves", "cool")
	result.Add("achieves", "best")
	result.Add("achieves", achievement)

	return result
}

func main() {
	vals := FormValues("80 level")

	fmt.Println(vals["name"])
	fmt.Println(vals["nick"])
	fmt.Println(vals["achieves"])
}
