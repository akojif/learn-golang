package main

import "fmt"

func main() {
	array1 := [...]int{10, 20, 30, 40, 50}

	slice1 := array1[:5]

	array2 := [...]int{60, 70, 80, 90, 100}

	slice2 := array2[:5]

	new_slice := append(slice1, slice2...)

	fmt.Println(new_slice)
}
