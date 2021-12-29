package main

import (
	"fmt"
	"math"
)

// struct
type Circle struct {
	x, y, r float64
}

// interface
type Shape interface {
	area() float64
}
type MultiShape struct {
	shapes []Shape
	// multiShape := MultiShape{
	// 	shapes: []Shape{
	// 		Circle{0,0,5},
	// 	},
	// }
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.area()
	}
	return total
}

func (m *MultiShape) area2() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	x := MultiShape{}
	fmt.Println(x.area2())
}
