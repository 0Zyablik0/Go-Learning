package main

import (
	"fmt"
)

func main() {
	var strArray = [3]string{
		"buy milk",
		"buy eggs",
		"buy coffee",
	}

	fmt.Printf("Number of items in array: %d\n", len(strArray))

	for index, item := range strArray {
		fmt.Printf("%d: %s\n", index, item)
	}

	for index := range strArray {
		fmt.Printf("%d: %s\n", index, strArray[index])
	}

	for _, item := range strArray {
		fmt.Printf("%s\n", item)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)

		if i == 5 {
			break
		}
	}

}
