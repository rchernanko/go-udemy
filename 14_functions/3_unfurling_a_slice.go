package main

import "fmt"

/*

	Unfurling - Todd's term

*/

func main() {
	xs := []int{2, 3, 4, 5, 6, 7, 8, 9}

	/*

		foo3(xs)

		We cannot do the above because the function signature requires an unlimited number of ints AND we are trying to pass
		in a SLICE of ints - 2 different types

		HOWEVER, we can add ... to the end of xs to UNFURL the slice of ints and pass them in one by one as single ints

		Language Spec - https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	*/

	sum := foo3(xs...)
	fmt.Println("the total is:", sum)
}

func foo3(x ...int) int {
	fmt.Printf("%T\n", x) //prints []int

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
