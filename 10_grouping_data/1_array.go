package main

import "fmt"

func main() {

	//CREATE AN ARRAY

	var x [5]int   // We MUST specify the number of elements the array has + each element must be of the same TYPE
	fmt.Println(x) // [0 0 0 0 0]

	// ASSIGNING VALUES TO ELEMENTS
	x[3] = 6
	fmt.Println(x) // [0 0 0 6 0]

	//x[5] = 321    //This won't compile (because 5 is out of bounds for a 5-element array

	// GET THE SIZE IOF THE ARRAY
	fmt.Println(len(x)) // 5

	/*

		- Arrays are seen as 'building blocks' in go
		- WE DO NOT REALLY USE ARRAYS IN GO. INSTEAD WE USE SLICES
		- States this in effective go - https://golang.org/doc/effective_go.html#arrays


	*/
}
