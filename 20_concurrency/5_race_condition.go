package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

- Let's create a race condition

*/

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	counter := 0 //shared variable
	var wg sync.WaitGroup
	const numOfGoRoutines = 100
	wg.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			v := counter
			runtime.Gosched() //allows other go routines to run - "go ahead and run something else"
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("Counter", counter)

	/*

			- The counter is never 100
			- It's always different. This is because we have a race condition above
			- There are up to 100 go routines accessing and writing to the counter variable and it's never necessarily in the right state

			- There's actually a go command that you can run to identify race conditions - is brilliant!:
			- `go run --race <file to run>`

		go run --race 5_race_condition.go
		CPUs: 8
		Go Routines: 1
		==================
		WARNING: DATA RACE
		Read at 0x00c0000b2030 by goroutine 8:
		  main.main.func1()
		      /Users/richardchernanko/Development/go-udemy/20_concurrency/5_race_condition.go:26 +0x3c

		Previous write at 0x00c0000b2030 by goroutine 7:
		  main.main.func1()
		      /Users/richardchernanko/Development/go-udemy/20_concurrency/5_race_condition.go:29 +0x5c

		Goroutine 8 (running) created at:
		  main.main()
		      /Users/richardchernanko/Development/go-udemy/20_concurrency/5_race_condition.go:25 +0x239

		Goroutine 7 (finished) created at:
		  main.main()
		      /Users/richardchernanko/Development/go-udemy/20_concurrency/5_race_condition.go:25 +0x239
		==================
		Go Routines: 1
		Counter 100
		Found 1 data race(s)

		- So how to fix the race condition?
		- We can use Mutex - let's that in the next tutorial

	*/
}
