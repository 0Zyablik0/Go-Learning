package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	printTriangleSizes(1, 2, 4)
	printTriangleSizes(-1, -1, -1)
	printTriangleSizes(3, 4, 5)
	printTriangleSizes(3, 3, 3)

	printCircleSizes(-1)
	printCircleSizes(0)
	printCircleSizes(1)
	printCircleSizes(2)

}

func calculateTrianglePerimeter(a, b, c float64) (float64, error) {
	if a+b <= c || a+c <= b || b+c <= a {
		err := errors.New("invalid triangle")
		return 0, err
	}
	perimeter := a + b + c

	return perimeter, nil
}

func calculateCirclePerimeter(r float64) (float64, error) {
	if r <= 0 {
		err := errors.New("invalid circle radius")
		return 0, err
	}
	perimeter := 2 * math.Pi * r

	return perimeter, nil
}

func calculateCircleArea(r float64) (float64, error) {
	if r <= 0 {
		err := errors.New("invalid circle radius")
		return 0, err
	}

	area := math.Pi * r * r

	return area, nil
}

func calculateTriangleArea(a, b, c float64) (float64, error) {
	perimeter, err := calculateTrianglePerimeter(a, b, c)
	if err != nil {
		return 0, err
	}
	semiPerimeter := perimeter / 2

	area := math.Sqrt(semiPerimeter * (semiPerimeter - a) * (semiPerimeter - b) * (semiPerimeter - c))

	return area, nil
}

func printTriangleSizes(a, b, c float64) {

	perimeter, err := calculateTrianglePerimeter(a, b, c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	area, err := calculateTriangleArea(a, b, c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Perimeter of the triangle with sides %.2f, %.2f, %.2f is %.5f\n", a, b, c, perimeter)
	fmt.Printf("Area of the triangle with sides %.2f, %.2f, %.2f is %.5f\n\n", a, b, c, area)

}

func printCircleSizes(r float64) {

	perimeter, err := calculateCirclePerimeter(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	area, err := calculateCircleArea(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Perimeter of the circle with radius %.2f  is %.5f\n", r, perimeter)
	fmt.Printf("Area of the circle with radius %.2f  is %.5f\n\n", r, area)

}
