package main

import "fmt"

// Property описывает общие свойства предметов.
type Property struct {
	Name  string // название предмета
	Price int
}

// Loot содержит методы управления инвентарём.
type Loot interface {
	Apply()
	Properties() Property
	Sell() int
}

type Scroll struct {
	Property
}

type Sword struct {
	Property
}

// добавьте нужные методы для типов Scroll и Sword
func (s Scroll) Apply() {

}

func (s Scroll) Properties() Property {
	return s.Property
}

func (s Scroll) Sell() int {
	return s.Price
}

func (s Sword) Apply() {

}

func (s Sword) Properties() Property {
	return s.Property
}

func (s Sword) Sell() int {
	return s.Price
}

func main() {
	loot := []Loot{
		Scroll{Property{Name: "Свиток знаний"}},
		Sword{Property{Name: "Двуручный меч"}},
	}

	for _, v := range loot {
		fmt.Println(v.Properties())
	}
	fmt.Println("Успех!")
}
