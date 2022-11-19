package main

import (
	"fmt"

	"github/0Zyablik0/go/shape"
)

func main() {
	var f shape.Figure = &shape.Circle{R: 1}
	printFigure(f)
	f = &shape.Circle{R: -1}
	printFigure(f)

	f = &shape.Rectangle{A: 2, B: 3}
	printFigure(f)
	f = &shape.Rectangle{A: -2, B: 3}
	printFigure(f)
	f = &shape.Rectangle{A: 2, B: -3}
	printFigure(f)
	f = &shape.Rectangle{A: -2, B: -3}
	printFigure(f)

	f = &shape.Square{A: 5}
	printFigure(f)
	f = &shape.Square{A: -15}
	printFigure(f)

	f = &shape.Triangle{A: 3, B: 4, C: 5}
	printFigure(f)
	f = &shape.Triangle{A: 1, B: 1, C: 1}
	printFigure(f)
	f = &shape.Triangle{A: 1, B: 1, C: 3}
	printFigure(f)
	f = &shape.Triangle{A: 3, B: 1, C: 1}
	printFigure(f)
	f = &shape.Triangle{A: 1, B: 3, C: 1}
	printFigure(f)

}

func printFigure(f shape.Figure) {
	perimeter, err := f.Perimeter()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	area, err := f.Area()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	parameters := f.GetParam()

	fmt.Printf("This is %s with perimeter %.5f and area %.5f\n", parameters, perimeter, area)
}
