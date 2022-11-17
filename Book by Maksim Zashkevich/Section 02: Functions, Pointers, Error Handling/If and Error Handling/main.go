package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	printCircleArea(-1)
	printCircleArea(0)
	printCircleArea(1)
	printCircleArea(2)
	printCircleArea(3)

}

func circleArea(r float64) (float64, error) {
	if r < 0 {
		return 0, errors.New("Circle Radius must be positive")
	}
	return pi * r * r, nil
}

func printCircleArea(r float64) {

	area, err := circleArea(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Area of the circle of radius %.2f is %.2f\n", r, area)

}
