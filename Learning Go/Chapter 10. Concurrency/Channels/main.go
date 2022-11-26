package main

import "fmt"

func integers(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i

	}
	close(ch)
}

func main() {
	channel := make(chan int)
	go integers(channel)
	for v := range channel {
		fmt.Println(v)
	}

	buffered_channel := make(chan int, 1)
	buffered_channel <- 1
	value := <-buffered_channel
	fmt.Println(value)

}
