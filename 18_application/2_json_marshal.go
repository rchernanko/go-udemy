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

type Club struct {
	Name           string `json:"team_name"`
	BestPlayer     string `json:"best_player"`
	FoundedIn      int    `json:"founded,omitempty"`
	IgnoreMeInJson string `json:"-"`
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

	//Another example - we can use annotations to control what appears in the json:

	club := Club{
		Name:           "Tottenham Hotspur",
		BestPlayer:     "Harry Kane",
		FoundedIn:      0,
		IgnoreMeInJson: "whatever",
	}

	fmt.Printf("Club Before Marshalling: %v\n", club)
	bs, err = json.Marshal(club)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(bs)) //{"team_name":"Tottenham Hotspur","best_player":"Harry Kane"}
}
