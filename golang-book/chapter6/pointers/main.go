package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	// x = xPtr
	zero(&x)
	fmt.Println(x) // 0

	// new
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
