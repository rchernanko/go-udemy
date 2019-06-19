package main

import "fmt"

func main() {

	/*

		An expression is when we have something that comes down to a value.
		And so the value that we're going to have here is a function and we're going to sign it to a variable
		and that's what's known as a 'func expression'

	*/

	f := func() {
		fmt.Println("my first func expression")
	}

	f()
}

/*

EXPRESSION VS STATEMENT

From here - https://www.quora.com/Whats-the-difference-between-a-statement-and-an-expression-in-Python-Why-is-print-%E2%80%98hi%E2%80%99-a-statement-while-other-functions-are-expressions

Here’s a general rule of thumb: If you can print it, or assign it to a variable, it’s an expression. If you can’t, it’s a statement.

Here are some examples of expressions:

- 2 + 2
- 3 * 7
- 1 + 2 + 3 * (8 ** 9) - sqrt(4.0)
- min(2, 22)
- max(3, 94)
- round(81.5)
- "foo"
- "bar"
- "foo" + "bar"
- None
- True
- False
- 2
- 3
- 4.0

All of the above can be printed or assigned to a variable.

Here are some examples of statements:

- if CONDITION:
- else:
- for:
- while CONDITION:
- try:
- return SOMETHING

None of the above constructs can be assigned to a variable. They are syntactic elements that serve a purpose, but do
not themselves have any intrinsic “value”. In other words, these constructs don’t “evaluate” to anything.

*/
