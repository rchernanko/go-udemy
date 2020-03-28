package main

import "fmt"

func main() {
	fmt.Println("need this for go build i think")
}

/**

Went through a lot of this - just watched the videos

Go modules:

- Dependency management
- Create a go.mod file with the `go mod init` command
- I've created a new repo for this - just to have it a little cleaner. More notes in there
- https://github.com/rchernanko/go-modules-messing-around

Go Commands:

- go fmt ./... (formatting all files when run from root directory)

Difference between go build and go install...

- go build will generate a binary in my current location BUT NOT in my gopath > bin directory
- go install will generate a binary and put it in my GOPATH's bin directory
	- note that when I ran a go install from within this project, becasue this project sits outside the GOPATH on my laptop, I needed to do:
		- export GOBIN=$GOPATH/bin
		- go install
			- without the export, i got an error message:
				`go install: no install location for directory /Users/richardchernanko/Development/go-udemy/3_development_environment outside GOPATH`

		- important - "Go Modules is somewhat of a departure from this approach. You’re no longer required to keep all your projects under $GOPATH."
			- brilliant article - https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31

go run:
	- needs a file name, eg, go run main.go
	- go run <file name>
	- go run *.go

go build:
	- for an executable:
		- builds the file
		- reports errors, if any
		- if there are no errors, it puts an executable into the current folder
	- for a package:
		- builds the file
		- reports errors, if any
		- throws away binary

go install:
	- for an executable:
		- compiles the program (builds it)
		- names the executable
			- mac: the folder name holding the code
			- windows: file name
		- puts the executable in workspace / bin
			- $GOPATH / bin
	- for a package:
		- compiles the package (builds it)
		- puts the executable in workspace / pkg
			- $GOPATH / pkg
		- makes it an archive file

*/
