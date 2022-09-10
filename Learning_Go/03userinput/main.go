package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Namaskaram, please give input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza ")

	// 'Comma ok' or 'error ok' syntax
	inp, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: " + inp)
	fmt.Printf("Variable is of type: %T - %s \n", inp, inp)
}
