package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

Let's explore the fmt package a little more - https://godoc.org/fmt

There are basically loads of different options to print / scan

 */

var y = 42 //Here, I have declared a value of a certain type and assigned a value of that type to it

func main() {
	fmt.Println(y)

	fmt.Printf("%T\n", y) //int
	fmt.Printf("%b\n", y) //101010 (binary)
	fmt.Printf("%x\n", y) //2a (hexadecimal)
	fmt.Printf("%#x\n", y) //0x2a (hexadecimal with 0x)

	boolLove := true
	fmt.Printf("%t\n", boolLove) //true

	//Can pass in multiple values using printf too:
	fmt.Printf("%T\n%b\n%x\n", y, y, y)

	//Sprint, Sprintf, Sprintln are also useful to know - "Sprintf formats according to a format specifier and returns the resulting string."
	var name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	fmt.Println(s) //Kim is 22 years old.



	//Fprint, Fprintf, Fprintln are also useful to know. Prints to a file
	f, _ := os.Create("richard.txt")
	w := bufio.NewWriter(f)

	// Use Fprint to write things to the file.
	// ... No trailing newline is inserted.
	fmt.Fprint(w, "Hello")
	fmt.Fprint(w, 123)
	fmt.Fprint(w, "...")

	// Use Fprintf to write formatted data to the file.
	value1 := "cat"
	value2 := 900
	fmt.Fprintf(w, "%v %d...", value1, value2)

	fmt.Fprintln(w, "DONE...")

	// Done.
	w.Flush()

}