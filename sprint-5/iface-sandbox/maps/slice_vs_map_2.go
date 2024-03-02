package maps

import (
	"fmt"
	"time"
)

func SliceVsMapExample2() {
	const elementsCount = 50_000_000
	s, m := getFilledData(elementsCount)

	start := time.Now()
	for _, v := range s {
		v++
	}
	sliceElapsedTime := time.Since(start)
	fmt.Printf("Перебор элементов слайса занял: %s\n", sliceElapsedTime)

	start = time.Now()
	for _, v := range m {
		v++
	}
	mapElapsedTime := time.Since(start)
	fmt.Printf("Перебор элементов мапы занял: %s\n", mapElapsedTime)

	fmt.Printf("Перебор в слайсе быстрее перебора в мапе в %d раз(а)!\n", int64(mapElapsedTime)/int64(sliceElapsedTime))
}
