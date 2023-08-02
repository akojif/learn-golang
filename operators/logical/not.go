package main

import "fmt"

func main() {
	var x, y int = 10, 20

	fmt.Println(!(x < y))
	fmt.Println(!(x > y))
	fmt.Println(!true)
	fmt.Println(!false)
}
