package main

import (
	"fmt"
	"sync"
)

/*

- Fanning out and fanning in are common concurrent design patterns
- I have a chunk of work, I don't know how much it's going to be. Let's fan that out to as many go routines as possible.
They'll all be working on that as much as they can, and then when they get results, we'll fan those results back in to
another channel and we'll get a channel with just the results

Fanning In - Taking values from many channels and putting those values onto one channel

*/

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send2(even, odd)
	go receive2(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send2(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive2(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even { //the code will block here until the channel is closed (see explanation below)
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd { //the code will block here until the channel is closed (see explanation below)
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}

/*

https://imil.net/blog/2018/12/31/Understanding-golang-channel-range/

"we must understand that a channel range doesn’t stop when there’s no more elements to be read unless the channel is closed"

*/
