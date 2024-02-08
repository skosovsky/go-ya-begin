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

type Floor struct {
	Name  string
	Rooms []Room
}

type House struct {
	Name   string
	Floors []Floor
}

func main() {
	var pen Thing
	pen.Name = "карандаш"
	pen.Weight = 50

	var roomLib Room
	roomLib.Name = "Библиотека"
	roomLib.Things = append(roomLib.Things, pen)
	fmt.Println(roomLib)

	room := Room{
		Name:   "Столовая",
		Things: []Thing{{Name: "карандаш", Weight: 50}, {Name: "кольцо", Weight: 30}},
	}
	fmt.Println(room)

	list := []struct {
		value int
		want  int
	}{
		{value: 10, want: 23},
		{value: 11, want: 25},
		{value: 5, want: 13},
	}
	fmt.Println(list)

	house := House{Name: "Дом v1", Floors: []Floor{{
		Name: "Подвал",
		Rooms: []Room{{Name: "Кладовка", Things: []Thing{
			{Name: "топор", Weight: 3000},
			{Name: "фонарик", Weight: 0},
			{Name: "брелок", Weight: 3000},
		}}, {Name: "Котельная", Things: []Thing{
			{Name: "верёвка", Weight: 200},
			{Name: "рюкзак", Weight: 500},
		}}},
	}, {
		Name: "1 этаж",
		Rooms: []Room{
			{Name: "Столовая", Things: []Thing{
				{Name: "карандаш", Weight: 0},
				{Name: "кольцо", Weight: 0},
				{Name: "карта", Weight: 0},
				{Name: "бинт", Weight: 0},
			}},
		},
	}}}
	fmt.Println(house)
}
