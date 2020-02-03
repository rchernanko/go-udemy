package main

import (
	"fmt"
	"math"
)

/*

- When we have a type, we can attach methods to it
- Those methods attached to a type are known as its method sets

Non pointer vs pointer receivers:
- A non-pointer receiver
	- works with values that are pointers or non-pointers
- A pointer receiver
 	- only works with values that are pointers

Receivers		Values
------------------------
(t T)			T and *T
(t *T)			*T

 */

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 { //c is a non-pointer receiver  - so works with values that are pointers or non-pointers
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(c) //info method works with a non-pointer value

	c1 := &circle{4}
	info(c1) //info method works with a pointer value
}
