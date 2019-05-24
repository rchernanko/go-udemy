package main

import "fmt"

func main() {

	j := 42

	if j == 42 {
		fmt.Println("our value is 42")
	} else if j == 41 {
		fmt.Println("our value is 41")
	} else {
		fmt.Println("our value is not 41 or 42")
	}

}
