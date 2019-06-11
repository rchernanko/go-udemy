package main

import "fmt"

func main() {

	/*

		- When you create a slice, as well as using a composite literal, you can also use the built in function 'make'
		- Todd says though that you should only really do this when you know for sure the final size of the slice's underlying array

	*/

	x := make([]int, 10, 100) //here, we are allocating values to the first 10 indexes of the slice
	fmt.Println(x)
	fmt.Println(len(x)) //prints 10
	fmt.Println(cap(x)) //the underlying array has a capacity of 100 elements

	/*

		However, I can't do this:
		x[10] = 23 	// this results in a panic: runtime error: index out of range

	*/

	//Instead, we can do this:
	x = append(x, 23)
	fmt.Println(x)
	fmt.Println(len(x)) //prints 11
	fmt.Println(cap(x)) //prints [0 0 0 0 0 0 0 0 0 0 23]

	/*

		What happens when you've added 100 elements to the slice though?
		Remember, 100 relates to the size of the slice's underlying array

		Well, when we add the 101st value, the old array was discarded, a new one was created, and it is now doubled in size
		So the size of new underlying array would now be 200

	*/

}
