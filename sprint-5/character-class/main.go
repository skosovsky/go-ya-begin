package main

import (
	"fmt"
)

// Hero описывает героя с общими полями для всех классов
type Hero struct {
	Name      string
	ClassName string
	Health    int
	Damage    int
	Def       int
	Inventory
}

// Attack возвращает строку с информацией о нанесённом уроне
func (h *Hero) Attack() string {
	return fmt.Sprintf("Ваш персонаж нанёс урон, равный %d", h.Damage)
}

// Defense возвращает строку с информацией о заблокированном уроне
func (h *Hero) Defense() string {
	return fmt.Sprintf("Ваш персонаж заблокировал %d урона", h.Def)
}

// Inventory описывает инвентарь
type Inventory struct {
	Items map[string]int
}

// Take добавляет предмет в инвентарь
func (i *Inventory) Take(item string) string {
	i.Items[item]++
	if _, isAvailable := i.Items[item]; isAvailable && i.Items[item] > 1 {
		return fmt.Sprintf("Вы положили %s в инвентарь. Количество: %d", item, i.Items[item])
	}
	return fmt.Sprintf("Вы положили %s в инвентарь", item)
}

// Drop удаляет предмет из инвентаря
func (i *Inventory) Drop(item string) string {
	if _, isAvailable := i.Items[item]; isAvailable {
		if i.Items[item] == 1 {
			delete(i.Items, item)
			return fmt.Sprintf("Вы выбросили %s", item)
		}

		i.Items[item]--
		return fmt.Sprintf("Вы выбросили %s\nОсталось %d", item, i.Items[item])
	}

	return fmt.Sprintf("У вас нет предмета %s", item)
}

// Warrior описывает класс «Воин»
type Warrior struct {
	Hero
}

// Buff — специальное умение для Warrior
func (w *Warrior) Buff() string {
	w.Def += 20
	return fmt.Sprintf("%s класса %s увеличил свою защиту.\nЗащита %s теперь %d.\n", w.Name, w.ClassName, w.Name, w.Def)
}

// Mage описывает класс «Маг»
type Mage struct {
	Hero
}

// Buff — специальное умение для Mage
func (m *Mage) Buff() string {
	m.Damage += 30
	return fmt.Sprintf("%s класса %s усилил свою атаку.\nАтака %s теперь %d.\n", m.Name, m.ClassName, m.Name, m.Damage)
}

// Healer описывает класс «Лекарь»
type Healer struct {
	Hero
}

// Buff — специальное умение для Healer
func (h *Healer) Buff() string {
	h.Health += 50
	return fmt.Sprintf("%s класса %s увеличил своё здоровье.\nЗдоровье %s теперь %d.\n", h.Name, h.ClassName, h.Name, h.Health)
}

func main() {
	// Маг
	// Инвентарь для мага
	mageInventory := Inventory{make(map[string]int)}

	// Структура для мага
	mage := Mage{Hero{Name: "Мерлин", ClassName: "Маг", Health: 100, Damage: 30, Def: 20, Inventory: mageInventory}}

	fmt.Println("Я", mage.Name, "класса", mage.ClassName) // Представимся

	fmt.Println(mage.Attack())  // Маг атакует
	fmt.Println(mage.Defense()) // Маг защищается
	fmt.Println(mage.Buff())    // Маг баффается

	fmt.Println(mage.Take("Посох")) // Маг кладёт посох в инвентарь
	fmt.Println(mage.Drop("Посох")) // Посох не понравился, выбрасывает посох

	fmt.Println() // Разделим вывод для персонажей

	// Воин
	// Инвентарь для воина
	warriorInventory := Inventory{make(map[string]int)}

	// Структура для воина
	warrior := Warrior{Hero{Name: "Арагорн", ClassName: "Воин", Health: 100, Damage: 30, Def: 20, Inventory: warriorInventory}}

	fmt.Println("Я", warrior.Name, "класса", warrior.ClassName) // Представимся

	fmt.Println(warrior.Attack())  // Воин атакует
	fmt.Println(warrior.Defense()) // Воин защищается
	fmt.Println(warrior.Buff())    // Воин баффается

	fmt.Println(warrior.Take("Меч"))    // Воин кладёт в инвентарь меч
	fmt.Println(warrior.Take("Шлем"))   // Воин кладёт в инвентарь шлем
	fmt.Println(warrior.Take("Наручи")) // Воин кладёт в инвентарь наручи
	fmt.Println(warrior.Take("Шлем"))   // Ещё один шлем в инвентарь
	fmt.Println(warrior.Drop("Шлем"))   // Много шлемов, один выкинул
	fmt.Println(warrior.Drop("Сапоги")) // Попробовал выкинуть сапоги, но их же и так нет

	fmt.Println() // Разделим вывод для персонажей

	// Лекарь
	// инвентарь для лекаря
	healerInventory := Inventory{make(map[string]int)}

	// Структура для лекаря
	healer := Healer{Hero{Name: "Елена Малышева", ClassName: "Лекарь", Health: 100, Damage: 30, Def: 20, Inventory: healerInventory}}

	fmt.Println("Я", healer.Name, "класса", healer.ClassName) // Представимся

	fmt.Println(healer.Attack())  // Лекарь атакует
	fmt.Println(healer.Defense()) // Лекарь защищается
	fmt.Println(healer.Buff())    // Лекарь накладывает на себя заклинание, или бафф

	fmt.Println(healer.Take("Амулет")) // Лекарь кладёт в инвентарь амулет
	fmt.Println(healer.Take("Плащ"))   // Лекарь кладёт в инвентарь плащ
}
