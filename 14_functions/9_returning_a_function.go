package main

import "fmt"

func main() {

	//Returning a string - basic example
	s1 := fooHere()
	fmt.Println(s1)

	//Now to return a function
	s2 := barHere()
	fmt.Printf("%T\n", s2) // Prints 'func() int'

	//And now to run the 'func() int'
	s3 := s2()
	fmt.Println(s3)   // Prints '451'
	fmt.Println(s2()) ///You can also just do this (and it will also print '451')

	//In fact, you can even just do this
	fmt.Println(barHere()()) //this will also print '451'

}

func fooHere() string {
	s := "hello world"
	return s
}

func barHere() func() int {
	return func() int { //returns this function
		return 451
	}
}

//When you pass in a func as an argument, this is what is known as a 'callback' - we will look at this in the next tutorial
