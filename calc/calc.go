package calc

// AddInts складывает все полученные параметры.
func AddInts(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
