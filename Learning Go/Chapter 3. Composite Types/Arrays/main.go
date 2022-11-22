package main

import "fmt"

func main() {

	var testArray [3]int // array of 3 int elements
	fmt.Println(testArray)

	var initializedTestArray = [3]int{10, 20, 30}
	fmt.Println(initializedTestArray)

	var sparseTestArray = [12]int{1, 2, 3, 4, 5}
	fmt.Println(sparseTestArray)

	var guessedSizeTestArray = [...]int{1, 2, 3}
	fmt.Println(guessedSizeTestArray)

	var multidimensionalTestArray [2][3]int
	fmt.Println(multidimensionalTestArray)

	testArray[0] = -5
	fmt.Println(testArray)

	fmt.Println(len(testArray))

}
