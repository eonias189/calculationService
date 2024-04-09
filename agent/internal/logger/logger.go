package logger

import (
	"fmt"
)

func Info(msgs ...any) {
	fmt.Println(msgs...)
}

func Error(error error) {
	fmt.Println("Error:", error.Error())
}
