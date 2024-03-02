package maps

import "fmt"

func MapExample1() {
	// Инициализация
	//m := make(map[int]int)
	//var m map[int]int
	m := map[int]int{}

	// Вставка
	m[2] = 0
	//fmt.Println(m)

	// Поиск
	if v, ok := m[2]; ok {
		fmt.Println(v)
	}
	//if v, ok := m[2]; ok {
	//	fmt.Println("Value found", v)
	//}

	// Удаление
	//delete(m, 2)
	//fmt.Println(m)
}
