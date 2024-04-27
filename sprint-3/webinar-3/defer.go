package main

import (
	"bufio"
	"log"
	"os"
)

func ExampleDefer() {
	ExampleDeferStack()
	// ExampleDeferRegular()
	// ExampleDeferParams()
	// ExampleDeferReturn()
	// ExampleDeferPanic()
	// ExampleDeferFile()
}

func ExampleDeferStack() {
	defer log.Println("defer 1")

	log.Println("function 1")
	log.Println("function 2")

	defer log.Println("defer 2")
	defer log.Println("defer 3")

	log.Println("function 3")
}

func deferHelpeer(s string) {
	log.Println("Defer helper", s)
}
func ExampleDeferRegular() {
	defer deferHelpeer("1")
	log.Println("Regular 1")
}

func ExampleDeferParams() {
	counter := 0
	deferHelper := func(v int) {
		log.Println("deferHelper", "counter", counter, "v", v)
	}
	defer deferHelper(1)
	counter = 1
	defer deferHelper(2)
	log.Println("ExampleDeferParams")
}

func returnHelper(s string) (ret string) {
	defer func() {
		if s == "defer" {
			ret = "This is SPARTA!!!"
		}
	}()
	return s
}
func ExampleDeferReturn() {
	log.Println("regular", returnHelper("regular"))
	log.Println("defer", returnHelper("defer"))
}

func helperDeferPanic() {
	// panic("I am panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered", err)
		}
	}()
	panic("I am panic")

	log.Println("After panic")
}
func ExampleDeferPanic() {
	log.Println("Before panic helper")
	helperDeferPanic()
	log.Println("After panic helper")
}

func ExampleDeferFile() {
	fd, err := os.OpenFile("./defer.go", os.O_RDONLY, 0666)
	if err != nil {
		log.Println("OpenFile error", err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println("ReadLine error", err)
		return
	}
	log.Println("ReadLine", line)
}
