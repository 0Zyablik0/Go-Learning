package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	result1 := make(chan int)
	result2 := make(chan int)
	fmt.Printf("\nTest starts at %s\n", t.Format(time.RFC3339))

	go process(500, result1)
	go process(1000, result2)

	fmt.Printf("result1 is %d", <-result1)
	fmt.Printf("result2 is %d", <-result2)

	fmt.Printf("Test  spends: %s\n", time.Since(t))

}

func process(n int, out chan int) {
	t := time.Now()
	result := 0
	for i := 0; i < n; i++ {
		result += i * i
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Result: %d, time spent: %s\n", result, time.Since(t))
	out <- result

}
