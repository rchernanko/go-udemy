package main

import "fmt"

func main() {

	for d := 0; d <= 20; d++ {

		if d%2 != 0 {
			fmt.Println(d) //Print the odd numbers
		}
	}
}
