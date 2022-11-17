package main

import (
	"fmt"
)

func main() {
	x := 10
	p := &x
	fmt.Println("Original x: ", x)
	fmt.Println("Address of x: ", &x)
	fmt.Println("Value of p: ", p)
	fmt.Println("Value of *p: ", *p)

	*p = 15
	fmt.Println("New value of x: ", x)

}
