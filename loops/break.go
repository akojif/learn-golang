package main

import "fmt"

func main() {
	for a := 1; a <= 5; a++ {
		if a == 4 {
			break
		}
		fmt.Println(a)
	}
}
