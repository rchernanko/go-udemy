package main

import "fmt"

/*

write a program that
- launches 10 goroutines
	- each goroutine adds 10 numbers to a channel
- pull the numbers off the channel and print them

*/

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-ch)
	}

	fmt.Println("about to exit")
}
