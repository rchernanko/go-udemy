package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

- In addition to the main goroutine, launch two additional goroutines
	- each additional goroutine should print something out
	- use waitgroups to make sure each goroutine finishes before your program exists

*/

var wg sync.WaitGroup

func main() {

	fmt.Println("Begin CPUs:", runtime.NumCPU())
	fmt.Println("Begin go routines:", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("Mid CPUs:", runtime.NumCPU())
	fmt.Println("Mid go routines:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End CPUs:", runtime.NumCPU())
	fmt.Println("End go routines:", runtime.NumGoroutine())
	fmt.Println("finished")
}

func foo() {
	fmt.Println("Hello I am foo")
	wg.Done()
}

func bar() {
	fmt.Println("Hello I am bar")
	wg.Done()
}
