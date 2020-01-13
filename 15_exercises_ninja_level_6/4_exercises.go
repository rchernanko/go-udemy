package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is: %s %s and I am %d years old", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Richard",
		last:  "Chernanko",
		age:   33,
	}

	p1.speak()
}
