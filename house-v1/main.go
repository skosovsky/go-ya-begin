package main

import "fmt"

type Thing struct {
	Name   string
	Weight int
}

type Room struct {
	Name   string
	Things []Thing
}

type House struct {
	Name  string
	Rooms [][]Room
}

func main() {
	house := House{
		Name: "Дом v1",
		Rooms: [][]Room{
			{
				{Name: "Кладовка", Things: []Thing{
					{Name: "топор", Weight: 3000},
					{Name: "фонарик", Weight: 0},
					{Name: "брелок", Weight: 0},
				}},
				{Name: "Котельная", Things: []Thing{
					{Name: "верёвка", Weight: 200},
					{Name: "рюкзак", Weight: 500},
				}}},
			{
				{Name: "Столовая", Things: []Thing{
					{Name: "карандаш", Weight: 0},
					{Name: "кольцо", Weight: 0},
					{Name: "карта", Weight: 0},
					{Name: "бинт", Weight: 0},
				}}},
		},
	}
	fmt.Println(house)
}
