package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(200 * time.Millisecond) //nolint:mnd // it's learning code
	timer := time.NewTimer(4 * time.Second)          //nolint:mnd // it's learning code
	count := 0

	for count != 50 {
		select {
		case <-ticker.C:
			count++
			if count%5 == 0 {
				fmt.Print("o") //nolint:forbidigo // it's learning code
			} else {
				fmt.Print(".") //nolint:forbidigo // it's learning code
			}
		case <-timer.C:
			return
		}
	}
}
