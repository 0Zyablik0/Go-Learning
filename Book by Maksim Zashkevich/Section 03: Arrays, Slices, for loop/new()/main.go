package main

import (
	"fmt"
)

func main() {
	var p = new(int)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 100
	fmt.Println(p)
	fmt.Println(*p)

}
