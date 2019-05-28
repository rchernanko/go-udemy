package main

import "fmt"

func main() {

	/*

		- https://golang.org/ref/spec#For_statements (For statements with range clause)
		- "A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map,
		or values received on a channel."

	*/

	z := []int{4, 3, 2, 5, 8, 1}

	/*

		RANDOM DIGRESSION

		- using len and cap functions
		- retrieve element at a certain position in a slice

	*/

	fmt.Printf("len: %v\n", len(z))
	fmt.Printf("cap: %v\n", cap(z))
	fmt.Printf("full slice: %v\n", z)
	fmt.Printf("element at index position 3 is: %v\n", z[3])

	/*

		- NOW BACK TO THE RANGE

	*/

	for index, value := range z {
		fmt.Println(index, value)
	}
}
