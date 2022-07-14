package main

import "fmt"

type Logger struct {
	level int `default:1`
}

func (logger Logger) log(message string) {
	fmt.Println(message)
}

var loggerInstance *Logger

func getLoggerInstance() *Logger {
	if loggerInstance == nil {
		fmt.Println("create new instance")
		loggerInstance = &Logger{}
	}
	fmt.Println("return   instance")
	return loggerInstance
}
