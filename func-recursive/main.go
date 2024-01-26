package main

func main() {
	numsSum([]int{1, 1, 1})
}

func numsSum(s []int) int {
	if len(s) == 0 {
		return 0
	}

	return s[0] + numsSum(s[1:])
	// Берем первый элемент слайса и прибавляем к нему вызов функции с оставшимися элементами
}
