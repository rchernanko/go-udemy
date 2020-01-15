package main

import "fmt"

func main() {

	func() {
		fmt.Println("i am an anonymous function")
	}()

	func(word string) {
		fmt.Printf("i am using %s", word)
	}("something")

}
