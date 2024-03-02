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
	Location      string
}

func (h *Hero) String() string {
	s := fmt.Sprintf(`Имя: %s
Здоровье: %d
Сила: %d
Местоположение: %s
`, h.Name, h.Health, h.Power, h.Location)

	return s
}

func (h *Hero) Say(msg string) {
	fmt.Printf("%s: %s\n", h.Name, msg)
}

func (h *Hero) UpHealth(points int) {
	fmt.Println("Количество здоровья героя", h.Name, "было", h.Health)
	h.Health += points
	fmt.Println("Количество здоровья героя", h.Name, "стало", h.Health)
}

func (h *Hero) MoveTo(location string) {
	h.Location = location
	fmt.Printf("%s переместился в %s\n", h.Name, h.Location)
}

type Magician struct {
	Hero
	Magic int
}

func (m *Magician) String() string {
	s := fmt.Sprintf(`Имя: %s
Здоровье: %d
Сила: %d
Магия: %d
Местоположение: %s
`, m.Name, m.Health, m.Power, m.Magic, m.Location)

	return s
}

func main() {
	ken := Hero{
		Name:     "Ken",
		Health:   100,
		Power:    1000,
		Location: "столовая",
	}
	ken.Say("Привет! Как тебя зовут?")

	merlin := Magician{
		Hero: Hero{
			Name:     "Merlin",
			Health:   100,
			Power:    400,
			Location: "библиотека",
		},
		Magic: 700,
	}
	fmt.Println(merlin)

	var hero Hero
	hero.Name = "Арчибальд"
	hero.Loot = make(map[string]Thing)
	fmt.Println(hero)

	hero.UpHealth(20)
	fmt.Println(hero)

	myHero := Hero{Name: "Артур", Health: 100}
	myHero.MoveTo("тронный зал")
}
