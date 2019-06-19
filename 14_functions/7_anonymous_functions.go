package main

import "fmt"

func main() {

	func() {
		fmt.Println("anonymous func ran")
	}() //these empty brackets are the arguments being passed in to the anonymous function's signature

	func(name string) {
		fmt.Println("anonymous func printing name:" + name)
	}("richard")
}
