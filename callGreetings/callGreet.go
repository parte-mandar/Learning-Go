package main

import (
    "fmt"
	"log"
    "example.com/greetings"
	"example.com/greetings/package2"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Get a greeting message and print it.
    message, err := greetings.Hello("M") // function from greetings module
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err) // This stops the execution completely
		// fmt.Println(err)
    } else {
		fmt.Println(message)
	}

	// If no error was returned, print the returned message
    // to the console.
	fmt.Println(message) // Here nothing prints if message is empty

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