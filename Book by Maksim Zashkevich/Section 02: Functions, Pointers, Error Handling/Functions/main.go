package main

import (
	"fmt"
)

const pi = 3.1415

func main() {
	fmt.Println(CircleArea(1))
	fmt.Println(CircleArea(2))
	fmt.Println(CircleArea(3))

}

func CircleArea(r float64) float64 {
	return pi * r * r
}
