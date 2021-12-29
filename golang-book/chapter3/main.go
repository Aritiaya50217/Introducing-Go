package main

import "fmt"

var z string = "World"

func main() {
	var x string = "Hello"
	fmt.Println(x)

	y := "Hello"
	fmt.Println(y)

	// scope
	fmt.Println(z)

	// constants
	const i string = "constants"
	// i = "Some other string" ---> cannot assign to i
	fmt.Println(i)

	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
