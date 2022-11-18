package main

import (
	"fmt"
)

func main() {

	printType(1)
	printType("this is a string")
	printType(1.2)
	printType(map[string]int{"1": 1})

}

func printType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("value is int")
	case string:
		fmt.Println("value is string")
	default:
		fmt.Println("value is neither int nor string")
	}
}
