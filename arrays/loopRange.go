package main

import "fmt"

func main() {
	grades := [...]int{34, 23, 24, 45, 483, 392}

	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}
}
