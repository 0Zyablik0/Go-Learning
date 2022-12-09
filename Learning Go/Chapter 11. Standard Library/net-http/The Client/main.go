package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	http
	fmt.Println(client)
}
