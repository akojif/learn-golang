package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	for index, element := range slice { // can be either value or element
		fmt.Println(index, "=>", element)
	}
}
