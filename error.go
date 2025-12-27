package logger

import "fmt"

func Error(msg string) {
	fmt.Println("[ERROR] " + msg)
}
