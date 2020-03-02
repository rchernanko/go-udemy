package main

import "fmt"

/*

In the last tutorial, we created a bi-directional channel, that is, a channel that we could send information to, and also
receive information from.

But when creating a channel, you can also specify it to be one direction, e.g. you can only send to this channel. Or you can
only receive from it. Todd thinks that this could make your code more readable.

 */

func main() {

	/*
	SEND ONLY channel
	By adding the `<-` after the chan, this means that I only send info to this channel. I cannot receive from it
	 */
	cs := make(chan <- int, 1)

	go func(){
		cs <- 42
	}()

	//fmt.Println(<-c) `invalid operation: <-c (receive from send-only type chan<- int)`

	fmt.Printf("%T\n", cs) //prints `chan<- int`
}