package main

import "fmt"

func main() {
	var fruit string = "random"
	if fruit == "apple" {
		fmt.Println("I love", fruit+"s")
	} else if fruit == "oranges" {
		fmt.Println("Oranges are not apples")
	} else {
		fmt.Println("No appetite")
	}
}
