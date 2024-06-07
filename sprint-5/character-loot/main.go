package main

import (
	"log"
)

// Sword описывает меч.
type Sword struct {
	Power int
}

// Scroll описывает свиток.
type Scroll struct {
	Magic int
}

// Loot опишите интерфейсный тип Loot.
type Loot interface {
	Apply()
}

// Apply добавьте нужный метод для типов Sword и Scroll.
func (s Sword) Apply() {
	log.Printf("Меч %d\n", s.Power)
}

func (s Scroll) Apply() {
	log.Printf("Свиток %d\n", s.Magic)
}

func main() {
	// Loot - это слайс интерфейсного типа Loot. Так как типы Sword и Scroll должны удовлетворять этому интерфейсу, то можно использовать эти структуры
	// как элементы слайсы. Этот слайс создан исключительно для проверки того, правильно ли реализован тип Loot и метод Apply() для структур
	loot := []Loot{Sword{Power: 50}, Scroll{Magic: 20}, Scroll{Magic: 70}} //nolint:mnd // it's learning code

	for _, v := range loot {
		v.Apply()
	}
}
