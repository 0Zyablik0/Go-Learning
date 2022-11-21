package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// Literals
	binary := 0b010001
	octal := 0o1234
	hexadecimal := 0xACBD
	underscored := 1_234_5678
	fp := 1.2345
	efp := 12.5e-2
	hex_efp := 0xABp-10

	fmt.Println("0b010001 is: \t", binary)
	fmt.Println("0o1234 is: \t", octal)
	fmt.Println("0xACBD is: \t", hexadecimal)
	fmt.Println("1_234_5678 is: \t", underscored)
	fmt.Println("1.2345 is: \t", fp)
	fmt.Println("12.5e-2 is: \t", efp)
	fmt.Println("0xABCp-2 is: \t", hex_efp)

	// Rune Literals
	unicode_rune := 'a'
	octal_rune := '\141'
	hex_rune_8_bit := '\x61'
	hex_rune_16_bit := '\U00000061'

	fmt.Printf("'a': type %T, value: %c \n", unicode_rune, unicode_rune)
	fmt.Printf("'\\141': type %T, value: %c \n", octal_rune, octal_rune)
	fmt.Printf("'\\x61': type %T, value: %c \n", hex_rune_8_bit, hex_rune_8_bit)
	fmt.Printf("'\\U00000061': type %T, value: %c \n", hex_rune_16_bit, hex_rune_16_bit)

	string_1 := "Greetings and\n\"Salutations\""
	string_2 :=
		`Greetings and
"Salutations`

	fmt.Println(string_1)
	fmt.Println(string_2)

	// Booleans
	var flag_1 bool
	var flag_2 = true
	fmt.Printf("flag_1: %t, flag_2: %t\n", flag_1, flag_2)

	/* Numeric types:
	int8, int16, int32, int64
	uint8, uint16, uint32, uint64
	byte == uint8
	int == int32 /int64 depending on CPU
	uint == uint32 / uint64 depending on CPU
	rune == int32
	float32, float 64
	complex64, complex128

	*/

	// Complex numbers
	z1 := 1 + 2i
	z2 := -1.4 + 7i
	fmt.Println(z1+z2, z1-z2, z1*z2, z1/z2)
	fmt.Println(real(z1), imag(z2))
	fmt.Println(cmplx.Abs(z1 * z2))

}
