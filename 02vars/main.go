package main

import "fmt"

// walrus operator not valid outside function scope

// Big letter in start of variable name means its public
const Token string = "token"

func main() {
	var username string = "name"
	fmt.Printf("variable is of type %T \n", username)

	var loggedIn bool = false
	fmt.Printf("variable is of type %T \n", loggedIn)

	var smallVal uint8 = 255
	fmt.Printf("variable is of type %T \n", smallVal)

	string := 500.0
	fmt.Printf("variable %.1f is of type %T \n", string, string)
}
