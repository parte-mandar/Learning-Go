package main

import (
    "fmt"
    "example.com/greetings"
	"example.com/greetings/package2"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys") // function from greetings module
    fmt.Println(message)

	// fmt.Println(greetings.namaskaram("Pure Soul")) 
	// Above line gives an error as namaskaram() starts with small 'n' which is not accessible outside greetings module
	// Instead, following is the same function with Namaskaram ('N' capital)
	fmt.Println(greetings.Namaskaram("Pure Soul")) 
	// The function Namaskaram() is in different file, but in same module.

	fmt.Println(greetings.AddingNamaskaram("Pure", "Soul"))
	// This returns content of namaskaram() in greetings module through different functn

	// Same module but different package
	fmt.Println(pkg2.Namaskaram("Pure Soul") + " (From package2)")
	// Different package can have function with same name
	// Package needs to be imported seperately
}