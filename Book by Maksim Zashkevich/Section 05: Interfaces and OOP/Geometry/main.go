package main

import (
	"errors"
	"fmt"
	"math"
)

type figure interface {
	area() (float64, error)
	perimeter() (float64, error)
	isValid() bool
	getParam() string
}

type circle struct {
	r float64
}

type rectangle struct {
	a, b float64
}

type square struct {
	a float64
}
type triangle struct {
	a, b, c float64
}

func printFigure(f figure) {
	perimeter, err := f.perimeter()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	area, err := f.area()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	parameters := f.getParam()

	fmt.Printf("This is %s with perimeter %.5f and area %.5f\n", parameters, perimeter, area)
}

func main() {

	var f figure = &circle{r: 1}
	printFigure(f)
	f = &circle{r: -1}
	printFigure(f)

	f = &rectangle{a: 2, b: 3}
	printFigure(f)
	f = &rectangle{a: -2, b: 3}
	printFigure(f)
	f = &rectangle{a: 2, b: -3}
	printFigure(f)
	f = &rectangle{a: -2, b: -3}
	printFigure(f)

	f = &square{a: 5}
	printFigure(f)
	f = &square{a: -15}
	printFigure(f)

	f = &triangle{a: 3, b: 4, c: 5}
	printFigure(f)
	f = &triangle{a: 1, b: 1, c: 1}
	printFigure(f)
	f = &triangle{a: 1, b: 1, c: 3}
	printFigure(f)
	f = &triangle{a: 3, b: 1, c: 1}
	printFigure(f)
	f = &triangle{a: 1, b: 3, c: 1}
	printFigure(f)

}

func (c *circle) isValid() bool {
	return c.r >= 0
}

func (c *circle) area() (float64, error) {
	if !c.isValid() {
		return 0, errors.New("the circle is invalid")
	}
	return math.Pi * c.r * c.r, nil
}

func (c *circle) perimeter() (float64, error) {
	if !c.isValid() {
		return 0, errors.New("the circle is invalid")
	}
	return 2 * math.Pi * c.r, nil
}

func (c *circle) getParam() string {
	return fmt.Sprintf("circle of radius: %.2f", c.r)
}

func (r *rectangle) isValid() bool {
	return r.a >= 0 && r.b >= 0
}

func (r *rectangle) area() (float64, error) {
	if !r.isValid() {
		return 0, errors.New("the rectangle is invalid")
	}
	return r.a * r.b, nil
}

func (r *rectangle) perimeter() (float64, error) {
	if !r.isValid() {
		return 0, errors.New("the rectangle is invalid")
	}
	return (r.a + r.b) * 2, nil
}

func (r *rectangle) getParam() string {
	return fmt.Sprintf("rectangle with sides: %.3f and %.3f", r.a, r.b)
}

func (s *square) isValid() bool {
	return s.a >= 0
}

func (s *square) area() (float64, error) {
	if !s.isValid() {
		return 0, errors.New("the square is invalid")
	}
	return s.a * s.a, nil
}

func (s *square) perimeter() (float64, error) {
	if !s.isValid() {
		return 0, errors.New("the square is invalid")
	}
	return s.a * 4, nil
}

func (s *square) getParam() string {
	return fmt.Sprintf("square with side: %.3f", s.a)
}

func (t *triangle) isValid() bool {
	return t.a+t.b > t.c && t.a+t.c > t.b && t.c+t.b > t.a
}

func (t *triangle) area() (float64, error) {
	if !t.isValid() {
		return 0, errors.New("the triangle is invalid")
	}
	p := (t.a + t.b + t.c) / 2
	S := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return S, nil
}

func (t *triangle) perimeter() (float64, error) {
	if !t.isValid() {
		return 0, errors.New("the triangle is invalid")
	}
	return t.a + t.b + t.c, nil
}

func (t *triangle) getParam() string {
	return fmt.Sprintf("triangle with sides: %.3f, %.3f and %.3f", t.a, t.b, t.c)
}
