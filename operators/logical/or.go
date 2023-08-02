package main

import "fmt"

func main() {
	x := 10

	fmt.Println((x < 0) || (x < 20))
	fmt.Println((x > 100) || (x < -2))

}
