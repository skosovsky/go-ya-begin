package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	slc := []int{1, 1, 1, 1, 2, 2}
	slc = sortSliceLibNew(slc)

	fmt.Println(slc)
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
