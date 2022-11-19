package main

import (
	"fmt"
)

func main() {
	buffer := make(chan int)
	go generateSquares(10, buffer)

	for square := range buffer {
		fmt.Println(square)
	}

	fmt.Println("----------------------------------------------------------------")

	buffer = make(chan int)
	go generateSquares(10, buffer)
	for {
		square, ok := <-buffer
		if !ok {
			break
		}
		fmt.Println(square)

	}

}

func generateSquares(n int, out chan int) {
	for i := 0; i <= n; i++ {
		out <- i * i
	}

	close(out) // wee need to close the channel to prevent deadlock

}
