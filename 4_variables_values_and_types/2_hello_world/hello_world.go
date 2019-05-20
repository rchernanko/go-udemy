package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("final thing before exiting")

}

func foo() {
	fmt.Println("I'm in foo")
}

/*

- The entry point of any program is 'package main' and 'func main'
- The exit from a program is when the flow of execution exits 'func main'
- In the above, he quickly wanted to demonstrate 'control-flow'

- To run this file:
	- navigate to the path of this file and run `go run hello_world`:

	Hello world
	I'm in foo
	something else
	0
	2
	4
	6
	...
	final thing before exiting

*/