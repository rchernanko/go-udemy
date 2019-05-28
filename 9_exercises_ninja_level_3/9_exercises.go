package main

import "fmt"

func main() {

	favSport := "football"

	switch favSport {
	case "basketball", "tennis":
		fmt.Println("not my favourite sport")
	case "cricket":
		fmt.Println("not my favourite sport either")
	case "football":
		fmt.Println("love it")
	}
}
