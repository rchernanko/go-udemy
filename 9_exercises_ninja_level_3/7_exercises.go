package main

import "fmt"

func main() {

	for r := 1; r < 10; r++ {
		if r == 1 {
			fmt.Println("printing out the if statement")
		} else if r%2 == 0 {
			fmt.Println("printing the even numbers")
		} else {
			fmt.Println("printing out the odd numbers")
		}
	}
}
