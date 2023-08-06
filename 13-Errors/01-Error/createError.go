package main

import (
	"errors"
	"fmt"
)

func main() {
	output, error := CreateErrorMethod1(0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(output)
	}

	output, error = CreateErrorMethod2(10)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(output)
}

func CreateErrorMethod1(number int) (int, error) {
	if number == 0 {
		return 0, errors.New("Number is not valid")
	}
	return number, nil
}

func CreateErrorMethod2(number int) (int, error) {
	if number == 0 {
		return 0, fmt.Errorf("Number is not valid: %d", number)
	}
	return number, nil
}
