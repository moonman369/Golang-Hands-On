package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func calcArea(sh Shape) float64 {
	return sh.area()
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func main() {
	c1 := Circle{
		radius: (1 / math.Pi),
	}

	r1 := Rectangle{
		length: 98.2343,
		width:  10,
	}

	fmt.Println("Area of circle, c1 =", calcArea(c1))
	fmt.Println("Area of rectangle, r1 =", calcArea(r1))
}
