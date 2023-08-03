package main

import "fmt"

func main() {
	fruits := [...]string{"Orange", "Apple", "Grape", "Pear", "Mango"}

	fmt.Println(fruits[2])

	grades := [...]int{97, 85, 93}

	fmt.Println(grades)

	grades[2] = 100

	fmt.Println(grades)

}
