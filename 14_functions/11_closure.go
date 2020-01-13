package main

import "fmt"

//Closure = enclosing a variable in some code so that the variable scope is limited to that one area of the code

func main() {

	x := 1
	fmt.Println(x)

	//Code block within a code block! y's scope is limited to within the block below
	{
		y := 42
		fmt.Println(y)
	}

	a := incrementor()
	b := incrementor()
	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(b()) //1
	fmt.Println(b()) //2

}

func incrementor() func() int {
	var x int
	return func() int { //this function is a closure around the 'x' variable. I can never directly access 'x'
		x++
		return x
	}
}
