package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}

	slice := array[1:5]

	sub_slice := slice[0:3]

	fmt.Println(sub_slice)
}
