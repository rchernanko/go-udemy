package main

import "fmt"

/*

write a program that:

- puts 100 numbers to a channel
- pull the numbers off the channel and print them

Me - could do this a few ways, below uses a range loop. Next file will use a select

*/

func main() {
	c := make(chan int)

	go send3(c)
	receive3(c)

	fmt.Println("about to exit")
}

func send3(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive3(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
