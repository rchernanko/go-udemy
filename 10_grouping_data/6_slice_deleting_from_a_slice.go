package main

import "fmt"

func main() {

	x := []int{234, 5, 1, 6, 2}
	fmt.Println(x) //prints [234 5 1 6 2]

	//I want to delete '1' from the the slice 'x'
	x = append(x[:2], x[3:]...) //we have to 'unfurl'
	fmt.Println(x)              //prints [234 5 6 2]
}
