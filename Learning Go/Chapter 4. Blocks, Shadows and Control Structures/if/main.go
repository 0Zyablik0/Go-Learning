package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if n := rand.Intn(10); n == 0 {
		fmt.Println("that's too low")
	} else if n > 5 {
		fmt.Println("that's too big", n)
	} else {
		fmt.Println("that's a good number", n)
	}

}
