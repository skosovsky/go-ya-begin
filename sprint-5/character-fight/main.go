package main

import (
	"fmt"
)

// Hero описывает главного героя игры.
type Hero struct {
	Name   string
	Health int
	Damage int
	Def    int
}

// Создаём метод атаки.
func (h *Hero) Attack() {
	fmt.Printf("%s нанес урон %d\n", h.Name, h.Damage)
}

// Создаём метод защиты.
func (h *Hero) Defense() {
	fmt.Printf("%s заблокировал %d единиц урона\n", h.Name, h.Def)
}

// Создаём специальный метод.
func (h *Hero) Special(healthpoints int) {
	fmt.Printf("Количество здоровья было %d\n", h.Health)
	h.Health += healthpoints
	fmt.Printf("Количество здоровья стало %d\n", h.Health)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100, Damage: 30, Def: 20}

	myHero.Attack()
	myHero.Defense()
	myHero.Special(30)
}
