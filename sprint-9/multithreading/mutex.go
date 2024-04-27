package multithreading

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	//fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
}

func RunMutex() {
	var wg sync.WaitGroup
	balance = 1000
	wg.Add(2)
	go deposit(500, &wg)
	go deposit(700, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}
