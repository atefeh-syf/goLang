package main

import (
	"fmt"
	"time"
)

func main() {
	numChannel := make(chan int)
	go SendDataToChannel(numChannel)

	receivedNum := <-numChannel
	PrintLnWithTime("Receive number :", receivedNum)
	time.Sleep(time.Second * 2)
}

func SendDataToChannel(numChannel chan int) {
	PrintLnWithTime("Before Sending 1 to channel")
	numChannel <- 1 // channelName
	PrintLnWithTime("After Sending 1 to channel")

	PrintLnWithTime("Before Sending 2 to channel")
	numChannel <- 2
	PrintLnWithTime("After Sending 2 to channel")

	PrintLnWithTime("Before Sending 3 to channel")
	numChannel <- 3
	PrintLnWithTime("After Sending 3 to channel")
}

func PrintLnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
