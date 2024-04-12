package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

// rnd - генератор псевдослучайных чисел.
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// randNum возвращает случайное число в интервале [min, max].
func randNum(min, max int) int {
	return rnd.Intn(max-min+1) + min
}

// input запрашивает и возвращает ввод пользователя в консоли.
func input(title string) string {
	log.Print(title)
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		log.Println(err)
	}
	return s
}

func setEnemyHealth() int {
	return randNum(80, 120) //nolint:gomnd // it's learning code
}

func getLiteAttack() int {
	return randNum(2, 5) //nolint:gomnd // it's learning code
}

func getMidAttack() int {
	return randNum(15, 25) //nolint:gomnd // it's learning code
}

func getHardAttack() int {
	return randNum(30, 40) //nolint:gomnd // it's learning code
}

func compareValues(enemyHealth, userTotalAttack int) bool {
	pointDifference := enemyHealth - userTotalAttack
	if pointDifference < 0 {
		return true
	}
	return pointDifference <= 10 //nolint:gomnd // it's learning code
}

func getUserAttack() int {
	total := 0

	for i := 0; i < 5; i++ {
		inputAttack := input("Введи тип атаки: ")

		var attackValue int
		switch inputAttack {
		case "lite":
			attackValue = getLiteAttack()
		case "mid":
			attackValue = getMidAttack()
		case "hard":
			attackValue = getHardAttack()
		default:
			log.Println("Неизвестный тип атаки:", inputAttack)
			continue
		}
		total += attackValue
		log.Println("Количество очков твоей атаки:", attackValue)
	}
	return total
}

func runGame() bool {
	enemyHealth := setEnemyHealth()
	userTotalAttack := getUserAttack()
	log.Println("Тобой нанесён урон противнику равный", userTotalAttack)
	log.Println("Очки здоровья противника до твоей атаки", enemyHealth)
	if compareValues(enemyHealth, userTotalAttack) {
		log.Println("Ура! Победа за тобой!")
	} else {
		log.Println("В этот раз не повезло :( Бой проигран.")
	}
	answer := input("Чтобы сыграть ещё раз, введи букву [y] или [Y]: ")
	return strings.ToUpper(answer) == "Y"
}

func main() {
	intro := `РАССЧИТАЙ И ПОБЕДИ!
Загрузка...

Твоя цель — за 5 ходов набрать такое количество очков урона противнику,
которое попадет в диапазон +– 10 от значения здоровья противника.

Значение здоровья противника генерируется случайным образом
в диапазоне от 80 до 120 очков.

В твоём распоряжении три вида атак:
lite — урон от 2 до 5 очков;
mid — урон от 15 до 25 очков;
hard — урон от 30 до 40 очков.
ВПЕРЁД К ПОБЕДЕ!!!`

	log.Println(intro)

	for runGame() {
	}
}
