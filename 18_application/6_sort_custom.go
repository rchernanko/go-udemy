package main

import (
	"fmt"
	"sort"
)

type personSort struct {
	first string
	age   int
}

func main() {
	p1 := personSort{"James", 32}
	p2 := personSort{"Moneypenny", 27}
	p3 := personSort{"Q", 64}
	p4 := personSort{"M", 56}

	people := []personSort{p1, p2, p3, p4}
	fmt.Println("Not sorted by age:", people)

	//2) And after making the changes described below...
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age:", people)

	//3) I will now also create some functions below that allow me to sort by name:
	sort.Sort(ByName(people))
	fmt.Println("Sorted by name:", people)

	/*

		1)

		Sorting a slice of struct...

		What if we want to sort all the above values by age or by name?
		Let's take a look at the documentation here - https://godoc.org/sort#example-package
		And now we will add a "ByAge" type and associated methods that implement the Sort interface - see below
		Then in our main function we can call the sort function and pass in ByAge(people) - this should sort our people by age

	*/
}

type ByAge []personSort

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

type ByName []personSort

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].first < a[j].first
}