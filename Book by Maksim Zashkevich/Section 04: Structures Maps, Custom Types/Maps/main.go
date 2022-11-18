package main

import (
	"fmt"
)

func main() {
	emptyMap()
	basicMap()
	deleteKey()

}

func emptyMap() {
	fmt.Println("\ntestMap after make(map[string]int):")
	var testMap map[string]int
	fmt.Println(testMap)
	fmt.Println("testMap == nil -> ", testMap == nil)

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("\ntestMap after make(map[string]int):")
	testMap = make(map[string]int)
	fmt.Println(testMap)
	fmt.Println("testMap == nil -> ", testMap == nil)
	fmt.Println("----------------------------------------------------------------")
}

func basicMap() {
	var testMap = make(map[string]int)
	testMap["one"] = 1
	testMap["two"] = 2
	testMap["three"] = 3
	fmt.Println("\nSimple nonempty map:")
	for key, val := range testMap {
		fmt.Printf("\t%s - %d\n", key, val)
	}
	fmt.Println("----------------------------------------------------------------")

	fmt.Println("\nAccess to nonexistent key:")

	fmt.Printf("\t four - %d\n", testMap["four"]) // nil value

	_, exist := testMap["four"]
	fmt.Printf("Key \"four\"  exists in map: %t\n", exist) // nil value

	fmt.Println("----------------------------------------------------------------")

}

func deleteKey() {
	var testMap1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	var testMap2 map[string]int

	fmt.Println("\nAfter initialization")
	fmt.Printf("First map: %#v\n", testMap1)
	fmt.Printf("Second map: %#v\n", testMap2)

	fmt.Println("----------------------------------------------------------------")

	testMap2 = testMap1 // only pointer is copied

	delete(testMap2, "two")
	delete(testMap1, "three")

	fmt.Println("\nAfter deletes")
	fmt.Printf("First map: %#v\n", testMap1)
	fmt.Printf("Second map: %#v\n", testMap2)

	fmt.Println("----------------------------------------------------------------")

}
