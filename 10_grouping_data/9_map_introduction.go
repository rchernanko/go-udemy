package main

import "fmt"

func main() {

	/*

		- Maps are 'key/value' store
		- They are unordered lists
		- https://golang.org/doc/effective_go.html#maps
		- FYI from above - "Like slices, maps hold references to an underlying data structure. If you pass a map to a
			function that changes the contents of the map, the changes will be visible in the caller."

	*/

	m := map[string]int{ //This is the composite literal
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m) //prints map[James:32 Miss Moneypenny:27]

	//ACCESS value by key
	fmt.Println(m["James"]) //prints 32

	//ACCESSING a value where the key does not exist
	fmt.Println(m["Barnabas"])
	//prints 0. i.e. if you enter a key and that value is not stored in the map, the zero value for that type is returned
	//but there is a better way to check whether the key exists (and to not just trust the zero value) - as below

	//the below technique is called the 'comma ok idiom' in golang - check effective go

	if v, ok := m["Barnabas"]; !ok {
		fmt.Println("Barnabas does not exist in the map!")
		fmt.Println(v) //prints 0
	}

	//some extra playing around
	m2 := map[string]int{}
	fmt.Println(m2) //prints map[]
}
