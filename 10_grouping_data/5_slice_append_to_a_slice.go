package main

import "fmt"

func main() {

	/*

		- Appending to a slice.
		- 'append' is a built-in function - https://godoc.org/builtin
		- More info here - https://golang.org/doc/effective_go.html#append
		- Interesting bit in the above link is that you can append a slice to a slice like this:

				x := []int{1,2,3}
				y := []int{4,5,6}
				x = append(x, y...)
				fmt.Println(x)

		- The above would print out [1 2 3 4 5 6]

	*/

	x := []int{3, 4, 5, 6, 7}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 100)
	fmt.Println(x) //prints [3 4 5 6 7 77 88 99 100]

	y := []int{234, 1123, 45}
	x = append(x, y...)
	//In the above, we 'unfurl' (a term Todd has coined) all the values in y. I.E. take all the values out of y and put them in the append function

	/*

		IMPORTANT:

		'...' BEFORE something in the definition of a function e.g. func append(slice []T, elements ...T)

			- This means that the function takes an unlimited number the T type.
			- This is a variadic parameter

		'...' AFTER something e.g. 'x = append(x, y...)'

			- Take an aggregate data structure that has a whole bunch of values in it and puts them in the e.g. append function
			- Is like doing this:

				x = append(x, 234, 1123, 45)


	*/

}
