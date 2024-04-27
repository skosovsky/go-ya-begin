package multithreading

import (
	"fmt"
	"sync"
	"time"
)

func wgWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func RunWgWorkers() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go wgWorker(i, &wg)
	}
	wg.Wait() // Ожидание завершения всех горутин
}
