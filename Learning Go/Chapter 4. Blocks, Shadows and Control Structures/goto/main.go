package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10)
	for n < 100 {
		if n%5 == 0 {
			goto done
		}
		n = n*2 + 1
	}
	fmt.Println("do something if loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(n)
}
