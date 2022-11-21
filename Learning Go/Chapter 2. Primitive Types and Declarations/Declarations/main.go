package main

import "fmt"

func main() {

	// Declaration with var
	var x int = 10
	var y = 10 // type guessed at compile time
	var z int
	var x1, x2 int
	var y1, y2 = 10, "string"

	fmt.Printf("x: type %T value %#v\n", x, x)
	fmt.Printf("y: type %T value %#v\n", y, y)
	fmt.Printf("z: type %T value %#v\n", z, z)
	fmt.Printf("x1: type %T value %#v\n", x1, x1)
	fmt.Printf("x2: type %T value %#v\n", x2, x2)
	fmt.Printf("y1: type %T value %#v\n", y1, y1)
	fmt.Printf("y2: type %T value %#v\n", y2, y2)
	var (
		z1     int
		z2         = 20
		z3     int = 30
		z4, z5     = 40, "hello"
		z6, z7 string
	)

	fmt.Printf("z1: type %T value %#v\n", z1, z1)
	fmt.Printf("z2: type %T value %#v\n", z2, z2)
	fmt.Printf("z3: type %T value %#v\n", z3, z3)
	fmt.Printf("z4: type %T value %#v\n", z4, z4)
	fmt.Printf("z5: type %T value %#v\n", z5, z5)
	fmt.Printf("z6: type %T value %#v\n", z6, z6)
	fmt.Printf("z7: type %T value %#v\n", z7, z7)

	// Short declaration

	a1 := 1
	a2, a3 := 30, "hello"
	fmt.Printf("a1: type %T value %#v\n", a1, a1)
	fmt.Printf("a2: type %T value %#v\n", a2, a2)
	fmt.Printf("a3: type %T value %#v\n", a3, a3)

}
