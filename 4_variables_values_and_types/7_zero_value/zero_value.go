package main

import "fmt"

func main() {

	/*

	DECLARE a variable of a certain TYPE and then ASSIGN a VALUE of that type to the variable

	E.g.

	 */

	var1 := 24
	fmt.Println("printing the value of var1:", var1)
	fmt.Printf("%T\n", var1)

	var number2 string = "richard"
	fmt.Println("printing the value of number2:", number2)
	fmt.Printf("%T\n", number2)

	var number3 = "sebastian"
	fmt.Println("printing the value of number3:", number3)
	fmt.Printf("%T\n", number3)

	/*

	ZERO VALUES - https://golang.org/ref/spec#The_zero_value:

	Some examples:

	 */

	//String
	var zeroValueString string
	fmt.Println("printing the zero value of zeroValueString", zeroValueString, "ending")

	//Int
	var zeroValueInt int
	fmt.Println("printing the zero value of zeroValueInt", zeroValueInt)

	//Bool
	var zeroValueBool bool
	fmt.Println("printing the zero value of zeroValueBool", zeroValueBool)

	//Struct
	var zeroStruct Person
	fmt.Println("printing the zero value of zeroStruct", zeroStruct) //printing the zero value of zeroStruct { 0 false}
	fmt.Printf("%T\n", zeroStruct) //main.Person
}


type Person struct {
	Name string
	Age int
	IsMale bool
}