package logger

import (
	"fmt"
	"os"
)

type FileLogger struct {
	filePath string
}

func NewFileLogger(filePath string) FileLogger {
	return FileLogger{
		filePath: filePath,
	}
}

func (l FileLogger) Info(msg string) {
	err := l.writeFile(fmt.Sprintf("[INFO] %s", msg))
	if err != nil {
		fmt.Println(err)
	}
}

func (l FileLogger) Warn(msg string) {
	err := l.writeFile(fmt.Sprintf("[WARN] %s", msg))
	if err != nil {
		fmt.Println(err)
	}
}

func (l FileLogger) Error(msg string) {
	err := l.writeFile(fmt.Sprintf("[ERROR] %s", msg))
	if err != nil {
		fmt.Println(err)
	}
}

func (l FileLogger) writeFile(msg string) error {
	f, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("opening file error: %s", err)
	}
	defer f.Close()

	if _, err = f.WriteString(msg + "\n"); err != nil {
		return fmt.Errorf("writing message error: %s", err)
	}
	return nil
}
