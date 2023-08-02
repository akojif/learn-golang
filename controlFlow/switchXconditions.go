package main

import "fmt"

func main() {
	var a, b int = 10, 2
	switch {
	case a+b == 12:
		fmt.Println("equals to 30")
		fallthrough
	case a+b <= 12:
		fmt.Println("less than or equals to 30")
	default:
		fmt.Println("greater than 30")

	}
}
