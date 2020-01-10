package main

import "fmt"

/*
	Defer is a keyword
	https://golang.org/ref/spec#Defer_statements

	And from https://golang.org/doc/effective_go.html#defer:

	"Go's defer statement schedules a function call (the deferred function) to be run immediately before the function
	executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must
	be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or
	closing a file."

*/

func main() {
	defer foo4()
	bar4()

	/*

		The above would print out

		bar
		foo

		the execution of foo4 is deferred right until the end of the function executing the defer returns
		(in our case, the main function). The last thing before the final '}' right before the func 'main' exits

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
