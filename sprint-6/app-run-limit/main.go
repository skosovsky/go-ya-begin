package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Counter возвращает строку для записи в файл.
func Counter(count int, t time.Time) string {
	date := t.Format("02.01.2006")

	return fmt.Sprintf("%d %s", count, date)
}

// Limits возвращает количество дней и запусков.
func Limits() (int, int, error) {
	const hoursInDay = 24
	app, err := os.Executable()
	if err != nil {
		return 0, 0, fmt.Errorf("filed to get path file: %w", err)
	}

	name := filepath.Join(filepath.Dir(app), "data.txt")
	if _, err = os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			out := Counter(1, time.Now())
			err = os.WriteFile(name, []byte(out), 0600)
			return 0, 1, fmt.Errorf("filed to write file: %w", err)
		}
		return 0, 0, fmt.Errorf("filed to get stat from file: %w", err)
	}

	var data []byte

	data, err = os.ReadFile(name)
	if err != nil {
		return 0, 0, fmt.Errorf("filed to read file: %w", err)
	}

	counter, date, err := ParseCounter(string(data))
	if err != nil {
		return 0, 0, err
	}

	counter++
	if err = os.WriteFile(name, []byte(Counter(counter, date)), 0600); err != nil {
		return 0, 0, fmt.Errorf("filed to write file: %w", err)
	}
	duration := time.Since(date)
	return int(duration.Hours()) / hoursInDay, counter, nil
}

// ParseCounter разбирает информацию из файла, возвращает значение счётчика и дату первого запуска.
func ParseCounter(input string) (int, time.Time, error) {
	counterDate := strings.Split(input, " ")
	counter, err := strconv.Atoi(counterDate[0])
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("filed to convert string: %w", err)
	}

	date, err := time.Parse("02.01.2006", counterDate[1])
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("filed to parse date: %w", err)
	}

	return counter, date, nil
}

func main() {
	days, counter, err := Limits()

	if err != nil {
		fmt.Println("Ошибка", err) //nolint:forbidigo // need fmt
		return
	}

	fmt.Printf("Количество дней: %d\nКоличество запусков: %d\n", days, counter) //nolint:forbidigo // need fmt
	if days > 14 || counter > 50 {
		fmt.Println("Запросите новую версию") //nolint:forbidigo // need fmt
		return
	}
	fmt.Println("Программа готова к работе") //nolint:forbidigo // need fmt
}
