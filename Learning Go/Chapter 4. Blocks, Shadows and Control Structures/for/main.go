package main

import "fmt"

func main() {
	// Complete for statement:

	for i := 0; i < 5; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	//The condition only

	i := 0
	for i < 100 {
		i += 2
	}
	fmt.Println(i)

	// for-range statement

	evenValues := []int{2, 4, 6, 8, 10}
	for i, v := range evenValues {
		evenValues[i] -= 1
		fmt.Printf("%d \t", v)
	}
	fmt.Println("\n", evenValues)

	for _, v := range evenValues {
		fmt.Printf("%d \t", v)
	}
	fmt.Println()

}
