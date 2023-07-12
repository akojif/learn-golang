package main

import (
	"fmt"
)
const PI float64 = 3.142 //global variable

func main() {
	
	var radius float64 = 5.0
	var area float64 
	area = PI * radius * radius

	fmt.Println("The area of the circle is", area)

}
