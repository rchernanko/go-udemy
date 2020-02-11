package main

import (
	"fmt"
	"sort"
)

/*

We are going to go on another adventure!

We will take a slice and sort it.

*/

func main() {
	xi := []int{10, 4, 73, 3, 84, 22}
	xs := []string{"richard", "james", "hallie", "jan", "anne"}

	fmt.Println("Unsorted int slice:", xi)
	fmt.Println("Unsorted string slice", xs)

	//How do we sort the values in the slices above?
	//Let's use the sort package - https://godoc.org/sort

	//ints - https://godoc.org/sort#example-Ints

	sort.Ints(xi)
	fmt.Println("Sorted int slice: ", xi)

	//strings - https://godoc.org/sort#Strings

	sort.Strings(xs)
	fmt.Println("Sorted string slice: ", xs)

	//Interesting, I randomly went off on a quick digression re: searching for strings
	index := sort.StringSlice(xs).Search("jan")
	fmt.Println(index)
	fmt.Println(len(xs))
	if index == len(xs) {
		fmt.Println("item not present in slice of strings")
	}

	//or...
	index = sort.SearchStrings(xs, "richard")
	fmt.Println(index)
	//Which does the same as above StringSlice.Search...
}
