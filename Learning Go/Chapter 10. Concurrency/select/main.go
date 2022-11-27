package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v1 := 1
		ch1 <- v1
		v2 := <-ch2
		fmt.Println(v1, v2)
	}()

	var v1 int = 2
	var v2 int

	select {
	case ch2 <- v1:
	case v2 = <-ch1:
	}
	fmt.Println(v1, v2)

	done := make(chan bool)

	select {
	case <-done:
	default:
		fmt.Println("no value read from channel")
	}

}
