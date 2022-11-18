package main

import (
	"fmt"
)

func main() {
	basicArray()
	partialInit()
	guessedNumberInit()

}

func basicArray() {
	var strArray = [3]string{
		"buy milk",
		"buy eggs",
		"buy coffee",
	}

	fmt.Printf("Number of items in array: %d\n", len(strArray))

	for index, item := range strArray {
		fmt.Printf("%d: %s\n", index, item)
	}

}

func partialInit() {
	var strArray = [3]string{
		"buy milk",
	}

	strArray[1] = "buy eggs"

	fmt.Printf("Number of items in array: %d\n", len(strArray))

	for index, item := range strArray {
		fmt.Printf("%d: %s\n", index, item)
	}

}

func guessedNumberInit() {
	var strArray = [...]string{
		"buy milk",
		"buy eggs",
		"buy coffee",
	}

	fmt.Printf("Number of items in array: %d\n", len(strArray))

	for index, item := range strArray {
		fmt.Printf("%d: %s\n", index, item)
	}

}
