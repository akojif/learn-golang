package main

import (
	"fmt"

	"francisakoji.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Francis")
	fmt.Println(message)
}
