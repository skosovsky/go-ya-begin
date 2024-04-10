package main

import (
	"fmt"
	"time"
)

const Count = 1000
const Lines = 4

var data [Lines][Count]int //nolint:gochecknoglobals // it's learning code

func main() {
	for i := range Lines {
		go func() {
			for j := range Count {
				data[i][j] = j
			}
		}()
	}

	time.Sleep(100 * time.Millisecond)
	// проверим как заполнен массив
	sum := 0
	for i := 0; i < Lines; i++ {
		for j := 0; j < Count; j++ {
			sum += data[i][j]
		}
	}
	fmt.Println(sum) //nolint:forbidigo // it's learning code
}
