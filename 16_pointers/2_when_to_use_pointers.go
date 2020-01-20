package main

import "fmt"

/*

- When should you use pointers?
1)
- Pointers are good if you have a large chunk of data and you don't want to pass that large chunk of data around your application
- You can instead just pass that address where that data is stored
- So that can help in terms of performance (because everything in go is 'pass by value' - and so if you don't use
pointers to pass around the large chunk of data, when its passed into a function, it will be copied and then added to
a new memory address. So you'd be continually creating the same large chunk of data in different memory addresses).
2)
- You can also use pointers when you want to change a value that's at a certain location.

KEY POINT OF THIS VIDEO - EVERYTHING IN GO IS PASS BY VALUE

 */

func main() {
	x := 0
	withoutPointer(x)
	fmt.Println(x) //x will still be 0 when we print it here

	z := 1
	fmt.Println(&z) //memory address
	withPointer(&z)
	fmt.Println(z) //90

	//TODO up 5 mins 13 in video
}

func withoutPointer(y int) {
	fmt.Println(y) //0
	y = 43
	fmt.Println(y) //43
}

func withPointer(y *int) {
	fmt.Println(*y) //1
	fmt.Println(y) //memory address
	*y = 90
	fmt.Println(*y) //90
}