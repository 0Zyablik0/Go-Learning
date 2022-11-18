package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	Human  /// anonymous field is used for embedding
	school string
}

func main() {
	student := Student{}
	fmt.Printf("Empty Student: %+v\n", student)

	student = Student{
		Human: Human{
			name: "John",
			age:  13,
		},
		school: "High School",
	}
	fmt.Printf("Initialized Student: %+v\n", student)
	student.sayHi()
	student.Human.sayHi()

}

func (h *Human) sayHi() {
	fmt.Printf("Hi, I'm %s\n", h.name)
}
