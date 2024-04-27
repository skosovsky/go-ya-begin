package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	BytesFlag    = 0x1
	SymbolsFlag  = 0x2
	LinesFlag    = 0x4
	WordsFlag    = 0x8
	LenLinesFlag = 0x10
)

type setting int //nolint:unused // it's not done

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
	var total answer
	flags, files := getParams()
	if files == nil {
		printAnswer(flags, calcFromStdin(), "file")
	} else {
		for _, file := range files {
			fileValue := calcFromFile(file)
			total.lines += fileValue.lines
			total.words += fileValue.words
			total.bytes += fileValue.bytes
			total.symbols += fileValue.symbols
			total.lenLine += fileValue.lenLine
			printAnswer(flags, fileValue, file)
		}
		if len(files) > 1 {
			printAnswer(flags, total, "total")
		}
	}
}

func getFlags() (flags setting, files []string) { //nolint:unparam // it's not done
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		// default settings
		flags = BytesFlag | LinesFlag | WordsFlag

		return flags, files // files = nil
	}

	arg := strings.Split(arguments[0], "")

	log.Println(arg) // TODO: удалить комментарий ниже и исправить

	// if arg[0] == "-" && len(arg) > 1 { // Check flags
	//	for _, flag := range arg[1:] {
	//		switch flag {
	//		case "l":
	//			flags = flags | LinesFlag
	//		case "w":
	//			flags = flags | WordsFlag
	//		case "m":
	//			flags.symbolsFlag = true
	//			flags.bytesFlag = false // if m, !c
	//		case "L":
	//			flags = flags | LenLinesFlag
	//		case "c":
	//			if flags.symbolsFlag == true { // if m, !c
	//				flags.bytesFlag = false
	//				break
	//			}
	//			flags.bytesFlag = true
	//		default: // Exit, if no flags
	//			fmt.Println("wc: illegal option --", flag)
	//			fmt.Println("usage: wc [-Lclmw] [file ...]")
	//			os.Exit(0)
	//		}
	//	}
	//
	//	files = arguments[1:]
	//	if len(files) == 0 {
	//		files = nil
	//	}
	//
	//	return flags, files
	// }

	// default settings
	flags = BytesFlag | LinesFlag | WordsFlag

	files = arguments
	if len(files) == 0 {
		files = nil
	}

	return flags, files
}

func getParams() (settings, []string) {
	var flags settings
	var files []string

	arguments := os.Args[1:]

	if len(arguments) == 0 {
		// default settings
		flags.linesFlag = true
		flags.wordsFlag = true
		flags.bytesFlag = true

		return flags, files // files = nil
	}

	arg := strings.Split(arguments[0], "")

	if arg[0] == "-" && len(arg) > 1 { // Check flags
		for _, flag := range arg[1:] {
			switch flag {
			case "l":
				flags.linesFlag = true
			case "w":
				flags.wordsFlag = true
			case "m":
				flags.symbolsFlag = true
				flags.bytesFlag = false // if m, !c
			case "L":
				flags.lenLineFlag = true
			case "c":
				if flags.symbolsFlag { // if m, !c
					flags.bytesFlag = false
					break
				}
				flags.bytesFlag = true
			default: // Exit, if no flags
				fmt.Println("wc: illegal option --", flag)   //nolint:forbidigo // it's console app
				fmt.Println("usage: wc [-Lclmw] [file ...]") //nolint:forbidigo // it's console app
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

func calcFromStdin() answer {
	// 世界 qwer
	// 1 2 12 8 7
	// l-w-c--m-L-

	var result answer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result.lines++ // Count lines

		words := strings.Split(line, " ")
		result.words += len(words) // Count words

		bytes := []byte(line)
		result.bytes += len(bytes) + 1 // Count bytes + /n

		runes := []rune(line)
		result.symbols += len(runes) + 1 // Count symbols + /n

		if result.lenLine < len(runes) {
			result.lenLine = len(runes) // Count len
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return result
}

func calcFromFile(pathFile string) answer {
	var result answer
	file, err := os.Open(pathFile)
	if err != nil {
		result.warning = "wc: " + pathFile + ": open: No such file or directory\n"
		return result
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	reader := bufio.NewReader(file)
	for {
		var line string
		line, err = reader.ReadString('\n')
		result.lines++ // Count lines

		words := strings.Split(line, " ")
		for _, v := range words {
			if v == "" || v == "\n" {
				continue
			}
			result.words++ // Count words
		}

		bytes := []byte(line)
		result.bytes += len(bytes) // Count bytes

		runes := []rune(line)
		result.symbols += len(runes) // Count symbols

		if result.lenLine < len(runes) {
			result.lenLine = len(runes) // Count len
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Println(err)
			return result
		}
	}

	result.lines--

	return result
}

func printAnswer(flags settings, value answer, label string) {
	if value.warning != "" {
		fmt.Printf("%s", value.warning) //nolint:forbidigo // it's console app
		return
	}

	if flags.linesFlag {
		fmt.Printf("%8d", value.lines) //nolint:forbidigo // it's console app
	}
	if flags.wordsFlag {
		fmt.Printf("%8d", value.words) //nolint:forbidigo // it's console app
	}
	if flags.bytesFlag {
		fmt.Printf("%8d", value.bytes) //nolint:forbidigo // it's console app
	}
	if flags.symbolsFlag {
		fmt.Printf("%8d", value.symbols) //nolint:forbidigo // it's console app
	}
	if flags.lenLineFlag {
		fmt.Printf("%8d", value.lenLine) //nolint:forbidigo // it's console app
	}
	if label != "" {
		fmt.Printf(" %s", label) //nolint:forbidigo // it's console app
	}

	fmt.Printf("\n") //nolint:forbidigo // it's console app
}
