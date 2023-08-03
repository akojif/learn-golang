package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93, 56, 33}

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
}
