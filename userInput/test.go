package main

import (
	"fmt"
)

func main() {
	var a string
	var c int

	fmt.Print("Enter the value of a and b")
	count, err := fmt.Scanf("%s %d", &a, &c)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(count)
	fmt.Println(err)

}
