package main

import "fmt"

/**

- We saw the BREAK keyword in the last tutorial
- We will look at the CONTINUE keyword now

*/

func main() {

	/*

		First of all, he shows us how to get the remainder of a division calculation

		So a normal division would look like this:

	*/

	f := 83 / 40
	fmt.Println(f) //2

	//And to get the remainder, we use the percentage symbol...this is called the 'modulo' operator
	g := 83 % 40
	fmt.Println(g) //3
	fmt.Println("done")

	/*

		Now, back to the for loop with a BREAK and CONDITION clause.

		So now, we only want to print even numbers

	*/

	y := 0

	for {

		y++
		if y > 100 {
			break
		}

		if y%2 != 0 {
			continue
		}

		fmt.Println(y)
	}
	fmt.Printf("done\n")

}
