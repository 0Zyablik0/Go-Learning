package main

import (
	"fmt"
)

var globalVar string = "This is a global variable"

func main() {
	localVar := "This is a local variable" // short syntax for local variables
	fmt.Println(globalVar)
	fmt.Println(localVar)
}
