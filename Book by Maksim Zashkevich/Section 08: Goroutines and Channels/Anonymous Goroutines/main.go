package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	time.Sleep(time.Second * 3)
}
