package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred person
	fmt.Println("Fred: ", fred)

	bob := person{}
	fmt.Println("Bob: ", bob)

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println("Julia: ", julia)

	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println("Beth: ", beth)

	bob.name = "Bob"
	fmt.Println("Bob: ", bob)

}
