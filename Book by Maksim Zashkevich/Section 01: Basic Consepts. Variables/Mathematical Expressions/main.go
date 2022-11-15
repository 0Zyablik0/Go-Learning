package main

import (
	"fmt"
)

const pi = 3.1415

func main() {
	circleRadius := 2
	circleArea := float32(circleRadius) * float32(circleRadius) * pi
	fmt.Printf("radius: %d m\n", circleRadius)
	fmt.Printf("area: %.4f m^2\n", circleArea)
}
