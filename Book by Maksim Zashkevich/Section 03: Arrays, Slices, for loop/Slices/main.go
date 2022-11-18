package main

import (
	"fmt"
)

func main() {
	var strSlice = []string{
		"buy milk",
		"buy eggs",
		"buy coffee",
	}

	fmt.Printf("Slice length : %d, slice capacity : %d\n", len(strSlice), cap(strSlice))

	newStrSlice := append(strSlice, "buy tea")

	fmt.Printf("Slice length : %d, slice capacity : %d\n", len(strSlice), cap(strSlice))
	fmt.Printf("New slice length : %d, New slice capacity : %d\n", len(newStrSlice), cap(newStrSlice))

	var tasks []string
	fmt.Println(tasks == nil)

	tasks = newStrSlice[1:3] // if [:] then we choose all elements
	fmt.Println(tasks == nil)

	for i := range tasks {
		fmt.Println(tasks[i])
	}
	newStrSlice[1] = "buy water"
	fmt.Println("---------Original array changed------------")
	for i := range tasks {
		fmt.Println(tasks[i])
	}

}
