package main

import (
	"fmt"
)

func main() {
	var name string
	var score int

	fmt.Print("What's your name and score in the test? ")
	fmt.Scanf("%s %d", &name, &score)
	fmt.Println("The score of", name, "in the test is", score)
}
