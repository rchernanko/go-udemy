package main

import (
	"encoding/json"
	"fmt"
)

//intentionally leaving everything here as lower case
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "richard",
		last:  "chernanko",
		age:   33,
	}

	p2 := person{
		first: "lucy",
		last:  "chernanko",
		age:   30,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	//Now let's try to marshal the above people slice

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(bs)) // everything ran BUT because the fields are in lower case, '[{},{}]' is printed

	/*

		IF we were to make the person fields with Capitals, we would get:

		[{"First":"richard","Last":"chernanko","Age":33},{"First":"lucy","Last":"chernanko","Age":30}]

	*/

	//TODO up to 5 20 in video
}
