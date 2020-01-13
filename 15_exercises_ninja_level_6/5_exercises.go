package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	area := s.area()
	fmt.Printf("shape area is: %v\n", area)
}

func main() {
	c := circle{
		radius: 12.345,
	}

	info(c)

	s := square{
		length: 15,
	}

	info(s)
}
