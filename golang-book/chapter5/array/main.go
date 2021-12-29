package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 100

	var total int = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / len(x))

	y := [5]float64{11.11, 22.22, 33.33, 44.44, 555.5}

	var result float64 = 0
	for _, value := range y {
		result += value
	}
	fmt.Println(result / float64(len(y)))
}
