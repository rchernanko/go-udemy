package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 42

	/*

		Below, I am trying to put on a second value to the buffered channel
		But, because of the way we've declared the buffered channel at the beginning, it can only hold 1 value (that's not yet been received) at a time

		So when I try to put on a second value (before the first has been received), it blocks, and we get the error below:

		`fatal error: all goroutines are asleep - deadlock!`

		The way to fix this is to create a buffered channel with 2 items...
		And you would then need to add another fmt.Println(<-c) to print out the 43 too

		But when using buffered channels, you of course need to know what your absolute limit is...which can sometimes
		be quite difficult to know...

	*/

	c <- 43

	fmt.Println(<-c)
}
