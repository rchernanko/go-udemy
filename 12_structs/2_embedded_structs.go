package main

import (
	"fmt"
)

type person1 struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person1
	licenseToKill bool
}

func main() {

	p1 := secretAgent{
		person1: person1{
			first: "james",
			last:  "bond",
			age:   34,
		},
		licenseToKill: true,
	}

	fmt.Println(p1) //{{james bond 34} true}

	//The 'inner type' gets promoted to the 'outer type' - i.e. i can access the fields of person1 with the dot notation
	//Here, the 'inner type' is person1 and the 'outer type' is secretAgent
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)
	fmt.Println(p1.licenseToKill)

	//I.E you don't have to do this:
	fmt.Println(p1.person1.first)
	fmt.Println(p1.person1.last)
	fmt.Println(p1.person1.age)
}
