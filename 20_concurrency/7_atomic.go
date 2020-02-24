package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*

- In the previous tutorial, we used mutex to control go routine access to a shared variable
- We can actually achieve the same thing by using the atomic package (a sub package within the sync package)
- We will keep it simple in this tutorial...it's just a quick demo

*/

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	var counter int64 //you must use an int64 to use the 'atomic' package. Every time you see an int64, think of that package

	const numOfGoRoutines = 100
	var wg sync.WaitGroup
	wg.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)                         //write to my counter
			runtime.Gosched()                                    //allows other go routines to run - "go ahead and run something else"
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) //read from my counter
			/*

				The above is a 'safe' way to access the counter variable
				Running this code with a race condition check results in no race conditions reported
					- go run --race 7_atomic.go`
			*/
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}
