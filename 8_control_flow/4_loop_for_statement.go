package main

import "fmt"

/**

- In the video, he first jumps into the documentation - https://golang.org/ref/spec#For_statements
- There are 3 types of 'for' statements (as below)
- At the end of the video, he also checks out the 'effective go' documentation - https://golang.org/doc/effective_go.html#for

*/

func main() {

	//1 - for statement with single condition

	fmt.Println("1 - FOR statement with single condition")
	x := 1

	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Printf("done\n")

	//2 - for statements with a clause

	fmt.Println("2 - FOR statement with a clause")

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("done\n")

	//3 - for statements with a RANGE clause - we will look at this in a latter tutorial
	fmt.Println("3 - FOR statement with a range clause - TO BE COVERED IN LATTER TUTORIAL")

	//3a - Eternal for statement with an IF condition and a BREAK clause
	fmt.Println("3a - External FOR statement with an IF condition and a BREAK clause")

	//Interestingly, you can use 'for' on its own (without any condition)!
	y := 1
	for {
		if y > 10 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Printf("done\n")

}
