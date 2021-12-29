package main

import (
	"fmt"
	"math"
)

// structs
type Circle struct {
	// fields
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

// embedded types
type Android struct {
	Person
	Model string
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleArea2(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleArea3(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// methods
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area2() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is", p.Name)
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	// struct
	c := Circle{0, 0, 5}
	fmt.Println(c)
	// fields
	fmt.Println(c.x, c.y, c.r)

	fmt.Println(circleArea2(c))

	// methods
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area2())

	// embedded
	a := new(Android)
	a.Talk()
}
