package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile",
		"gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); {
		case size < 5:
			fmt.Println(word, " is a short word!")
		case size > 10:
			fmt.Println(word, " is a long word!")
		default:
			fmt.Println(word, " is exactly the right length:", size)

		}
	}
}
