package main

/*

Get this code working using a buffered channel:

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

*/

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
