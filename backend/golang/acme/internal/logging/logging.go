package logging

import "fmt"

var L = &LoggerStdOut{}

type LoggerStdOut struct{}

func (l LoggerStdOut) Debug(message string, args ...any) {
	fmt.Printf("[DEBUG] "+message, args...)
}

func (l LoggerStdOut) Info(message string, args ...any) {
	fmt.Printf("[INFO] "+message, args...)
}

func (l LoggerStdOut) Warn(message string, args ...any) {
	fmt.Printf("[WARN] "+message, args...)
}

func (l LoggerStdOut) Error(message string, args ...any) {
	fmt.Printf("[ERROR] "+message, args...)
}
