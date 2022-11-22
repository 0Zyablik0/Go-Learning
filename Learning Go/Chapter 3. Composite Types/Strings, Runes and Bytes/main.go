package main

import "fmt"

func main() {
	var s string = "😁 - smile"
	var b = s[0] // type byte
	fmt.Printf("s = %s; len(s) = %d;\n", s, len(s))
	fmt.Printf("type of b: %T \n", b)

	var x int = 65
	var y string = string(x)
	fmt.Println(y) // A, not 65

	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
