package main

import "fmt"

func main() {

	/*

		- How to add something to a map + range over it

	*/

	m := map[string]int{
		"James": 35,
	}

	m["Richard"] = 32

	for k, v := range m {
		fmt.Println(k, v)
	}
}
