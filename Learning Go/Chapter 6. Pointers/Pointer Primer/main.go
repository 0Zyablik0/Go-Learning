package main

import (
	"fmt"
)

type Foo struct{}

type testStruct struct {
	name    string
	surname *string
}

func main() {
	x := 10
	var pointerToX *int = &x
	fmt.Println(pointerToX, &x)
	fmt.Println(x)

	var y = new(int)
	fmt.Println(y)
	fmt.Println(*y)

	structTest := &Foo{}
	fmt.Println(structTest, *structTest)

	z := testStruct{
		"alex",
		stringPointer("marty"),
	}

	fmt.Println(z, *z.surname)
}

func stringPointer(s string) *string {
	return &s
}
