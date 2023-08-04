package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 6, 7, 8, 9, 10}

	slice := array[0:7]

	fmt.Println(cap(array))
	fmt.Println(cap(slice))
}
