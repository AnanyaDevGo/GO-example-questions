package main

import "fmt"

type Shape interface {
	Area() float32
}
type Rectangle struct {
	width  float32
	height float32
}
type Circle struct {
	radius float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}
func (c Circle) Area() float32 {
	return 2 * 3.14 * c.radius
}
func main() {
	rect := Rectangle{width: 5, height: 2}
	circ := Circle{radius: 5}
	shape := []Shape{rect, circ}
	for _, shape := range shape {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
