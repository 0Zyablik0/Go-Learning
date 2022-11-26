package main

import "fmt"

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)

}

func main() {
	for _, v := range []int{1, 2, 0, 5, 0, 3, 4, 5, 0, 6, 10, 12, 15, 0, 20, 0, 30} {
		div60(v)
	}

}
