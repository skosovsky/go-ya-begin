package main

import (
	"fmt"
	"log"
)

// Character содержит общую информацию о герое.
type Character struct {
	Name   string
	Health int
}

// Warrior описывает воина.
type Warrior struct {
	Character
	Power int
}

// Mage описывает мага.
type Mage struct {
	Character
	Magic int
}

// определите методы String() для Warrior и Mage.
func (w Warrior) String() string {
	return fmt.Sprintf("* %s", w.Name)
}

func (m Mage) String() string {
	return fmt.Sprintf("# %s", m.Name)
}

func main() {
	// проверяем работу метода String()
	log.Println(Warrior{Character: Character{Name: "Крестоносец"}})
	log.Println(Warrior{Character: Character{Name: "Коммандос"}})
	log.Println(Mage{Character: Character{Name: "Шаман"}})
	log.Println(Mage{Character: Character{Name: "Друид"}})
}
