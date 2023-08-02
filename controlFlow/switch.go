package main

import "fmt"

func main() {
	var i int = 20
	switch i {
	case 10:
		fmt.Println("i is equal to 10")
	case 100, 200:
		fmt.Println("i is either 100 or 200")
	default:
		fmt.Println("All I know is i is an integer")
	}
}
