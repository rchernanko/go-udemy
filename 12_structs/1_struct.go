package main

import "fmt"

/*

	- A struct is a data structure which allows us to compose together values of different types.
	- So it's an aggregate data type. It aggregates to gather values of different type.

*/

type person struct { //Mix of int and strings - We're composing together values of a different type.
	first string
	last  string
	age   int
}

func main() {

	//Here, we are CREATING a VALUE of TYPE 'PERSON'
	p1 := person{
		first: "James",
		last:  "bond",
		age:   32,
	}

	p2 := person{
		first: "miss",
		last:  "moneypenny",
		age:   45,
	}

	fmt.Println(p1) // {James bond 32}
	fmt.Println(p2) // {miss moneypenny 45}

	fmt.Println(p1.last)  // bond
	fmt.Println(p2.first) // miss

}
