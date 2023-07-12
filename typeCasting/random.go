package main

import (
	"fmt"
	"strconv" // Using the string conversion packaage
)

func main() {
	var score int = 390
	var num string = strconv.Itoa(score) // Convert integer to string using the imported package (Strconv)

	fmt.Printf("%q \n", num)

}
