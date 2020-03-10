package main

import "fmt"

func main() {

	/*
	RECEIVE ONLY channel
	By adding the `<-` before the chan, this means that I only receive from this channel. I cannot send to it
	*/
	cr := make(<- chan int, 2)

	//cr <- 34 `invalid operation: c <- 34 (send to receive-only type <-chan int)`

	go func(){
		fmt.Println(<-cr)
	}()

	fmt.Printf("%T\n", cr) //prints `<-chan int`
}
