package main

import "fmt"

func main() {
	defer finally()
	initially()
}

func initially() {
	fmt.Println("initially running")
}

func finally() {
	fmt.Println("running via the defer")
}
