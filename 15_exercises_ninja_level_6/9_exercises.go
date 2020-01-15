package main

import "fmt"

func main() {
	something(calculateNumbers, "sdf")
}

func something(f func(num1, num2 int) int, word string) {
	fmt.Println(word)

	fmt.Println(f(1, 2))
}

func calculateNumbers(num1, num2 int) int {
	return num1 + num2
}
