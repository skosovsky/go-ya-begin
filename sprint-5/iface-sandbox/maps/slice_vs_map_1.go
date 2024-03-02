package maps

import (
	"fmt"
	"time"
)

func SliceVsMapExample1() {
	const elementsCount = 50_000_000
	s, m := getFilledData(elementsCount)

	valueToFind := elementsCount - 1

	start := time.Now()
	for _, v := range s {
		if v == valueToFind {
			fmt.Println("Элемент в слайсе найден")
			break
		}
	}
	sliceElapsedTime := time.Since(start)
	fmt.Printf("Поиск элемента занял: %s\n", sliceElapsedTime)

	start = time.Now()
	if _, ok := m[valueToFind]; ok {
		fmt.Println("Элемент в мапе найден")
	}
	mapElapsedTime := time.Since(start)
	fmt.Printf("Поиск элемента занял: %s\n", mapElapsedTime)

	fmt.Printf("Поиск в мапе быстрее поиска в слайсе в %d раз(а)!\n", int64(sliceElapsedTime)/int64(mapElapsedTime))
}
