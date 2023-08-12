package main

import (
	//"log"
	"fmt"
	"os"
	"runtime/debug"
)

func main() {

	defer func() {
		fmt.Println("Defer of main")

		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Recovering error ... ")
			debug.PrintStack()
			os.Exit(1)
		}
	}()

	numbers := []int{1, 2, 3, 4, 5}
	function1(numbers, 2)
	Divide(10, 0)
}

func function1(numbers []int, index int) {
	defer func() {
		fmt.Println("defer of function1")
	}()
	// 1 - index out of range
	//fmt.Println(numbers[index])
	// 2 - panic
	//panic("panic")
	// 3 - log
	//log.Fatal("Fatal")
	//fmt.Println(12)
}

func Divide(a int, b int) {
	defer func() {
		fmt.Println("defer of divide")
	}()

	fmt.Printf("Start of divide: %d , %d\n", a, b)
	fmt.Printf("Result: %d\n", a/b)
	fmt.Printf("End of divide: %d , %d\n", a, b)

}
