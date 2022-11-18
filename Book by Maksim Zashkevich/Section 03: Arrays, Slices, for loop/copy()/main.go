package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)

	n_copied := copy(dst, src)

	fmt.Println("Number of elements copied: ", n_copied)
	fmt.Println("source: ", src)
	fmt.Println("Destination: ", dst)

	dst[2] = -1
	fmt.Println("\nChanging dst slice")
	fmt.Println("source: ", src)
	fmt.Println("Destination: ", dst)

	dst = make([]int, 4)
	n_copied = copy(dst, src)

	fmt.Println("\nThe Length of destination is less than the length of source")
	fmt.Println("Number of elements copied: ", n_copied)
	fmt.Println("source: ", src)
	fmt.Println("Destination: ", dst)

	dst = make([]int, 6)
	n_copied = copy(dst, src)

	fmt.Println("\nThe Length of destination is more than the length of source")
	fmt.Println("Number of elements copied: ", n_copied)
	fmt.Println("source: ", src)
	fmt.Println("Destination: ", dst)

}
