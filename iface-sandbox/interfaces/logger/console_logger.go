package logger

import (
	"fmt"

	"github.com/fatih/color"
)

type ConsoleLogger struct{}

func NewConsoleLogger() ConsoleLogger {
	return ConsoleLogger{}
}

func (l ConsoleLogger) Info(msg string) {
	blue := color.New(color.FgHiBlue).SprintfFunc()
	fmt.Printf("%s %s\n", blue("[INFO]"), msg)
}

func (l ConsoleLogger) Warn(msg string) {
	blue := color.New(color.FgHiYellow).SprintfFunc()
	fmt.Printf("%s %s\n", blue("[WARN]"), msg)
}

func (l ConsoleLogger) Error(msg string) {
	blue := color.New(color.FgHiRed).SprintfFunc()
	fmt.Printf("%s %s\n", blue("[ERROR]"), msg)
}
