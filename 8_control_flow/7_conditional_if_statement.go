package main

import "fmt"

func main() {

	//Let's go through a few basic examples:

	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	fmt.Println("done with basic examples")

	// More interestingly, we can have an IF statement with an initialisation
	fmt.Println("IF statement with an initialisation")

	if j := 42; j == 42 {
		fmt.Printf("J is %v\n", j)
	}
}
