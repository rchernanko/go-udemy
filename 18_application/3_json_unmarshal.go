package main

import (
	"encoding/json"
	"fmt"
)

type club struct {
	Name           string `json:"team_name"`
	BestPlayer     string `json:"best_player"`
	FoundedIn      int    `json:"founded,omitempty"`
	IgnoreMeInJson string `json:"-"`
}

func main() {
	club1 := club{
		Name:           "Tottenham Hotspur",
		BestPlayer:     "Harry Kane",
		FoundedIn:      0,
		IgnoreMeInJson: "whatever",
	}

	fmt.Printf("Club Before Marshalling: %v\n", club1)
	bs, err := json.Marshal(club1)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(bs))

	//OK now let's unmarshal the 'bs' back to a club struct

	club2 := &club{}
	err = json.Unmarshal(bs, club2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*club2)

	//Great website he shows to convert json to go code - https://mholt.github.io/json-to-go/ (in my bookmarks now)

	//We could also unmarshall with a raw string:
	s := `[{"team_name":"Tottenham Hotspur","best_player":"Harry Kane"},{"team_name":"Manchester United","best_player":"Paul Pogba"}]`
	bs = []byte(s)
	fmt.Printf("%T\n", s)  //string
	fmt.Printf("%T\n", bs) //[]uint8 - NOTE that as a byte is an alias for uint8

	var clubs []club
	err = json.Unmarshal(bs, &clubs)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(clubs) // [{Tottenham Hotspur Harry Kane 0 } {Manchester United Paul Pogba 0 }]

}
