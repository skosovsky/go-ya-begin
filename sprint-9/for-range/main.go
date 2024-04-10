package main

import (
	"fmt"
)

func do(in, out chan int) {
	defer close(out)
	for v := range in {
		// отправляем результат в другой канал
		out <- 2 * v
	}
}

func main() {
	chIn := make(chan int)
	chOut := make(chan int)

	// запускаем горутину, которая преобразует числа
	go do(chIn, chOut)

	// горутина, которая отправляет числа на обработку
	go func() {
		// закрываем канал после отправки всех чисел
		defer close(chIn)

		for i := 0; i <= 50; i++ {
			chIn <- i
		}
	}()
	// в цикле читаем числа из результирующего канала
	for v := range chOut {
		fmt.Print(v, " ")
	}
}
