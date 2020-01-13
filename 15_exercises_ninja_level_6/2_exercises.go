package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := foo1(nums...)
	fmt.Println(sum)
}

func foo1(numbers ...int) int {
	var sum int
	for _, v := range numbers {
		sum = sum + v
		//could also do `sum += v`
	}

	return sum
}
