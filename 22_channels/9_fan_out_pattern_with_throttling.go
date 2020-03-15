package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*

- Let's have a look at the fan out pattern
- Let's say we have a piece of work that needs to be done over and over and over again, we can 'fan that out' to many
go routines, as opposed to running everything sequentially. This way, the piece of work can be run in parallel and can be
completed quicker (than running sequentially).
- And we can also use throttling to e.g. only fan out to 10 go routines at a time

*/

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	start := time.Now()

	go populate(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("took: ", time.Since(start))
	fmt.Println("about to exit")
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)

		//This is fanning out! Basically fanning out the work to multiple go routines! Awesome!
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)

		/*
			You could use throttling with the above so that we only fan out to e.g. 10 go routines instead of 100
			See fanOutInWithThrottling function below

		*/
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)

}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutInWithThrottling(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}
