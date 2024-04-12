package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var database = map[string]string{ //nolint:gochecknoglobals // it's learning code
	"Сергей":  "Омск",
	"Соня":    "Москва",
	"Алексей": "Калининград",
	"Миша":    "Москва",
	"Дима":    "Челябинск",
	"Алина":   "Красноярск",
	"Егор":    "Пермь",
	"Коля":    "Красноярск",
	"Артём":   "Владивосток",
	"Петя":    "Михайловка",
}

var offsetUTC = map[string]int{ //nolint:gochecknoglobals // it's learning code
	"Москва":          3,  //nolint:gomnd // it's learning code
	"Санкт-Петербург": 3,  //nolint:gomnd // it's learning code
	"Новосибирск":     7,  //nolint:gomnd // it's learning code
	"Екатеринбург":    5,  //nolint:gomnd // it's learning code
	"Нижний Новгород": 3,  //nolint:gomnd // it's learning code
	"Казань":          3,  //nolint:gomnd // it's learning code
	"Челябинск":       5,  //nolint:gomnd // it's learning code
	"Омск":            6,  //nolint:gomnd // it's learning code
	"Самара":          4,  //nolint:gomnd // it's learning code
	"Ростов-на-Дону":  3,  //nolint:gomnd // it's learning code
	"Уфа":             5,  //nolint:gomnd // it's learning code
	"Красноярск":      7,  //nolint:gomnd // it's learning code
	"Воронеж":         3,  //nolint:gomnd // it's learning code
	"Пермь":           5,  //nolint:gomnd // it's learning code
	"Волгоград":       3,  //nolint:gomnd // it's learning code
	"Краснодар":       3,  //nolint:gomnd // it's learning code
	"Калининград":     2,  //nolint:gomnd // it's learning code
	"Владивосток":     10, //nolint:gomnd // it's learning code
}

func whatTime(city string) string {
	utcTime := time.Now().UTC()
	friendTime := utcTime.Add(time.Duration(offsetUTC[city]) * time.Hour)
	return friendTime.Format("15:04")
}

func formatCountFriends(count int) string {
	if count == 1 {
		return "1 друг"
	}
	if count >= 2 && count <= 4 {
		return fmt.Sprintf("%d друга", count)
	}
	return fmt.Sprintf("%d друзей", count)
}

func processAlice(query string) string {
	if query == "сколько у меня друзей?" {
		return fmt.Sprintf("У тебя %s", formatCountFriends(len(database)))
	}
	if query == "кто все мои друзья?" {
		var friends []string
		for name := range database {
			friends = append(friends, name)
		}
		// сортируем друзей по алфавиту
		sort.Strings(friends)
		return fmt.Sprintf("Твои друзья: %s", strings.Join(friends, ", "))
	}

	if query == "где все мои друзья?" {
		uniqueCities := make(map[string]int)
		// заполняем мапу без дублирования городов
		for _, city := range database {
			uniqueCities[city] = 1
		}
		var cities []string
		// получаем уникальные названия городов
		for city := range uniqueCities {
			cities = append(cities, city)
		}

		// сортируем города по алфавиту
		sort.Strings(cities)
		return fmt.Sprintf("Твои друзья в городах: %s", strings.Join(cities, ", "))
	}
	return "неизвестный запрос"
}

func processFriend(name, query string) string {
	city, found := database[name]
	if found {
		if query == "ты где?" {
			return fmt.Sprintf("%s в городе %s", name, city)
		}
		if query == "который час?" {
			if _, ok := offsetUTC[city]; !ok {
				return fmt.Sprintf("Не могу определить время в городе %s", city)
			}
			// добавьте проверку запроса - который час?
			// проверьте, существует ли город в offsetUTC
			// если нет, то верните строку:
			// "Не могу определить время в городе <название_города>"
			// если город есть, то вызовете whatTime() и
			// верните "Там сейчас <время>"
			return fmt.Sprintf("Там сейчас %s", whatTime(city))
		}
		return "неизвестный запрос"
	}
	return fmt.Sprintf("У тебя нет друга по имени %s", name)
}

func processQuery(query string) string {
	elements := strings.Split(query, ", ")
	if len(elements) != 2 {
		return "неизвестный запрос"
	}
	if elements[0] == "Алиса" {
		return processAlice(elements[1])
	}
	return processFriend(elements[0], elements[1])
}

func runner() {
	queries := []string{
		"Алиса, сколько у меня друзей?",
		"Алиса, кто все мои друзья?",
		"Алиса, где все мои друзья?",
		"Алиса, кто виноват?",
		"Коля, ты где?",
		"Соня, что делать?",
		"Антон, ты где?",
		"Алексей, который час?",
		"Артём, который час?",
		"Антон, который час?",
		"Петя, который час?",
	}
	for _, query := range queries {
		fmt.Println(query, "-", processQuery(query))
	}
}

func main() {
	runner()
}
