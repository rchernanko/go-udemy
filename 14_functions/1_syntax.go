package main

import "fmt"

func main() {

	/*
		Syntax is:

			func (r receiver) identifier(parameter(s)) (return(s)) { ... }

			fyi - the receiver is optional

		IMPORTANT: by default, everything in Go is PASS BY VALUE

	*/

	foo()
	bar("richard")
	s1 := woo("james")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)

}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("hello", s)
}

func woo(st string) string {
	return fmt.Sprint("hello from woo, ", st)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true

	return a, b
}
