package main

import "fmt"

func main() {

	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice := array[0:9]

	fmt.Println(slice)

	slice_2 := append(slice, 10, 11, 12)

	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))

}
