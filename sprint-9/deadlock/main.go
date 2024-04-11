package main

import (
	"log"
	"sync"
	"time"
)

var muA, muB sync.Mutex //nolint:gochecknoglobals // it's learning code
var a, b int            //nolint:gochecknoglobals // it's learning code

func A() {
	muA.Lock()
	a++
	muB.Lock()
	b++
	muB.Unlock()
	// 3 ...
	muA.Unlock()
}

func B() {
	muB.Lock()
	b++
	muA.Lock()
	a++
	muA.Unlock()
	// 3 ...
	muB.Unlock()
}

func main() {
	for range 1000000 {
		go A()
		go B()
	}

	log.Println(a)
	log.Println(b)

	time.Sleep(2 * time.Second) //nolint:gomnd // it's learning code
}
