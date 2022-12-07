package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("local time is: ", now)
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2022-11-02 07:50:40 +0300")
	if err != nil {
		fmt.Println("Parse error: ", err)
	} else {
		fmt.Println(t.Format(time.RFC850))
	}
}
