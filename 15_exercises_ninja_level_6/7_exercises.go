package main

import "fmt"

var g func()

func main() {
	hi := func() {
		fmt.Println("printing out hi!")
	}

	hi()

	g = hi
	g()
}
