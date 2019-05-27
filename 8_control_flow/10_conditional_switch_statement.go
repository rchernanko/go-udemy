package main

import "fmt"

func main() {

	//Switch statement with no fallthrough keyword

	fmt.Println("FIRST SWITCH WITHOUT FALLTHROUGH")

	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print2")
	case 3 == 3:
		fmt.Println("prints")
	case 4 == 4:
		fmt.Println("also true, does it print?")
	}

	/*

		Interestingly enough, only the 3rd case prints. By default, the switch statement doesn't fall through (i.e. once
		the 3rd case has executed, the 4th case is not executed). You can however change this by using the 'fallthrough' keyword
		Let's specify a fallthrough in the switch statement below (so that the 3rd AND 4th cases both print)

	*/

	fmt.Println("SECOND SWITCH WITH FALLTHROUGH - 3rd and 4th statements should print")

	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print2")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 == 4:
		fmt.Println("also true, does it print?")
	}

	/*

		One thing to be very aware of when using the fallthrough keyword. Even if the switch case equates to false, its
		block will still be executed. So in the example below, even though 7 is not equal to 9, because of the fallthrough
		keyword in the previous case block, 'not true1' will be executed: BE CAREFUL. In fact, Todd said at the end...

		"don't use fallthrough"

	*/

	fmt.Println("THIRD SWITCH WITH FALLTHROUGH - BE CAREFUL WHEN USING FALLTHROUGH KEYWORD!!!")

	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print2")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 == 4:
		fmt.Println("also true, does it print?")
		fallthrough
	case 7 == 9:
		fmt.Println("not true1")
		fallthrough
	case 11 == 14:
		fmt.Println("not true2")
		fallthrough
	case 15 == 15:
		fmt.Println("true 15")
	}

	fmt.Println("FOURTH SWITCH - WITH DEFAULT")

	switch {
	case false:
		fmt.Println("this won't print")
	case 4 == 2:
		fmt.Println("either will this")
	default:
		fmt.Println("this will print because it's default and no other cases are true")
	}

	fmt.Println("FIFTH SWITCH - EVALUATES A VALUE")

	switch "spurs" {
	case "spurs":
		fmt.Println("best team in the land")
	case "goons":
		fmt.Println("absolute rubbish")
	default:
		fmt.Println("print default")
	}

	fmt.Println("SIXTH SWITCH - EVALUATES A VALUE AND HAS MULTIPLE VALUES IN CASE")

	switch "bond" {
	case "bond", "moneypenny", "dr no":
		fmt.Println("bond or moneypenny or dr no")
	case "M":
		fmt.Println("I am M")
	default:
		fmt.Println("this is default")
	}
}
