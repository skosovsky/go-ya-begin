package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {
	slc := []string{"Иван Фролов", "Ян Косовский", "Сергей Косовский", "Марк Косовский"}
	fmt.Println(nameSort(slc))
}

func sortSlice(slc []int) []int {
	var result []int
	var minNum, maxNum int

	if len(slc) <= 1 {
		return slc
	}

	minNum, maxNum = slc[0], slc[0]

	for _, v := range slc {
		if minNum > v {
			minNum = v
		}

		if maxNum < v {
			maxNum = v
		}
	}

	for i := maxNum; i >= minNum; i-- {
		for _, v := range slc {
			if i == v {
				result = append(result, i)
			}
		}
	}
	return result
}

func sortSliceLib(slc []int) []int {
	sort.Slice(slc, func(i, j int) bool {
		return slc[i] > slc[j]
	})

	return slc
}

func sortSliceLibNew(slc []int) []int {
	var result []int

	if len(slc) == 0 {
		return slc
	}

	slices.Sort(slc)

	for i := len(slc) - 1; i >= 0; i-- {
		result = append(result, slc[i])
	}

	return result
}

func findX(slc []int, num int) bool {
	for _, v := range slc {
		if v == num {
			return true
		}
	}

	return false
}

func findXLib(slc []int, num int) bool {
	idx := sort.Search(len(slc), func(i int) bool {
		return slc[i] <= num
	})

	if idx < len(slc) && slc[idx] == num {
		return true
	}

	return false
}

func nameSort(slc []string) []string {
	var swapSlc, sortSlc []string
	sortSlc = []string{}

	// Меням фамилию и имя местами
	for _, v := range slc {
		fullName := strings.Split(v, " ")

		if len(fullName) != 2 {
			continue
		}

		swapSlc = append(swapSlc, strings.Join([]string{fullName[1], fullName[0]}, " "))
	}

	// Сортируем по фамилию и имени
	slices.Sort(swapSlc)

	// Меням фамилию и имя обратно
	for _, v := range swapSlc {
		fullName := strings.Split(v, " ")

		sortSlc = append(sortSlc, strings.Join([]string{fullName[0], fullName[1]}, " "))
	}

	return sortSlc
}
