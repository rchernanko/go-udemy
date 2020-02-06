package main

import "fmt"

func main() {
	p1 := person{
		name: "Richard",
		age:  33,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	p.age = 34
	p.name = "James"

	/*

		Can also do:

		(*p).age = 34
		(*p).name = "James"

	*/
}
