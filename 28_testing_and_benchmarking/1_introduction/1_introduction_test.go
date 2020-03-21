package main

//mmm weird becasue i thought this needed to be main_test.
//see notes below. something not quite right perhaps in my gopath...
//anyway, todd actually does the tutorial like this though... (with just main, instead of main_test)

import (
	"testing"
)

//very strange, when i run this with the GoLand IDE, the MySum function is undefined. It might have something to do with my go path...
// but running it on the command line with a `go test`, it works TODO - come back to this though...
func TestMySum(t *testing.T) {
	i := mySum(2, 3) //note that i am testing a non exportable function because i am in the package level scope
	if i != 5 {
		t.Error("Expected", 5, "Got", i)
	}
}
