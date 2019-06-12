package main

import "fmt"

func main() {

	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb) // [james bond chocolate martini]
	mp := []string{"miss", "moneypenny", "strawberry", "hazelnut"}
	fmt.Println(mp) // [miss moneypenny strawberry hazelnut]

	//Let's create a slice of a slice of strings
	xp := [][]string{jb, mp}
	fmt.Println(xp)
	//prints [[james bond chocolate martini] [miss moneypenny strawberry hazelnut]]

}
