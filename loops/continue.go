package main

import "fmt"

func main() {
	for i := 2; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
