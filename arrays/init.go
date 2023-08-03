package main

import "fmt"

func main() {
	var grade [3]int = [3]int{97, 85, 93}
	fmt.Println(grade)

	grades := [3]int{97, 85, 93}
	fmt.Println(grades)

	gradess := [...]int{97, 85, 93}
	fmt.Println(gradess)

	names := [...]string{"Francis", "Akoji", "AKF", "iDan"}
	fmt.Println(names)
}
