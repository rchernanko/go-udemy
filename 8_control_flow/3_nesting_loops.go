package main

import "fmt"

func main() {

	//Loop within a loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("The output loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}
}
