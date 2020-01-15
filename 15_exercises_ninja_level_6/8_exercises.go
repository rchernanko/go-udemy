package main

import "fmt"

func main() {
	variable1 := returnAFunc("tomorrow")
	hi := variable1()
	fmt.Println(hi)
}

func returnAFunc(word string) func() string {
	return func() string {
		return fmt.Sprintf("saying this word: %s", word)
	}
}
