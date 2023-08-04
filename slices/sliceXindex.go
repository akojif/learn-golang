package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	slice := array[:5]

	fmt.Println(array)
	fmt.Println(slice)

	slice[1] = 15

	fmt.Println("after modification of slices")
	fmt.Println(array)
	fmt.Println(slice)
}
