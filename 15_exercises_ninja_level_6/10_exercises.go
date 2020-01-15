package main

import "fmt"

func main() {
	h := bumperUpper()
	fmt.Println(h()) //1
	fmt.Println(h()) //2
	fmt.Println(h()) //3
	fmt.Println(h()) //4
	fmt.Println(h()) //5
	fmt.Println(h()) //6

	g := bumperUpper()
	fmt.Println(g()) //1
	fmt.Println(g()) //2
	fmt.Println(g()) //3
	fmt.Println(g()) //4
	fmt.Println(g()) //5
	fmt.Println(g()) //6
}

func bumperUpper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
