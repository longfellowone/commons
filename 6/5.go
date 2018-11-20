package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * 2.0
}

type Square struct {
	length float64
	width  float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

func info(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	newC := Circle{
		radius: 2.0,
	}

	newS := Square{
		length: 100,
		width:  100,
	}

	info(newC)
	info(newS)
}
