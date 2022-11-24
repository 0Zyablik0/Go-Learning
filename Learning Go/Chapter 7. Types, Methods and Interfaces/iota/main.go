package main

import "fmt"

type MailCategory int
type BitField int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisement
)

const (
	Field1 BitField = 1 << iota
	Field2
	Field3
	Field4
)

func main() {

	fmt.Println(Uncategorized, Personal, Spam, Social, Advertisement)
	fmt.Println(Field1, Field2, Field3, Field4)

}
