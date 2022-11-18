package main

import (
	"fmt"
)

type age int

func (a age) isAdult() bool {
	return a >= 18
}

func main() {
	testAge := age(23)
	fmt.Println("Is adult: ", testAge.isAdult())

}
