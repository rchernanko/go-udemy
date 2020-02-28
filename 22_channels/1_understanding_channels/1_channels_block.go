package main

import "fmt"

/*

- Channels are a way of synchronising your concurrent code when its running in parallel
- They are a way of co-ordinating the efforts of all code running in parallel, i.e. all the go routines

*/

func main() {
	//create a channel onto which I put integers
	c := make(chan int)

	//Put the int 42 on that channel
	c <- 42

	//Read the integer from that channel
	fmt.Println(<-c)

	/*

		When I actually try to run the above, it fails. I get the error:

		`fatal error: all goroutines are asleep - deadlock!` on line 17 (the one that tries to put 42 on the channel)

		And this is because CHANNELS BLOCK.

		When you send and receive on a channel, it's like relay racers in a track race that have to pass the batton exactly on time.
		Here, the putting of 42 on the channel cannot occur until the send (put on the channel) and receive (take from the channel)
		can happen at the same time. The channel blocks the send until the receive is ready, which is what's happened
		above. The receive is not ready. When we try to put on the channel, on line 17, there is nothing set up to receive
		from that channel.

		Let's fix this in the next file in this directory

	*/
}
