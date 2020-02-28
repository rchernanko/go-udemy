package main

import "fmt"

/*

- Let's fix the blocking channel issue from the previous file
- There are actually 2 ways to do this. Here is the second
- This way uses the concept of buffered channels

*/

func main() {

	/*

		By adding a second argument here, we are creating a buffered channel
		A buffered channel allows certain number of values to sit in that channel regardless of whether something is ready to pull it off.
		In our case, the buffered channel can have 1 item put on it

	*/

	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
