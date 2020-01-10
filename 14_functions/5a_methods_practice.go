package main

import "fmt"

type player struct {
	FirstName string
	LastName  string
	Age       int
	Team      string
}

func (pl player) stateTeam() {
	fmt.Printf("I am %s %s and I play for %s\n", pl.FirstName, pl.LastName, pl.Team)
}

func (pl player) stateAge() {
	fmt.Printf("I am %s %s and I am %d years old\n", pl.FirstName, pl.LastName, pl.Age)
}

func main() {

	pl1 := player{
		FirstName: "Harry",
		LastName:  "Kane",
		Age:       25,
		Team:      "Tottenham Hotspur",
	}

	pl1.stateTeam()
	pl1.stateAge()
}
