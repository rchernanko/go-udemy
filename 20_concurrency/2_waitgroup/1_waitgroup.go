package main

import (
	"fmt"
	"runtime"
)

func main() {
	//1) There are some really useful functions within Go's runtime package:
	fmt.Println("OS\t", runtime.GOOS)                                       //On my work laptop: darwin
	fmt.Println("ARCH\t", runtime.GOARCH)                                   //On my work laptop: amd64
	fmt.Println("Number of CPUs on this device\t", runtime.NumCPU())        //On my work laptop: 8
	fmt.Println("Number of existing go routines\t", runtime.NumGoroutine()) //On my work laptop: 1

	foo()
	bar()

	fmt.Println("Number of CPUs on this device\t", runtime.NumCPU())        //On my work laptop: 8
	fmt.Println("Number of existing go routines\t", runtime.NumGoroutine()) //AFTER launching 'foo' in go routine: 2
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

/*

- So when we first run this (without any go routines/waitgroup), we get the following printed to the console.
i.e. foo runs and then bar

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

*/
