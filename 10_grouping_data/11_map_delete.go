package main

import "fmt"

func main() {

	m := map[string]int{
		"Richard": 32,
		"John":    56,
		"Amy":     14,
	}

	fmt.Println("BEFORE DELETE", m) //BEFORE DELETE map[Richard:32 John:56 Amy:14]

	//Let's delete Amy
	delete(m, "Amy")

	fmt.Println("AFTER DELETE", m) //AFTER DELETE map[Richard:32 John:56]

	if _, ok := m["Amy"]; !ok {
		fmt.Println("Amy has been deleted from the map")
	}
}
