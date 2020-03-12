package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	//the for loop is an infinite for loop (until we get out of it via a value being sent to the quit channel)
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- 0
	close(q)

	/*
		Note re: why don't i need to close the e and o channels.
		"You needn't close every channel when you've finished with it. It's only necessary to close a channel when it
		is important to tell the receiving goroutines that all data have been sent. A channel that the garbage collector
		determines to be unreachable will have its resources reclaimed whether or not it is closed."

		https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open

		https://gobyexample.com/closing-channels

		"Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers."
	*/
}
