package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum2(2, 3))
}

func mySum2(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
