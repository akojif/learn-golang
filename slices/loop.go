package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	for index, element := range slice { // for use cases where you don't want to return the index, you can just represent the index as _
		fmt.Println(index, "=>", element) // can be either value or element
	}
}
