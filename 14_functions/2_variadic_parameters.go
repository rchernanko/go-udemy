package main

import "fmt"

func main() {
	sum := foo2(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the total is:", sum)
}

func foo2(x ...int) int {
	/*
		 	The ...int above is a variadic parameter - this allows me to pass in an unlimited number of 'ints' (including none)
			NOTE that as per the language specification, a variadic parameter must be the last parameter in a function signature

				i.e. you can't do this:

				foo2(x ...int, s string) { ... }

			Language Spec - https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	*/

	fmt.Println(x)        // prints [2 3 4 5 6 7 8 9]
	fmt.Printf("%T\n", x) //prints []int

	sum := 0
	for i, v := range x {
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total")
		sum += v
	}
	return sum
}
