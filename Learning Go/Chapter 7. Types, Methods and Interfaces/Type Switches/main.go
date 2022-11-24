package main

import "fmt"

func doThings(i interface{}) {
	switch t := i.(type) {
	case nil:
		fmt.Println("This is empty interface")

	case int:
		fmt.Println("This is int: ", t)
	case float64, float32:
		fmt.Println("This is float64: ", t)
	case string:
		fmt.Println("this is string: ", t)
	case bool:
		fmt.Println("This is bool: ", t)
	default:
		fmt.Println("This is unknown type")
	}

}

func main() {
	doThings(nil)
	doThings(10)
	doThings("hello")
	doThings(3.2)
	doThings(true)
	doThings(struct{}{})
}
