package main

import "fmt"

func main() {
	count := PrintPrimes(300)
	fmt.Printf("\nКоличество простых чисел: %d\n", count)
}

func PrintPrimes(n int) int {
	s := make([]bool, n+1)

	count := 0
	for i := 2; i <= n; i++ {
		if !s[i] {
			count++
			fmt.Printf("%4d", i)
			for k := i + i; k <= n; k += i {
				s[k] = true
			}
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
	return count
}
