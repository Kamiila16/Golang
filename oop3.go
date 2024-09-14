package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func PrintArea(s Shape) {
	fmt.Printf("The area is: %.2f\n", s.Area())
}
func main() {
	c := Circle{Radius: 10}
	r := Rectangle{width: 4, height: 5}
	PrintArea(c)
	PrintArea(r)
}
