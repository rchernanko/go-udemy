package main

import "fmt"

type hotdog int
var e hotdog
var f int

func main() {
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	e = 42
	fmt.Println(e)

	f = int(e) //convert e which is of type hotdog to an int
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
