package main

import "fmt"

var a int

type hotdog int
var b hotdog

func main() {
	a = 42
	b = 51

	fmt.Println(a) //42
	fmt.Printf("%T\n", a) //int

	fmt.Println(b) //51
	fmt.Printf("%T\n", b) //main.hotdog

	/*

	CONVERSION - we take a value of a certain type and convert it to another type

	In the last tutorial we saw that we could not do this:

	a = b

	because b is of type hotdog and a is of type int

	However, we know that hotdog's underlying type is int, so we can use conversion here:

	*/

	a = int(b)

	fmt.Println(a) //51
	fmt.Printf("%T\n", a) //int

	//And with another custom type (where the underlying type is a string)...

	var ottos = "ottos"
	fmt.Println(ottos) //ottos

	type burger string
	var otherBurger burger
	otherBurger = "maccas"

	ottos = string(otherBurger)
	fmt.Println(ottos) //maccas

	/*

	This is called CONVERSTION in golang. NOT CASTING. You won't find any mention of casting in any of the golang docs

	Key points from this video (and the last one):

		- You can create your own types
		- You can convert to the underlying type

	Next up, some exercises!!

	 */

}
