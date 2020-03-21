package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "welcome my dear James" {
		t.Error("got", s, "expected", "welcome my dear James")
	}
}

/*
- I didn't really watch through the examples videos properly - TODO actually watch through them - https://blog.golang.org/examples

Basically they are tests and used with go docs...

Godoc examples are snippets of Go code that are displayed as package documentation and that are verified by running them
as tests. They can also be run by a user visiting the godoc web page for the package and clicking the associated "Run" button.

*/
func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// welcome my dear James
}

/*
Here is my benchmark test
This will run x number of times to assess how quickly the Greet function runs
*/
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}

/*

go test -bench .
OR
go test -bench Greet

The output is something like this:

richardchernanko > 5_benchmarking > (master) $ go test -bench Greet
goos: darwin
goarch: amd64
BenchmarkGreet-8         8563302               131 ns/op

- Ran on 8 cores
- It ran 8563302 times
- It took 131 nanoseconds per operation

*/
