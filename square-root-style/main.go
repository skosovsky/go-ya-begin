// Package main is main package
package main

import (
	"fmt"
	"math"
)

// Message is a welcome message.
var Message = `Добро пожаловать в самую лучшую программу для вычисления
квадратного корня из заданного числа`

// CalculateSquareRoot return the square root of a number.
func CalculateSquareRoot(myNumber float64) float64 {
	return math.Sqrt(myNumber)
}

// calc check if a number is positive and print the square root if it's positive.
func calc(number float64) {
	if number < 0 {
		return
	} else {
		fmt.Println("Мы вычислили квадратный корень из введённого вами числа. Это будет:", CalculateSquareRoot(number))
	}
}

func main() {
	fmt.Println(Message)
	calc(25.5)
}
