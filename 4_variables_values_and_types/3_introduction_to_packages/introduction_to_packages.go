package main

import "fmt"

/*
	https://golang.org/pkg/

	Useful details of what each package in the Go standard library does / is used for

	When you click on a package in the above, if you click on the 'index' link, it then takes you to a nice list of the
	available functions in this package

	You can also do this at https://godoc.org (e.g. https://godoc.org/fmt) and Todd prefers the layout of this

*/

func main() {
	/*

	One thing that you can see when reviewing the docs for fmt.Println is that it's signature allows for variadic
	parameters to be passed in

	func Println(a ...interface{}) (n int, err error)

	So we c an do stuff like this:

	 */

	strings := []string{"richard", "learning", "go"}

	fmt.Println("hello", 34, false, strings)

	/*

	After running go run <file_name>...

	hello 34 false [richard learning go]

	To use the Println's returns:

	 */

	n, err := fmt.Println("returns")

	fmt.Println(n) //prints 8 (number of bytes written)

	if err == nil {
		fmt.Println(err) //prints <nil>
		fmt.Println("no errors in the Println execution") //gets printed
	}

	//To tell the compiler that you don't care / don't want to use the 'err' return:
	n, _ = fmt.Println("not needed the err")

	fmt.Println(n) //19


/*

Key points covered in this video:

- variadic parameters
   	- the “...<some type>” is how we signify a variadic parameter
   	- the type “interface{}” is the empty interface
		- every value is also of type “interface{}”
   	- so the parameter “...interface{}” means “give me as many arguments of any type as you’d like

- throwing away returns
   	- use the “_” underscore character to throw away returns

- you can’t have unused variables in your code
	- this is code pollution
   	- the compiler doesn’t allow it

- we use this notation in Go
   	- package.Identifier
	- ex: fmt.Println()
   		- we would read that: “from package fmt I am using the Println func”
   	- an identifier is the name of the variable, constant, func

- packages
   	- code that is already written which you can use
   	- imports

 */
}
