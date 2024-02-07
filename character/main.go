package main

import "fmt"

type Thing struct {
	Name   string
	Weight int
}

type Hero struct {
	Name          string           // имя героя
	Health, Power int              // Health здоровье, Power сила
	Loot          map[string]Thing // добыча
}

func main() {
	ken := Character{
		Name:   "Ken",
		Health: 100,
		Speed:  500,
		Power:  1000,
	}
	ken.Say("Привет! Как тебя зовут?")

	merlin := Magician{
		Character: Character{
			Name:   "Merlin",
			Health: 100,
			Speed:  250,
			Power:  400,
		},
		Magic: 700,
	}
	fmt.Println(merlin)

	var hero Hero
	hero.Name = "Арчибальд"
	hero.Loot = make(map[string]Thing)
	fmt.Println(hero)
}

type Character struct {
	Name   string // имя
	Health int    // здоровье
	Speed  int    // скорость
	Power  int    // сила
	Woman  bool   // true, если женский персонаж
}

func (c Character) String() string {
	s := fmt.Sprintf(`Имя: %s
Здоровье: %d
Скорость: %d
Сила: %d
`, c.Name, c.Health, c.Speed, c.Power)

	return s
}

func (c Character) Say(msg string) {
	fmt.Printf("%s: %s\n", c.Name, msg)
}

type Magician struct {
	Character
	Magic int
}

func (m Magician) String() string {
	s := fmt.Sprintf(`Имя: %s
Здоровье: %d
Скорость: %d
Сила: %d
Магия: %d
`, m.Name, m.Health, m.Speed, m.Power, m.Magic)

	return s
}
