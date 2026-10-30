package logger

import "fmt"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}
