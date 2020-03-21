package main

//INITIAL PROBLEMS

//mmm weird because i thought this needed to be main_test.
//see notes below. something not quite right perhaps in my gopath...
//anyway, todd actually does the tutorial like this though... (with just main, instead of main_test)

//OK FOUND THE ISSUE re: above:

//For a start, I think pretty much 100% of packages in this repo are 'main'. Think that's causing issues. Also with the issues described below re: goland ide
//
//https://stackoverflow.com/questions/19998250/proper-package-naming-for-testing-with-the-go-language
//If I want to test 'private', non exportable functions, my package must not end in _test,
//If I want to just check the exportable functions, the package must end in _test

import (
	"testing"
)

//initial issues described below...
//very strange, when i run this with the GoLand IDE, the MySum function is undefined. It might have something to do with my go path...
// but running it on the command line with a `go test`, it works
func TestMySum(t *testing.T) {
	i := mySum(2, 3) //note that i am testing a non exportable function because i am in the package level scope
	if i != 5 {
		t.Error("Expected", 5, "Got", i)
	}
}
