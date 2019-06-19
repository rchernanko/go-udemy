package main

import "fmt"

func main() {

	/*

		A 'callback' is passing a func as an argument
		When you do this, this is referred to as 'functional programming'

	*/

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...) //unfurling
	fmt.Println("all numbers", s)

	//And now let's call the function that has a function defined as a parameter.
	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	//And again for the odd numbers
	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

/*
	And now to define a function which has a function in its signature
	The function we are specifying in the signature is 'f func(xi ...int) int'
*/

func even(f func(xi ...int) int, vi ...int) int {
	var evenNums []int
	for _, v := range vi {
		if v%2 == 0 {
			evenNums = append(evenNums, v)
		}
	}

	return f(evenNums...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var oddNums []int

	for _, v := range vi {
		if v%2 != 0 {
			oddNums = append(oddNums, v)
		}
	}

	return f(oddNums...)
}
