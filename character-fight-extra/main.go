package main

import (
	"fmt"
)

const (
	msgTmplAttack = "%s класса %s нанёс %d урона"
	msgTmplDef    = "%s класса %s заблокировал %d урона"
)

// Hero описывает героя с общими полями для всех классов
type Hero struct {
	Name      string // имя
	ClassName string // имя класса
	Health    int    // здоровье
	Damage    int    // урон
	Def       int    // защита
}

// Attack возвращает строку с информацией о нанесенном уроне
func (h Hero) Attack() string {
	return fmt.Sprint("Ваш персонаж нанёс урон равный", h.Damage)
}

// Defense возвращает строку с информацией о заблокированном уроне
func (h Hero) Defense() string {
	return fmt.Sprint("Ваш персонаж заблокировал ", h.Def, " урона")
}

// Warrior описывает класс `Воин`
type Warrior struct {
	Hero            // встроенная структура Hero
	ClassDamage int // урон для класса Воин
	ClassDef    int // защита для класса Воин
}

// Метод Attack() для воина
func (w Warrior) Attack() string {
	return fmt.Sprintf(msgTmplAttack, w.Name, w.ClassName, w.Damage+w.ClassDamage)
}

// Метод Defense() для воина
func (w Warrior) Defense() string {
	return fmt.Sprintf(msgTmplDef, w.Name, w.ClassName, w.Def+w.ClassDef)
}

// Mage описывает класс `Маг`
type Mage struct {
	Hero            // встроенная структура Hero
	ClassDamage int // урон для класса Маг
	ClassDef    int // защита для класса Маг
}

// Метод Attack() для мага
func (m Mage) Attack() string {
	return fmt.Sprintf(msgTmplAttack, m.Name, m.ClassName, m.Damage+m.ClassDamage)
}

// Метод Defense() для мага
func (m Mage) Defense() string {
	return fmt.Sprintf(msgTmplDef, m.Name, m.ClassName, m.Def+m.ClassDef)
}

// Healer описывает класс `Лекарь`
type Healer struct {
	Hero            // встроенная структура Hero
	ClassDamage int // урон для класса Лекарь
	ClassDef    int // защита для класса Лекарь
}

// Метод Attack() для лекаря
func (h Healer) Attack() string {
	return fmt.Sprintf(msgTmplAttack, h.Name, h.ClassName, h.Damage+h.ClassDamage)
}

// Метод Defense() для лекаря
func (h Healer) Defense() string {
	return fmt.Sprintf(msgTmplDef, h.Name, h.ClassName, h.Def+h.ClassDef)
}

func main() {
	// воин
	warrior := Warrior{
		Hero:        Hero{Name: "Арагорн", ClassName: "Воин", Health: 100, Damage: 30, Def: 20},
		ClassDamage: 20,
		ClassDef:    30,
	}

	// воин атакует
	fmt.Println(warrior.Attack())
	// воин защищается
	fmt.Println(warrior.Defense())

	// маг
	mage := Mage{
		Hero:        Hero{Name: "Мерлин", ClassName: "Маг", Health: 100, Damage: 30, Def: 20},
		ClassDamage: 30,
		ClassDef:    10,
	}
	// маг атакует
	fmt.Println(mage.Attack())
	// маг защищается
	fmt.Println(mage.Defense())

	// лекарь
	healer := Healer{
		Hero:        Hero{Name: "Елена Малышева", ClassName: "Лекарь", Health: 100, Damage: 30, Def: 20},
		ClassDamage: 10,
		ClassDef:    20,
	}
	// лекарь атакует
	fmt.Println(healer.Attack())
	// лекарь защищается
	fmt.Println(healer.Defense())
}
