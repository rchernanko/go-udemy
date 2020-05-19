package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

/*

What is bcrypt?

- If you're storing passwords anywhere, bcrypt is an excellent way of storing password information
- As soon as you receive a password from e.g. a user registration, you can use bcrypt to encrypt it, store the
encrypted value + then you never even know the value of that password
- When you then receive a password from e.g. a user login, you run it through bcrypt and then you can compare the value
with the value you have stored
- The basics of encryption are taking a value, and jumbling it all up. You only ever want to store an encrypted version
of user's passwords

Go implementation of bcrypt:
- https://godoc.org/golang.org/x/crypto/bcrypt
- Interestingly, the "x" package here is basically some go code that is more experimental and has not yet made it into
the standard library

*/

func main() {
	s := `password123` //raw literal
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginPassword := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("passwords match!")
	}

	//TODO do we use bcrypt in the management application
}
