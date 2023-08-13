package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	value := 0
	go func1()
	go func2()
	go func3()

	go func() {
		value++
	}()

	go func() {
		value += 2
	}()

	go func() {
		value += 3
	}()

	println(value)

	time.Sleep(time.Second)
}

func func1() {
	println("func1")
}

func func2() {
	println("func2")
}

func func3() {
	println("func3")
}
