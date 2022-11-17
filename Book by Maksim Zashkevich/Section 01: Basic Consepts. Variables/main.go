package main

import (
	"fmt"
)

func main() {
	var phrase_1 string = "Hello, world!" // comment
	var phrase_2 string
	var phrase_3 = "Hello, world!"
	fmt.Println(phrase_1)
	fmt.Println(phrase_2)
	fmt.Println(phrase_3)

	var number1, number2, number3 int
	var boll_var, float_var, string_var = true, 2.3, "four"
	fmt.Println(number1, number2, number3)
	fmt.Println(boll_var, float_var, string_var)

	/**
		Possible types:
			int, int8, int16, int32, int64, uint8, uint16, uint32, uint64, uintptr, byte
			float32, float64
			complex64, complex128
			bool
			string
	**/

}
