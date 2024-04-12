package main

import (
	"sync"
)

var muA, muB sync.Mutex //nolint:gochecknoglobals // it's learning code

func A() {
	muA.Lock()
	// 1 ...
	muB.Lock()
	// 2 ...
	muB.Unlock() //nolint:staticcheck // it's learning code
	// 3 ...
	muA.Unlock()
}

func B() {
	muB.Lock()
	// 1 ...
	muA.Lock()
	// 2 ...
	muA.Unlock() //nolint:staticcheck // it's learning code
	// 3 ...
	muB.Unlock()
}

func main() {
	for range 30 {
		go A()
		B()
	}
}
