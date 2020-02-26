package main

import (
	"fmt"
	"runtime"
)

/*

Create a program that prints out your OS and ARCH. Use the following commands to run it
	- go run
	- go build
	- go install

*/
func main() {
	fmt.Println("Arch is:", runtime.GOARCH)
	fmt.Println("OS is:", runtime.GOOS)
}

/*

- go run

$ go run 6_exercises.go
Arch is: amd64
OS is: darwin

- go build

$ go build 6_exercises.go
Builds the binary which we can then run with ./6_exercises and this prints out
Arch is: amd64
OS is: darwin

- go install

	- should add the binary to my gopath > bin
	- and from then anywhere, you can just run 6_exercises and the binary will execute
	- but when i tried to do a go install from here, i couldn't - got the following error message:

		go install: no install location for directory /Users/richardchernanko/Development/go-udemy/21_exercises_ninja_level_9 outside GOPATH
        For more details see: 'go help gopath'

	TODO - should probably watch the installing go videos at the beginning

*/
