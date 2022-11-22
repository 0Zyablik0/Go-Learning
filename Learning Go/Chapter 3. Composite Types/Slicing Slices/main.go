package main

import "fmt"

func main() {
	var testSlice = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(testSlice)
	var testSliceBeginning = testSlice[:3]
	fmt.Println(testSliceBeginning)
	var testSliceEnd = testSlice[4:]
	fmt.Println(testSliceEnd)
	var testSliceMiddle = testSlice[2:5]
	fmt.Println(testSliceMiddle)
	var testWholeSlice = testSlice[:]
	fmt.Println(testWholeSlice)

	// slices share memory:

	testSliceBeginning[1] = -1
	testSliceEnd[0] = -4
	testSliceMiddle[0] = -2
	fmt.Println(testSlice)
	fmt.Println(testSliceBeginning)
	fmt.Println(testSliceEnd)
	fmt.Println(testSliceMiddle)
	fmt.Println(testWholeSlice)

	// confusion with append:
	fmt.Println("----------------------------------------------------------------")
	var x = []int{1, 2, 3, 4}
	var y = x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	// more confusion:
	fmt.Println("----------------------------------------------------------------")

	x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y = x[:2]  // use x[:2:2]
	z := x[2:] // use x[:2:4]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

	// slices from array

	var testArray = [4]int{1, 2, 3, 4}
	testSliceBeginning = testArray[:2]
	testSliceEnd = testArray[2:]
	testArray[0] = 10
	fmt.Println("testArray = ", testArray)
	fmt.Println("testSliceBeginning = ", testSliceBeginning)
	fmt.Println("testSliceEnd = ", testSliceEnd)

}
