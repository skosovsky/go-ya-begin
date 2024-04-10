package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(500 * time.Millisecond) //nolint:gomnd // it's learning code
	timer := time.NewTimer(10 * time.Second)    //nolint:gomnd // it's learning code

	var count int
	for {
		select {
		case <-ticker:
			count++
			fmt.Print(count, " ") //nolint:forbidigo // it's learning code
			if count == 10 {      //nolint:gomnd // it's learning code
				continue
				//	close(ticker) // never close <-chan only chan<- or chan
			}
		case <-timer.C:
			fmt.Println("Расчёт закончен") //nolint:forbidigo  // it's learning code
			return
		}
	}
}
