package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Fix the race condition you created in exercise #4 by using package atomic
*/

var counter2 int64
var wg3 sync.WaitGroup

func main() {
	var numOfGoRoutines = 100
	wg3.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go incrementCounter2()
	}

	wg3.Wait()
	fmt.Println("final counter:", counter2)
}

func incrementCounter2() {
	atomic.AddInt64(&counter2, 1)
	fmt.Println(atomic.LoadInt64(&counter2))
	wg3.Done()
}
