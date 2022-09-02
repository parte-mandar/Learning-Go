package main

import "fmt"

var web = "google.com" // Public

func main() {
	fmt.Println("Vars")
	var username string = "mandar"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T (%s)\n", username, username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T (%t)\n", isLoggedIn, isLoggedIn)

	var smallVal uint8 = 255 // More than 255 not allowed in uint8, see docs https://go.dev/ref/spec#Variables
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T (%d)\n", smallVal, smallVal)

	var smallFloat float64 = 255.78462387422423423
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T (%f)\n", smallFloat, smallFloat)

	// Default values and aliases
	var someValue int // Default value will be 0
	fmt.Println(someValue)
	fmt.Printf("Variable is of type: %T (%d)\n", someValue, someValue)

	// Implicit type of variable declaration
	var web = "google.com"
	fmt.Println(web)

	// no var style
	num := 78632
	fmt.Println(num)
	num = 89432
	fmt.Println(num)
}
