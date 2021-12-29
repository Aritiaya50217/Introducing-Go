package main

import "fmt"

func average(xs []float64) float64 {
	// panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

func main() {
	number := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(number))

	x, y := f()
	fmt.Println(x, y)
}
