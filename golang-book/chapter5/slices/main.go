package main

import "fmt"

func main() {
	// empty slices
	var x []float64
	fmt.Printf("%#v\n", x)

	// จอง memory ให้ slices
	y := make([]float64, 5, 10)
	fmt.Printf("%#v\n", y)

	// append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Printf("Before :%v After : %v\n", slice1, slice2)

	// copy
	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Printf("Slice 1 :%v\nSlice 3 :%v\n", slice1, slice3)
}
