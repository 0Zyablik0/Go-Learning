package main

import (
	"fmt"
	"time"
)

func main() {
	sequenceTest()
	concurrentTest()

}

func process(n int) {
	t := time.Now()
	result := 0
	for i := 0; i < n; i++ {
		result += i * i
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Result: %d, time spent: %s\n", result, time.Since(t))

}

func concurrentTest() {
	t := time.Now()
	fmt.Printf("\nConcurrent test starts at %s\n", t.Format(time.RFC3339))

	go process(500)
	process(1000)

	fmt.Printf("Concurrent test  spends: %s\n", time.Since(t))
}

func sequenceTest() {
	t := time.Now()
	fmt.Printf("\nSequence test starts at %s\n", t.Format(time.RFC3339))

	process(500)
	process(1000)

	fmt.Printf("Sequence test  spends: %s\n", time.Since(t))
}
