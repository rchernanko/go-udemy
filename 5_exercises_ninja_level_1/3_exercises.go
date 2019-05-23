package main

import (
	"fmt"
)

var a = 42
var b = "James Bond"
var c = true

func main() {
	s := fmt.Sprintf("%d, %s, %t", a, b, c)
	fmt.Println(s)

	/*

	I can actually also just use %v in the above:
	s := fmt.Sprintf("%v, %v, %v", a, b, c)

	I can also add tabs between each value:
	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)

	 */
}