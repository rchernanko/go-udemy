package main

import "testing"

func TestMySum3(t *testing.T) {

	type test struct {
		testName string
		data     []int
		answer   int
	}

	tests := []test{
		{
			testName: "5 + 2",
			data:     []int{5, 2},
			answer:   3,
		},
		{
			testName: "12 + 13",
			data:     []int{12, 13},
			answer:   7,
		},
		{
			testName: "45 + 3",
			data:     []int{45, 3},
			answer:   48,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			i := mySum3(test.data...)
			if i != test.answer {
				t.Error("Expected", test.answer, "Got", i) //if there's an error, it's a bit nicer output because we now print the name of the test that failed...
			}
		})
	}
}
