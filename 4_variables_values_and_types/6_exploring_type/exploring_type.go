package main

import "fmt"

/**

- Go is all about type

 */

var y = 42
var z = "shaken not stirred" //this is an "interpreted" string literal. Differs to "raw string literal" (see below)
//FYI - can also do this `var z string = "shaken not stirred"`

func main() {
	fmt.Println(y) //42
	fmt.Printf("%T\n", y) //int

	fmt.Println(z)
	fmt.Printf("%T\n", z) //string

	/*

	Go is a static programming language. So I cannot now do this:

	z = 43

	I will get the following error:

	./exploring_type.go:21:4: cannot use 43 (type int) as type string in assignment

	z will ONLY hold values of the TYPE int. This cannot change. It's not a dynamic programming language like JavaScript

	A variable is declared to hold a value of a certain type

	If you look at the language spec, you can also declare a 'raw string literal' using the ``

	E.g.

	 */

	rawStringLiteral := `James said, "shaken, not stirred"`

	fmt.Println(rawStringLiteral) //prints James said, "shaken, not stirred"

	rawStringLiteral2 := `James said, 

	"shaken,
	not stirred"`

	fmt.Println(rawStringLiteral2)

	/*

	prints:

	James said,

	        "shaken,
	        not stirred"

	i.e. also keeps the spacing, line breaks etc


	NOTE - one more thing:

	- "composite" data types refers to types such as array, slice, map types that can be used to group together primitive data types
		-  aka structure or aggregate data type
		- e.g. a slice of strings


	 */

}