package main

import "fmt"

var outside = "I am outside"

//outsideWillNotCompile := "I won't compile" //THIS WILL NOT WORK

//Declare there is a variable with the identifier 'z' and that 'z' is of type int. Assigns the zero value of type int to 'z'
var z int

/*
	If I try to `go run` here, I will see:

	# command-line-arguments
	./the_var_keyword.go:7_zero_value:1: syntax error: non-declaration statement outside function body

*/

func main() {

	//we saw in the last tutorial that this is the way to use the short declaration operator
	x := 42 //declare a variable and assign a value (of a certain type)
	fmt.Println(x)

	//now, let's use the var keyword
	var y = 43
	fmt.Println(y)

	/*

	what's the difference between short declaration operator and var?

	- I can declare a variable with the 'var' keyword OUTSIDE this function (and then e.g. print it)
	- But I can't do that with a short declaration operator.
	- E.g. outside + outsideWillNotCompile variables

	 */

	fmt.Println(outside) //will print I am outside
	//fmt.Println(outsideWillNotCompile) //this won't compile

	//Best practice is to use the short declaration operator as much as possible to limit the scope of variables

	foo()
}

func foo() {
	//Note - the scope of 'outside' and 'z' is 'package level'
	fmt.Println("again:", outside)
	fmt.Println(z) //prints the 'nil' value of 'int' (which is 0)
}

/*

RE: 'z' above...

Zero values (from the language spec) - https://golang.org/ref/spec#The_zero_value

"
When storage is allocated for a variable, either through a declaration or a call of new, or when a new value is created,
either through a composite literal or a call of make, and no explicit initialization is provided, the variable or value
is given a default value. Each element of such a variable or value is set to the zero value for its type: false for
booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed
if no value is specified.

These two simple declarations are equivalent:

var i int
var i int = 0
After

type T struct { i int; f float64; next *T }
t := new(T)
the following holds:

t.i == 0
t.f == 0.0
t.next == nil

The same would also be true after:

var t T

"






*/