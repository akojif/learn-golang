package main

import "fmt"

func main() {
	src_slice := []int{10, 20, 30, 40, 50}
	dest_slice := make([]int, 3)

	num := copy(dest_slice, src_slice)

	fmt.Println(num)

	fmt.Println("Number of elements copied:", dest_slice)
}
