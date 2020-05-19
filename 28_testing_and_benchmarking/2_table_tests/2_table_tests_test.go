package main

import "testing"

func TestMySum2(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{
			data:   []int{5, 2},
			answer: 7,
		},
		{
			data:   []int{12, 13},
			answer: 25,
		},
		{
			data:   []int{45, 3},
			answer: 48,
		},
	}

	for _, v := range tests {
		i := mySum2(v.data...)
		if i != v.answer {
			t.Error("Expected", v.answer, "Got", i)
		}
	}
}
