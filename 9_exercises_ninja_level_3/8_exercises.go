package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("won't come in here")
	case 21 == 45:
		fmt.Println("won't come in here either")
	case 25 == 25:
		fmt.Println("should come in here")
	default:
		fmt.Println("won't come in here")
	}
}
