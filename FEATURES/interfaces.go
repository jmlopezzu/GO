package main

import "fmt"

// Shape interface
type Shape interface {
	Area() float64
}

// Circle type implementing Shape interface
type Circle struct {
	Radius float64
}

// Rectangle type implementing Shape interface
type Rectangle struct {
	Width, Height float64
}

// Implementing Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implementing Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	printArea(circle)
	printArea(rectangle)
}
