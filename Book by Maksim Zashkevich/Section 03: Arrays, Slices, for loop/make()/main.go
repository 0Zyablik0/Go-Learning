package main

import (
	"fmt"
)

func main() {

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 1, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(name string, slice []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n", name, len(slice), cap(slice), slice)

}
