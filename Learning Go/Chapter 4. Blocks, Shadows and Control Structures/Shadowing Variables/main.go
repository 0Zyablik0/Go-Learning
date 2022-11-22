package main

import "fmt"

func main() {

	// Example 1
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// Example 2
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

	//Example 3
	fmt.Println(true)
	true := 10
	fmt.Println(true)

}
