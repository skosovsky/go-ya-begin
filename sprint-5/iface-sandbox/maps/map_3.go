package maps

import "fmt"

func MapExample3() {
	m := make(map[int]struct{})
	s := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}

	for _, v := range s {
		m[v] = struct{}{}
	}

	for k, _ := range m {
		fmt.Println(k)
	}
}
