package multithreading

import (
	"context"
	"log"
	"time"
)

func RunDanglingRoutines() {
	timeout := 5 * time.Second
	err := executeTaskWithTimeout(context.Background(), timeout)
	if err != nil {
		log.Println(err)
	}

	log.Println("point 1")
	time.Sleep(15 * time.Second)
	log.Println("point 2")

	//runtime.Goexit()
}

func executeTaskWithTimeout(ctx context.Context, timeout time.Duration) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	done := make(chan struct{})
	defer close(done)

	go func() {
		log.Println("point 3")
		executeTask()
		log.Println("point 4")
		done <- struct{}{}
		log.Println("point 5")
		//close(done)
		log.Println("point 6")
	}()

	select {
	case <-done:
		return nil
	case <-timeoutCtx.Done():
		return timeoutCtx.Err()
	}
}

func executeTask() {
	time.Sleep(10 * time.Second)
	log.Println("point 7")
}
