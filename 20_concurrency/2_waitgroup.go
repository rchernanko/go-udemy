package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//Start at the bottom comments first...

	//There are some really useful functions within Go's runtime package:
	fmt.Println("OS\t", runtime.GOOS)                                       //On my work laptop: darwin
	fmt.Println("ARCH\t", runtime.GOARCH)                                   //On my work laptop: amd64
	fmt.Println("Number of CPUs on this device\t", runtime.NumCPU())        //On my work laptop: 8
	fmt.Println("Number of existing go routines\t", runtime.NumGoroutine()) //On my work laptop: 1

	//Adding a wait group - 'there is 1 thing we should wait for before finishing the execution'
	wg.Add(1)

	//Now let's launch 'foo' into a go routine.
	go foo()
	bar()

	fmt.Println("Number of CPUs on this device\t", runtime.NumCPU())        //On my work laptop: 8
	fmt.Println("Number of existing go routines\t", runtime.NumGoroutine()) //AFTER launching 'foo' in go routine: 2
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() //remove that 1 thing we were waiting for
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

/*

- So when we first run this (without the go routines/waitgroup), we get the following printed to the console

OS       darwin
ARCH     amd64
Number of CPUs on this device    8
Number of existing go routines   1

foo: 0
foo: 1
foo: 2
foo: 3
foo: 4
foo: 5
foo: 6
foo: 7
foo: 8
foo: 9
bar: 0
bar: 1
bar: 2
bar: 3
bar: 4
bar: 5
bar: 6
bar: 7
bar: 8
bar: 9

Number of CPUs on this device    8
Number of existing go routines   1

- And after launching foo into a go routine:

OS       darwin
ARCH     amd64
Number of CPUs on this device    8
Number of existing go routines   1
bar: 0
bar: 1
bar: 2
bar: 3
bar: 4
bar: 5
bar: 6
bar: 7
bar: 8
bar: 9
Number of CPUs on this device    8
Number of existing go routines   2

- Why isn't foo executing?
- Well we need to add some sort of synchronisation here to tell go not to leave the main function without executing the foo go routine
- We will use a waitgroup to do this - godoc.org/sync. let's add this now
- After we add the wait group:

OS       darwin
ARCH     amd64
Number of CPUs on this device    8
Number of existing go routines   1
bar: 0
bar: 1
bar: 2
bar: 3
bar: 4
bar: 5
bar: 6
bar: 7
bar: 8
bar: 9
Number of CPUs on this device    8
Number of existing go routines   2
foo: 0
foo: 1
foo: 2
foo: 3
foo: 4
foo: 5
foo: 6
foo: 7
foo: 8
foo: 9

*/
