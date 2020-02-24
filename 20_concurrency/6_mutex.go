package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

- Let's use the race condition from the previous tutorial
- The reason we have a race condition is because there are multiple go routines accessing the shared 'counter' variable
- What we need to happen is that when one of the go routines accesses the variable, it 'checks it out' (like taking a book out
from the library). And when it's 'checked out', no other go routines can access it until it's checked back in. Nobody else
can access that book until we take it back to the library. Once it's checked back in, other go routines can then access the
variable.
Mutex locks access to a variable / certain block of code.

*/

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	counter := 0 //shared variable
	var wg sync.WaitGroup
	const numOfGoRoutines = 100
	wg.Add(numOfGoRoutines)

	var mu sync.Mutex //add a mutex here

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			mu.Lock() //lock (check out)
			v := counter
			runtime.Gosched() //allows other go routines to run - "go ahead and run something else"
			v++
			counter = v
			mu.Unlock() //unlock (check in)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}
