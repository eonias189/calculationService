package logger

import "fmt"

func Info(msgs ...any) {
	fmt.Print("INFO ")
	fmt.Println(msgs...)
}

func Error(err error) {
	fmt.Println("ERROR", err.Error())
}
