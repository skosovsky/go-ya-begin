package multithreading

import (
	"log"
)

/*
func RunBonus() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)
	out4 := fanOut(in)

	m := map[string]int{
		"case 1": 0,
		"case 2": 0,
		"case 3": 0,
		"case 4": 0,
	}

	counter := 0
	for {
		select {
		case value, ok := <-out1:
			m["case 1"]++
			if ok {
				fmt.Println("Output 1 got:", value)
			} else {
				counter |= 1 << 0
			}
		case value, ok := <-out2:
			m["case 2"]++
			if ok {
				fmt.Println("Output 2 got:", value)
			} else {
				counter |= 1 << 1
			}
		case value, ok := <-out3:
			m["case 3"]++
			if ok {
				fmt.Println("Output 3 got:", value)
			} else {
				counter |= 1 << 2
			}
		case value, ok := <-out4:
			m["case 4"]++
			if ok {
				fmt.Println("Output 4 got:", value)
			} else {
				counter |= 1 << 3
			}
		}

		if counter == 15 {
			break
		}
	}

	fmt.Println()
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
		return
	}()

	return ch
}
*/

// Optimal -------------------------------------------------------------------------------------------------------------

/*
func RunBonus() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)
	out4 := fanOut(in)

	out := fanIn(out1, out2, out3, out4)

	for value := range out {
		fmt.Println("Output got:", value)
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for data := range in {
			out <- data
		}
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, w := range work {
			ch <- w
		}
	}()
	return ch
}
*/

func RunBonus() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)
	out4 := fanOut(in)

	m := map[string]int{
		"case 1": 0,
		"case 2": 0,
		"case 3": 0,
		"case 4": 0,
	}

	counter := 0
	for {
		select {
		case value, ok := <-out1:
			m["case 1"]++
			if ok {
				log.Println("Output 1 got:", value)
			} else {
				counter |= 1 << 0
				out1 = nil
			}
		case value, ok := <-out2:
			m["case 2"]++
			if ok {
				log.Println("Output 2 got:", value)
			} else {
				counter |= 1 << 1
				out2 = nil
			}
		case value, ok := <-out3:
			m["case 3"]++
			if ok {
				log.Println("Output 3 got:", value)
			} else {
				counter |= 1 << 2
				out3 = nil
			}
		case value, ok := <-out4:
			m["case 4"]++
			if ok {
				log.Println("Output 4 got:", value)
			} else {
				counter |= 1 << 3
				out4 = nil
			}
		}

		if counter == 15 {
			break
		}
	}

	log.Println()
	for k, v := range m {
		log.Printf("%s: %d\n", k, v)
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
		return
	}()

	return ch
}
