package main

import (
	"errors"
	"fmt"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleOdd(i int) (int, error) {
	if i%2 == 0 {
		return 0, fmt.Errorf("%d is not an odd number", i)
	}
	return i * 2, nil
}
