package greetings

import "fmt"

// Hello returns a greeting for the named person.
// Hello has to start with capital 'H' to be accessed outside
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name) 
    return message
}


func namaskaram(name string) string {
    return "Namaskaram " + name
}