package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ExampleCycle() {
	ExampleCyclePrimes()
	// ExampleCycleVariable()
	// ExampleCycleFiles()
	// ExampleCycleFilesAnon()
}

// Дан список чисел, нужно вывести все простые числа из этого списка.
func ExampleCyclePrimes() {
	nums := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	response := []int{}

NUMS_LOOP:
	for _, num := range nums {
		if num < 2 {
			continue
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				continue NUMS_LOOP
			}
		}
		response = append(response, num)
	}

	log.Println(response)
}

// Поведение переменных в цикле range поменяется в Go 1.22.
func ExampleCycleVariable() {
	nums := []int{0, 10, 20, 30, 40, 50}
	for i, v := range nums {
		iCycle, vCycle := i, v
		log.Println(i, v)
		defer func() {
			log.Println(iCycle, vCycle, i, v)
		}()
	}
}

// Дан список файлов, для существующих файлов вывести название и заголовок
// Заголовок это часть файла до первой пустой строки.
func ExampleCycleFiles() {
	fileNames := []string{"cycle.txt", "not_exists.go", "cycle.go"}
	for _, fName := range fileNames {
		fd, err := os.OpenFile(fName, os.O_RDONLY, 0666)
		if err != nil {
			continue
		}
		defer func(fName string, fd *os.File) {
			fd.Close()
			log.Println("Defer", fName)
		}(fName, fd)

		var sb strings.Builder
		reader := bufio.NewReader(fd)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Println(fName, "ReadLine error", err)
				break
			}
			line = strings.TrimSpace(line)
			if line == "" {
				break
			}
			sb.WriteString(line)
			sb.WriteRune('\n')
		}
		log.Println(fName)
		log.Println("-----------------")
		log.Println(sb.String())
	}
}

func ExampleCycleFilesAnon() {
	fileNames := []string{"cycle.txt", "not_exists.go", "cycle.go"}
	for _, fName := range fileNames {
		func() { // Добавилась эта строка
			file, err := os.OpenFile(fName, os.O_RDONLY, 0666)
			if err != nil {
				return // continue заменил на return
			}
			defer func(fName string, fd *os.File) {
				err = fd.Close()
				if err != nil {
					log.Println(err)
				}
				log.Println("Defer", fName)
			}(fName, file)

			var builder strings.Builder
			reader := bufio.NewReader(file)
			for {
				var line string
				line, err = reader.ReadString('\n')
				if err != nil {
					log.Println(fName, "ReadLine error", err)
					return
				}
				line = strings.TrimSpace(line)
				if line == "" {
					break
				}
				builder.WriteString(line)
				builder.WriteRune('\n')
			}
			log.Println(fName)
			log.Println("-----------------")
			log.Println(builder.String())
		}() // Добавилась эта строка
	}
}
