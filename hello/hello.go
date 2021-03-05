package main

import (
	"fmt"

	"streetleague.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Utku")
	fmt.Println(message)
}
