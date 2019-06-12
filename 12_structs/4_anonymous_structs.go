package main

import "fmt"

func main() {

	//Here's an anonymous struct

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   32,
	}

	fmt.Println(p1) //prints {james bond 32}
}
