package main

import "fmt"

func addTo(base int, values ...int) []int {
	out := make([]int, 0, len(values))
	for _, v := range values {
		out = append(out, v+base)
	}
	return out
}

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

}
