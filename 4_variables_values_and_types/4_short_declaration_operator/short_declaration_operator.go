package main

import "fmt"

/*

- Some info on 'identifiers' from the GoLang language spec - https://golang.org/ref/spec#Identifiers

	"

	Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters
	and digits. The first character in an identifier must be a letter.

	identifier = letter { letter | unicode_digit } .

	a
	_x9
	ThisVariableIsExported
	αβ

	"

- And the same for keywords - https://golang.org/ref/spec#Keywords

"
The following keywords are reserved and may not be used as identifiers.

break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
"

- And the same for operators and punctuation - https://golang.org/ref/spec#Operators_and_punctuation

"
The following character sequences represent operators (including assignment operators) and punctuation:

+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
"

 */

func main() {

	//Some examples of the short declaration operator:

	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 24 //semantically, the 100 and the 24 are the 'operands' and the + is the 'operator'
	fmt.Println(y)
}