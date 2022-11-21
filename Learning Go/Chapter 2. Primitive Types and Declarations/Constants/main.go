package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)
const z = 10 * 20

func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	const untyped = 10
	var int_v int = untyped
	var float_v float64 = untyped
	var byte_v byte = untyped
	fmt.Printf("type of int_v: %T, type of float_v: %T, type of byte_v: %T", int_v, float_v, byte_v)

	ü := 10
	fmt.Println(ü)
}
