package main

import "fmt"

/*

- Let's fix the blocking channel issue from the previous file
- There are actually 2 ways to do this. Here is the first

*/

func main() {
	c := make(chan int)

	/*

		Wrap the putting of 42 on the channel in a go routine.
		When we run main.go, there will now be 2 x go routines (the main one and then this one below).

		In the main go routine, the channel will first be created, the second go routine will be kicked off and the execution
		will then continue so that the channel is receiving messages via the `fmt.Println(<-c)`

		Separately, the second go routine will be running in parallel, and as soon as the channel is ready to receive, it will
		put the int 42 onto the channel

	*/
	go func() {
		//Put the int 42 on that channel
		c <- 42
	}()

	//Read the integer from that channel
	fmt.Println(<-c)
}
