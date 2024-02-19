package main

import "fmt"

func main() {
	// исходный массив
	numbers := [...]int{1, 3, 8, 19, 42}
	numbersPtr := make([]*int, 0, len(numbers))

	// 1. создайте и заполните слайс указателей на элементы массива
	for i := range numbers {
		numbersPtr = append(numbersPtr, &numbers[i])
	}

	// 2. пройдитесь по слайсу и умножьте на три все значения, на которые ссылаются указатели
	for _, v := range numbersPtr {
		*v *= 3
	}

	fmt.Println(numbers)
}
