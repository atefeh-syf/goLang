package main

import "fmt"

type number interface {
	int | float64
}

func main() {
	slice1 := []int{1,2,3,4,5}
	fmt.Printf("%d\n",  Sum(slice1))

	slice2 := []float64{1.5, 2.5, 3.5, 4.5, 5.5}
	fmt.Printf("%f\n",  Sum(slice2))
}

func Sum[T number](slc []T) (sum T) {
	for _,  v := range slc {
		sum += v
	}
	return
}
