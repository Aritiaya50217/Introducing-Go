package main

import (
	"fmt"
)

func main() {
	// int
	fmt.Println("1 + 1 =", 1+1)

	// float
	fmt.Println("1.11 + 1.11 =", 1.11+1.11)

	// string
	s := "String"
	t := "test"

	fmt.Printf("%#v\n", s)
	fmt.Println("Lenght : ", len(s))
	fmt.Println("Character 1 :", s[0])
	fmt.Printf("%v"+"%v\n", t, s)

	// booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
