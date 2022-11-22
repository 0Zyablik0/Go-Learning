package main

import "fmt"

func main() {
	testSlice := []int{1, 2, 3, 4, 5}
	testCopySlice := make([]int, 5)
	n := copy(testCopySlice, testSlice)
	fmt.Println(testCopySlice, n)
}
