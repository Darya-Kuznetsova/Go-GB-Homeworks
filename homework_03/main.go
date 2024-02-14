package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle := Rectangle{width: 7, height: 11, Shape: Shape{}}
	fmt.Println("Rectangle area:", rectangle.Area())
	circle := Circle{radius: 4, Shape: Shape{}}
	fmt.Println("Circle area", circle.Area())
}

type Shape struct {
}

type Rectangle struct {
	width  float64
	height float64
	Shape
}

type Circle struct {
	radius float64
	Shape
}

func (s Shape) Area() float64 {
	return 0.0
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
