package main

import "fmt"

type myOwnType int
var d myOwnType

func main() {
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	d = 42
	fmt.Println(d)

	//Documentation on "underlying types" - https://golang.org/ref/spec#Types
}
