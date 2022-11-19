package shape

import (
	"errors"
	"fmt"
	"math"
)

type Figure interface {
	Area() (float64, error)
	Perimeter() (float64, error)
	GetParam() string

	isValid() bool
}

type Circle struct {
	R float64
}

type Rectangle struct {
	A, B float64
}

type Square struct {
	A float64
}
type Triangle struct {
	A, B, C float64
}

func (c *Circle) isValid() bool {
	return c.R >= 0
}

func (c *Circle) Area() (float64, error) {
	if !c.isValid() {
		return 0, errors.New("the circle is invalid")
	}
	return math.Pi * c.R * c.R, nil
}

func (c *Circle) Perimeter() (float64, error) {
	if !c.isValid() {
		return 0, errors.New("the circle is invalid")
	}
	return 2 * math.Pi * c.R, nil
}

func (c *Circle) GetParam() string {
	return fmt.Sprintf("circle of radius: %.2f", c.R)
}

func (r *Rectangle) isValid() bool {
	return r.A >= 0 && r.B >= 0
}

func (r *Rectangle) Area() (float64, error) {
	if !r.isValid() {
		return 0, errors.New("the rectangle is invalid")
	}
	return r.A * r.B, nil
}

func (r *Rectangle) Perimeter() (float64, error) {
	if !r.isValid() {
		return 0, errors.New("the rectangle is invalid")
	}
	return (r.A + r.B) * 2, nil
}

func (r *Rectangle) GetParam() string {
	return fmt.Sprintf("rectangle with sides: %.3f and %.3f", r.A, r.B)
}

func (s *Square) isValid() bool {
	return s.A >= 0
}

func (s *Square) Area() (float64, error) {
	if !s.isValid() {
		return 0, errors.New("the square is invalid")
	}
	return s.A * s.A, nil
}

func (s *Square) Perimeter() (float64, error) {
	if !s.isValid() {
		return 0, errors.New("the square is invalid")
	}
	return s.A * 4, nil
}

func (s *Square) GetParam() string {
	return fmt.Sprintf("square with side: %.3f", s.A)
}

func (t *Triangle) isValid() bool {
	return t.A+t.B > t.C && t.A+t.C > t.B && t.C+t.B > t.A
}

func (t *Triangle) Area() (float64, error) {
	if !t.isValid() {
		return 0, errors.New("the triangle is invalid")
	}
	p := (t.A + t.B + t.C) / 2
	S := math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	return S, nil
}

func (t *Triangle) Perimeter() (float64, error) {
	if !t.isValid() {
		return 0, errors.New("the triangle is invalid")
	}
	return t.A + t.B + t.C, nil
}

func (t *Triangle) GetParam() string {
	return fmt.Sprintf("triangle with sides: %.3f, %.3f and %.3f", t.A, t.B, t.C)
}
