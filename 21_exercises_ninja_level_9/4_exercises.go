package main

import (
	"fmt"
	"sync"
)

/*
Fix the race condition you created in the previous exercise by using a mutex
- it makes sense to remove runtime.Gosched()
*/

var counter1 = 0
var wg2 sync.WaitGroup
var mu1 sync.Mutex

func main() {
	var numOfGoRoutines = 100
	wg2.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go incrementCounter1()
	}

	wg2.Wait()
	fmt.Println(counter1)
}

func incrementCounter1() {
	//wrap the lock / unlock around the race condition
	mu1.Lock()
	updatedCounter := counter1
	updatedCounter++
	counter1 = updatedCounter
	mu1.Unlock()
	wg2.Done()
}
