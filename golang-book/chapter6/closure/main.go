package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	a := 0
	increment := func() int {
		a++
		return a
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// call function
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

}
