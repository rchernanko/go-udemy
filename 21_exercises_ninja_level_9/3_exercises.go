package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

- Using goroutines, create an incrementer program
	- have a variable to hold the incrementer value
	- launch a bunch of goroutines
		- each goroutine should
			- read the incrementer value
				- store it in a new variable
			- yield the processor with runtime.Gosched()
			- increment the new variable
			- write the value in the new variable back to the incrementer variable
- use waitgroups to wait for all of your goroutines to finish
- the above will create a race condition.
- Prove that it is a race condition by using the -race flag
- if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

*/

var counter = 0
var wg1 sync.WaitGroup

func main() {
	var numOfGoRoutines = 100
	wg1.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go incrementCounter()
	}

	wg1.Wait()
	fmt.Println(counter)
}

func incrementCounter() {
	updatedCounter := counter
	runtime.Gosched()
	updatedCounter++
	counter = updatedCounter
	wg1.Done()
}
