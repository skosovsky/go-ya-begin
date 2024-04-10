package main

import (
	"fmt"
	"sync"
)

var list []int        //nolint:gochecknoglobals // it's learning code
var wg sync.WaitGroup //nolint:gochecknoglobals // it's learning code
var mu sync.Mutex     //nolint:gochecknoglobals // it's learning code

func do() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		list = append(list, i)
		mu.Unlock()
	}
}

func main() {
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go do()
	}
	wg.Wait()
	// проверка содержимого у слайса
	sum := 0
	for _, v := range list {
		sum += v
	}

	fmt.Println(len(list), sum) //nolint:forbidigo // it's learning code
}
