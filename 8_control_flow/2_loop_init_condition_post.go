package main

import "fmt"

/*

- First thing he recommends is the 'go by example' website - https://gobyexample.com/
- There is NO WHILE LOOP in Go

- Basic loop is structured as below (as per other languages I've used so far):

	for init; condition; post {

	}

*/

func main() {

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
