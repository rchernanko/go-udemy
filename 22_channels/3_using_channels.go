package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)
	//the above used to be a go function, but we removed the keyword "go" so that the execution of main.go wouldn't
	//complete until the info from the channel had been received

	fmt.Println("about to exit")
}

//send func
func foo(cs chan<- int) {
	cs <- 42
}

//receive func
func bar(cr <-chan int) {
	fmt.Println(<-cr)
}