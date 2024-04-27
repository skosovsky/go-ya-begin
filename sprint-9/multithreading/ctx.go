package multithreading

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"time"
)

func ctxWorker(ctx context.Context, wg *sync.WaitGroup, id int, resChan chan<- string, timeout time.Duration) {
	defer wg.Done()
	log.Printf("worker %d starting with timeout %f\n", id, timeout.Seconds())
	delay := time.NewTimer(timeout)
	select {
	case <-delay.C:
		resChan <- fmt.Sprintf("worker %d finished", id)
	case <-ctx.Done():
		log.Printf("worker %d error: %s\n", id, ctx.Err())
		if !delay.Stop() {
			<-delay.C
		}
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func RunCtxExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const workersCount = 10
	resChan := make(chan string, workersCount)
	defer close(resChan)
	var wg sync.WaitGroup
	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		timeout := time.Second * time.Duration(randRange(1, 10))
		go ctxWorker(ctx, &wg, i, resChan, timeout)
	}

	log.Println(<-resChan)
	cancel()
	wg.Wait()
}
