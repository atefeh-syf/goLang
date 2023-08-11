package main

import (
	"fmt"

	"golang.org/x/text/number"
)

func main() {
		numbers := []int{1,2,3,4,5}
		function1(numbers, 6)
}

func function1(numbers []int, index int) {
		fmt.Println(numbers[index])
}