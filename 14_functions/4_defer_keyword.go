package main

import "fmt"

/*
	Defer is a keyword
	https://golang.org/ref/spec#Defer_statements


*/

func main() {
	defer foo4()
	bar4()

	/*

		The above would print out

		bar
		foo

		the execution of foo4 is deferred right until the end of the main function. The last thing before the final '}'
		right before the func 'main' exits

		e.g. you might use it when we open a file, we can defer the closing of it. As soon as you open it, you can
		defer to close it - is nice and readible

	*/

}

func foo4() {
	fmt.Println("foo")
}

func bar4() {
	fmt.Println("bar")
}
