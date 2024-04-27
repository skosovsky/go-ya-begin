package multithreading

import (
	"log"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		log.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func RunWorkersWithBuffering() {
	const jobsCount = 5
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)
	defer close(results)

	for i := 1; i <= jobsCount; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= jobsCount; i++ {
		log.Println(<-results)
	}
}

func RunWorkersWithoutBuffering() {
	const jobsCount = 5
	jobs := make(chan int)
	results := make(chan int)
	defer close(results)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	go func() {
		for i := 1; i <= jobsCount; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	for i := 1; i <= jobsCount; i++ {
		<-results
	}
}
