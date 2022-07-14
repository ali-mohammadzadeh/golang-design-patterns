package main

func main() {
	logger := getLoggerInstance()
	logger.log("my log")
	logger1 := getLoggerInstance()
	logger1.log("my logger1")
	logger2 := getLoggerInstance()
	logger2.log("my logger2")
	logger3 := getLoggerInstance()
	logger3.log("my logger3")
}
