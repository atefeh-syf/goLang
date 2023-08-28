package main

import (
	"Unbuffered/examples"
	"time"
)

func main() {
	examples.Emaple1()
}

func IntroExample() {
	numChannel := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		numChannel <- 1
	}()
	println("waiting for numChannel", time.Now().String())
	receivedNum := <-numChannel
	println(receivedNum)
}
