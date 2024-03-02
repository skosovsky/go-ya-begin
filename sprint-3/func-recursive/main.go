package main

func main() {
	numsSum([]int{1, 1, 1})
	binSearch([]int{1, 2, 3}, 1)
}

func numsSum(s []int) int {
	if len(s) == 0 {
		return 0
	}

	return s[0] + numsSum(s[1:])
	// Берем первый элемент слайса и прибавляем к нему вызов функции с оставшимися элементами
}

func binSearch(s []int, el int) int {
	if len(s) == 1 {
		if s[0] == el {
			return el
		}
		return -1
	}

	if el > s[:len(s)/2][len(s)/2-1] {
		res := binSearch(s[len(s)/2:], el)
		return res
	} else {
		res := binSearch(s[:len(s)/2], el)
		return res
	}
}
