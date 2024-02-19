package maps

import (
	"fmt"
	"sync"
	"time"
)

func getFilledData(elementsCount int) ([]int, map[int]int) {
	var s []int
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()

		for i := 0; i < elementsCount; i++ {
			s = append(s, i)
		}
		fmt.Printf("Запись слайса заняла: %s\n", time.Since(start))
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()

		for i := 0; i < elementsCount; i++ {
			m[i] = i
		}
		fmt.Printf("Запись мапы заняла: %s\n", time.Since(start))
	}()
	wg.Wait()

	return s, m
}
