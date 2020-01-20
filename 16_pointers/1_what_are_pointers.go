package main

import "fmt"

/*

- A pointer is something that points to a location in memory where a value is stored

 */


func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //when we use the '&' operator, it prints out the address in memory where the value '42' is stored

	//"Here is the value AAAANNNDDDD (&&&&&) here is the address"

	//Now let's print out the types of a and &a
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", &a)
	// The above will print "*int" - the type of &a (the address) is a 'pointer to an int'.
	// This type is 'pointing to the memory location of where an int is stored'

	//*int is its own type, so you can't do:
	//var int b = a

	//If I want to see the value stored at a memory address, I can use the * operator:
	b := &a
	fmt.Println(*b) //aka 'de-referencing an address' (i.e. gives me the value stored at that address) - prints 42

	//You can even do this - "give me the memory address for this value and then de-reference it":
	fmt.Println(*&a) //prints 42

	//I can even change the value of b (which will in turn change the value of a because of the assignment of &a to b above):
	*b = 45
	fmt.Println(a)
}
