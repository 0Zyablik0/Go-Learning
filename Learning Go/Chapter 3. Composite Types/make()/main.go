package main

import "fmt"

func main() {
	testSlice := make([]int, 5)
	fmt.Println(testSlice, len(testSlice), cap(testSlice))

	testSlice = append(testSlice, 10)
	fmt.Println(testSlice, len(testSlice), cap(testSlice))

	testSlice = make([]int, 5, 10)
	fmt.Println(testSlice, len(testSlice), cap(testSlice))

	testSlice = make([]int, 0, 10)
	testSlice = append(testSlice, 5, 6, 7, 8)
	fmt.Println(testSlice, len(testSlice), cap(testSlice))

}
