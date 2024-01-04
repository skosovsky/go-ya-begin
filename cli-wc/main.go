package main

import (
	"fmt"
	"os"
	"strings"
)

type settings struct {
	linesFlag   bool // l
	wordsFlag   bool // w
	bytesFlag   bool // c
	symbolsFlag bool // m
	lenLineFlag bool // L
}

func main() {
	flags, files := getParams()
	fmt.Println(flags, files)

	if files == nil {
		// вызываем функцию по считыванию с консоли
	}
}

func getParams() (flags settings, files []string) {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		// default settings
		flags.linesFlag = true
		flags.wordsFlag = true
		flags.bytesFlag = true

		return flags, files // files = nil
	}

	arg := strings.Split(arguments[0], "")

	if arg[0] == "-" && len(arg) > 1 { // И проверяем, а это флаги?
		for _, flag := range arg[1:] {
			switch flag {
			case "l":
				flags.linesFlag = true
			case "w":
				flags.wordsFlag = true
			case "m":
				flags.symbolsFlag = true
			case "L":
				flags.lenLineFlag = true
			case "c":
				flags.bytesFlag = true
			default: // Если попались не флаги, выходим
				fmt.Println("wc: illegal option --", flag)
				fmt.Println("usage: wc [-Lclmw] [file ...]")
				os.Exit(0)
			}
		}

		files = arguments[1:]
		if len(files) == 0 {
			files = nil
		}

		return flags, files
	}

	// default settings
	flags.linesFlag = true
	flags.wordsFlag = true
	flags.bytesFlag = true

	files = arguments
	if len(files) == 0 {
		files = nil
	}

	return flags, files
}
