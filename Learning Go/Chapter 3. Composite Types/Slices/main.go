package main

import "fmt"

func main() {
	var testSlice = []int{1, 2, 4}
	fmt.Println(testSlice)

	var partialInitTestSlice = []int{1, 5: 2, 3: 4, 7: 100, 15}
	fmt.Println(partialInitTestSlice)

	var multiDimTestSlice = [][][]int{}
	fmt.Println(multiDimTestSlice == nil)

	var testNilSlice []int
	fmt.Println(testNilSlice == nil)

}
