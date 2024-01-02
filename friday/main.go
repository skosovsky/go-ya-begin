package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	finish := time.Date(400, 12, 31, 0, 0, 0, 0, time.UTC)

	dayWeeks := map[time.Weekday]int{}

	for i := 0; start != finish; i++ {
		if start.Day() == 13 {
			dayWeeks[start.Weekday()]++
		}

		start = start.Add(time.Hour * 24)
	}

	fmt.Println(dayWeeks)
}
