package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("The type of 1000 is %v \n", reflect.TypeOf(1000))
	fmt.Printf("The data type of Francis is %v \n", reflect.TypeOf("Francis"))
	fmt.Printf("The data type of 43.67 is %v \n", reflect.TypeOf(43.67))

}
