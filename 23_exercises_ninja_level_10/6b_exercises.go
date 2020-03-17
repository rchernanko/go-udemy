package main

import "fmt"

/*

write a program that:

- puts 100 numbers to a channel
- pull the numbers off the channel and print them

Me - could do this a few ways, below uses a select statement. Previous file used a range loop

*/

func main() {
	c := make(chan int)
	q := make(chan bool)

	go send4(c, q)
	receive4(c, q)
}

func send4(c chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	q <- true
}

func receive4(c <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			fmt.Println("returning now")
			return
		}
	}
}
