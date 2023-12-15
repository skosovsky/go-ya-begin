package main

import "log"

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(i int) int {
		log.Println("Start", i)
		result := fn(i)
		log.Println("Finish", result)

		return result
	}
}

func Mult(n int) int {
	return n * 10
}

func Double(n int) int {
	return n * 2
}

func main() {
	f := LogDecorate(Mult)
	f(5)

	f = LogDecorate(Double)
	f(5)
}
