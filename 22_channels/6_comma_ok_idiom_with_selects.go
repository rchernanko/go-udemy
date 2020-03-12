package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send1(even, odd, quit)

	receive1(even, odd, quit)

	fmt.Println("about to exit")
}

func send1(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive1(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case i, ok := <-quit:

			/*
				You can use the comma, ok idiom to check when a channel is closed.
				If the channel is open, ok will be true.
				If the channel is closed, ok will be false AND the value in i will be the default value for that type
			*/
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			} else {
				fmt.Println("from comma ok:", i)
			}
		}
	}
}
