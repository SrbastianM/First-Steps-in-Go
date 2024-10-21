package main

import (
	"fmt"
	pkIsPair "learning_go/src/isPair"
	pkSquare "learning_go/src/square"
	pkValidation "learning_go/src/userValidation"
)

func main() {
	// password and username validation and is pair or not
	const pass string = "1234"
	const user string = "test"
	// is Pair
	fmt.Println(pkIsPair.IsPair(2))

	// Print Areas
	pkSquare.Square()

	// User Validation
	fmt.Println(pkValidation.ValidateUser(user, pass))

}
