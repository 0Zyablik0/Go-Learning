package main

import "fmt"

func main() {
	var nilTestArray []int
	nilTestArray = append(nilTestArray, 10)
	fmt.Println(nilTestArray)

	testArray := []int{1, 2, 3}
	testArray = append(testArray, 4)
	fmt.Println(testArray)

	testArray = append(testArray, 5, 6, 7)
	fmt.Println(testArray)

	var addTestSlice = []int{8, 9, 10}
	testArray = append(testArray, addTestSlice...)
	fmt.Println(testArray)

}
