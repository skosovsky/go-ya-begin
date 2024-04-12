package main

import (
	"log"
)

// Hero описывает главного героя игры.
type Hero struct {
	Name   string
	Health int
	Damage int
	Def    int
}

// Attack Создаём метод атаки.
func (h *Hero) Attack() {
	log.Printf("%s нанес урон %d\n", h.Name, h.Damage)
}

// Defense Создаём метод защиты.
func (h *Hero) Defense() {
	log.Printf("%s заблокировал %d единиц урона\n", h.Name, h.Def)
}

// Special Создаём специальный метод.
func (h *Hero) Special(healthPoints int) {
	log.Printf("Количество здоровья было %d\n", h.Health)
	h.Health += healthPoints
	log.Printf("Количество здоровья стало %d\n", h.Health)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100, Damage: 30, Def: 20} //nolint:gomnd // it's learning code

	myHero.Attack()
	myHero.Defense()
	myHero.Special(30) //nolint:gomnd // it's learning code
}
