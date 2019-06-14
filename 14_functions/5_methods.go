package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		licenseToKill: true,
	}

	sa1.speak()
}

//Let's add a method to secret agent
func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last)
}
