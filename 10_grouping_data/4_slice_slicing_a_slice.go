package main

import "fmt"

func main() {

	/*

		- We can use the colon operator to slice a slice

	*/

	x := []int{4, 5, 6, 7, 42}
	fmt.Println(x[1]) //5
	fmt.Println(x)    //[4, 5, 6, 7, 42]

	//And now to slice the slice...

	fmt.Println(x[1:])  //gives me from position 1 onwards i.e. prints [5 6 7 42]
	fmt.Println(x[1:3]) //gives me from position 1 until (and excluding) i.e. prints [5 6]
}
