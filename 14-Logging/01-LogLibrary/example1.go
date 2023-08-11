package main

import (
	"log"
	"os"
)

var (
	errorLogger *log.Logger
	warnLogger *log.Logger
	infoLogger *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt",  os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file ", err)
	}
	flags := log.Ldate | log.Ltime | log.Lshortfile;
	log.SetFlags(flags)
	log.SetOutput(file)

	errorLogger = log.New(file, "Error: ", flags)
	warnLogger = log.New(file, "Warn: ", flags)
	infoLogger = log.New(file, "Info: ", flags)
}

func main() {
	infoLogger.Println("Start to main ")
	sum(1,9)
}

func sum(a, b int) {
	warnLogger.Println("start sum")
	println(a +b)
	warnLogger.Println("end sum sum")
}