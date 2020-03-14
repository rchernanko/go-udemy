package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		//when using a range, we must close the channel (always on the sending side) otherwise we get a `fatal error: all goroutines are asleep - deadlock!` exception
		//With selects, I don't think we always need to close the channel
	}()

	/*
		Range over the channel to receive the values. This is a little design pattern when using channels.

		The range will continue to poll the channel until the channel is closed.
		When the channel is closed, it ranges over any last values on that channel and processes them.
		And when there are no more values on the closed channel, it leaves the range block
	*/
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
