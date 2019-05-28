package main

import "fmt"

/*

1)

- In this 'grouping data' section, we're learning how to get a bunch of values (of the same type) into one data structure
and have one data structure hold it.
- This one data structure is called an 'AGGREGATE DATA TYPE', e.g. an array, a list

2)

- FYI - Another 'grouped' data type is called the 'COMPOSITE DATA TYPE'
- This is something that allows different data types to be composed together
- In Go, we have a 'type' called a 'struct' which allows us to compose values of different types together
- I.e. a struct can group together ints, strings, bools etc
- A struct is a composite data type

*/

func main() {

	/*

		- We will now create a slice using a 'composite literal'
		- Composite literal = A way where we group together values

		x := type{values}     // composite literal

	*/

	y := []int{4, 5, 6, 7}
	fmt.Println(y)

	/*

		From language specification - https://golang.org/ref/spec#Composite_literals

		"Composite literals construct values for structs, arrays, slices, and maps and create a new value each time
		they are evaluated. They consist of the type of the literal followed by a brace-bound list of elements."

	*/
}
