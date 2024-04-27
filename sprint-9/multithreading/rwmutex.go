package multithreading

import (
	"fmt"
	"sync"
)

// Counter Счетчик с защитой мьютексом
type Counter struct {
	sync.RWMutex
	count int
}

// Increment увеличивает счетчик на 1
func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

// Value возвращает текущее значение счетчика
func (c *Counter) Value() int {
	c.RLock()
	defer c.RUnlock()
	return c.count
}

func RunRWMutex() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(2)

	inc := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter.Increment()
		}
	}

	// Горутина для инкремента
	go inc()

	// Горутина для инкремента
	go inc()

	wg.Wait()
	fmt.Println("Final Count:", counter.Value())
}
