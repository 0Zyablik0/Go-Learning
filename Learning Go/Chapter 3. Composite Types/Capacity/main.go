package main

import "fmt"

func main() {
	var tetsSlice []int
	fmt.Println(tetsSlice, len(tetsSlice), cap(tetsSlice))
	tetsSlice = append(tetsSlice, 10)
	fmt.Println(tetsSlice, len(tetsSlice), cap(tetsSlice))
	tetsSlice = append(tetsSlice, 20)
	fmt.Println(tetsSlice, len(tetsSlice), cap(tetsSlice))
	tetsSlice = append(tetsSlice, 30)
	fmt.Println(tetsSlice, len(tetsSlice), cap(tetsSlice))
	tetsSlice = append(tetsSlice, 40)
	fmt.Println(tetsSlice, len(tetsSlice), cap(tetsSlice))
	tetsSlice = append(tetsSlice, 50)
	fmt.Println(tetsSlice, len(tetsSlice), cap(tetsSlice))
}
