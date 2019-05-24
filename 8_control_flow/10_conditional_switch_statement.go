package main

import "fmt"

func main() {

	//Switch statement with no fallthrough keyword

	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print2")
	case 3 == 3:
		fmt.Println("prints")
	case 4 == 4:
		fmt.Println("also true, does it print?")
	}

	//Interestingly enough, only the 3rd case prints. By default, the switch statement doesn't fall through (i.e. once
	//the 3rd case has executed, the 4th case is not executed). You can however change this by using the 'fallthrough' keyword

	//TODO up to 2 mins 15
}
