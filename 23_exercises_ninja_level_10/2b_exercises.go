package main

import (
	"fmt"
)

/*

Get this code running:

package main

import (
	"fmt"
)

func main() {
	cr := make(<-chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}


*/

func main() {
	cr := make(chan int) //make the channel bi-directional instead of receive only

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
