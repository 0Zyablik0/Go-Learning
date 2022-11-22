package main

import "fmt"

func main() {
	intSet := map[int]bool{}
	values := []int{1, 2, 3, 4, 5, 2, 3, 5, 9, 10}
	for _, v := range values {
		intSet[v] = true
	}
	fmt.Println(len(intSet), len(values))

	// Alternative:
	testSet := map[int]struct{}{}
	for _, v := range values {
		testSet[v] = struct{}{}
	}
	fmt.Println(len(testSet), len(values))

}
