package main

import "fmt"

func main() {
	c := make(chan int)

	//send (and close the channel)
	go func(){
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	/*
	receive
	The range will keep looping over a channel until the channel is closed
	 */
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
