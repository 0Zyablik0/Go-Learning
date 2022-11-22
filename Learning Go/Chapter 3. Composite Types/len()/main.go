package main

import (
	"fmt"
)

func main() {
	var partialInitTestSlice = []int{1, 5: 2, 3: 4, 7: 100, 15}
	fmt.Println(len(partialInitTestSlice))

	var testNilSlice []int
	fmt.Println(len(testNilSlice))
}
