package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return numerator / denominator, numerator % denominator, nil

}

func main() {
	result, remainder, err := divAndRemainder(5, 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

}