package main

import (
	"bufio"
	"fmt"
	"log"
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

type answer struct {
	lines   int
	words   int
	bytes   int
	symbols int
	lenLine int
	warning string
}

func main() {
	flags, files := getParams()
	var testff answer
	fmt.Println(flags, files)

	if files == nil {
		fmt.Println(calcFromStdin())
	} else {
		for _, file := range files {
			fmt.Println(calcFromFile(file))
		}
	}

	fmt.Printf("%8d%8d%8d%8d%8d %s\n", testff.bytes, testff.bytes, testff.bytes, testff.bytes, testff.bytes, "file")

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

// 世界 qwer
// 1 2 12 8 11
// l-w-c--m-L-
func calcFromStdin() (result answer) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result.lines++ // Count lines

		words := strings.Split(line, " ")
		result.words += len(words) // Count words

		bytes := []byte(line)
		result.bytes += len(bytes) + 1 // Count bytes + /n

		runes := []rune(line)
		result.symbols += len(runes) + 1 // Count symbols

		if result.lenLine < len(line) {
			result.lenLine = len(line) // Count len
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return result
}

func calcFromFile(pathFile string) (result answer) {
	file, err := os.Open(pathFile)
	if err != nil {
		result.warning = "wc: " + pathFile + ": open: No such file or directory\n"
		return result
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lastLine := ""
	for scanner.Scan() {
		line := scanner.Text()
		result.lines++ // Count lines

		words := strings.Split(line, " ")
		result.words += len(words) // Count words

		bytes := []byte(line)
		result.bytes += len(bytes) + 1 // Count bytes + /n

		runes := []rune(line)
		result.symbols += len(runes) + 1 // Count symbols + /n

		if result.lenLine < len(line) {
			result.lenLine = len(line) // Count len
		}

		lastLine = line
	}

	if len(lastLine) != 0 { // if no /n in lastLine
		result.lines--
		result.bytes--
		result.symbols--
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return result
}
