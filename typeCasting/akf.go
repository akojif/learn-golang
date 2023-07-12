package main

import (
	"fmt"
	"strconv" // Use string conversion package
)

func main() {
	var score string = "500"      // If this value of the string was different, let's say "200abc" 0- you'd see the error that will be thrown when the code is ran
	i, err := strconv.Atoi(score) // Convert string to integer - it returns two values (one actual value and an error)

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", err, err)

}
