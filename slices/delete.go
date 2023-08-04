package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 4

	slice1 := array[0:i]
	slice2 := array[i+1:]
	new_slice := append(slice1, slice2...)

	fmt.Println(new_slice)
}
