package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	// defer moves the call to second to the end of the function
	defer second()
	first()

}
