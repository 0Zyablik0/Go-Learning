package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		channel <- 1
	}()

	time.Sleep(time.Millisecond * 100)

	select {
	case n := <-channel:
		fmt.Println(n)
	default:
		fmt.Println("nothing in the channel")
	}
}
