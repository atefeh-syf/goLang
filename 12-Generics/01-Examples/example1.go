package main

import "fmt"

type number interface {
	int | int16 | int32 | int64 | float32 | float64
}

func main() {
	x := Sum(2, 7)
	fmt.Printf("%d\n", x)

	y := Sum(2.3, 3.9)
	fmt.Printf("%f\n", y)
}

func Sum[T number](a, b T) T {
	return a + b
}
