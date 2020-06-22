package main

import "fmt"

/*

This exercise will reinforce our understanding of method sets:

- create a type person struct
- attach a method speak to type person using a pointer receiver
	- *person
- create a type human interface
	- to implicitly implement the interface, a human must have the speak method
- create func “saySomething”
	- have it take in a human as a parameter
 	- have it call the speak method
- show the following in your code
	- you CAN pass a value of type *person into saySomething
	- you CANNOT pass a value of type person into saySomething
- here is a hint if you need some help
	- https://play.golang.org/p/FAwcQbNtMG


*/

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old\n", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "richard",
		age:  33,
	}

	/*
		Even though p1 is not a pointer to a person, the below will still work
		See notes in `method_sets_revisited`
	*/

	p1.speak()

	//But with interfaces, this will not work.
	//saySomething(p1)

	//The thing we pass in MUST be a pointer to a person as defined in the person's `Speak` method
	saySomething(&p1)

}
