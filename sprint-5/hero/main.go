package main

import "fmt"

type Hero struct {
	Name   string // имя героя
	Level  int    // уровень героя
	Health int    // здоровье героя
}

// IncLevel увеличивает уровень героя на единицу.
// добавьте метод IncLevel().
func (h *Hero) IncLevel() {
	h.Level++
}

// ChangeHealth изменяет здоровье героя на указанную величину.
// добавьте метод ChangeHealth(dif int).
func (h *Hero) ChangeHealth(dif int) {
	h.Health += dif
}

func main() {
	hero := Hero{
		Name:   "Буратино",
		Level:  10,
		Health: 35,
	}
	// вызовите метод IncLevel()
	hero.IncLevel()
	// вызовите метод ChangeHealth(10)
	hero.ChangeHealth(10)

	fmt.Println(hero)
}
