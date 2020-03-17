package main

/*

Starting with the code below, pull the values off the channel using a select statement

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}

*/

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen1(q)

	receive1(c, q)

	fmt.Println("about to exit")
}

func gen1(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
	}()

	return c
}

func receive1(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("from the c channel:", v)
		case <-q:
			fmt.Println("returning")
			return
		}
	}
}
