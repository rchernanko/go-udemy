package main

import "fmt"

/**

Let's look at how to create our own type

"go is all about type"
"type is life"

Go is a static programming language so as soon as you declare that an e.g. variable is of a certain type, it will never change type

 */

var a int

type hotdog int //Here we have created our own type
var b hotdog //Here we have declared a variable of that above type

func main() {
	a = 42
	b = 51

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b) //52
	fmt.Printf("%T\n", b) //main.hotdog

	/*

	What happens if I try to assign the variable b to variable a?

	a = b

	I get an error message:

	"./creating_your_own_type.go:32:4: cannot use b (type hotdog) as type int in assignment"

	b is of type 'hotdog' and a is of type 'int'...so i cannot do the above. they are 2 different types

	however, there is a way around this...let's take a look at that in the next video

	 */
}