package structs

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		Radius: radius,
	}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Square struct {
	Width float64
}

func NewSquare(width float64) *Square {
	return &Square{
		Width: width,
	}
}

func (sq *Square) Area() float64 {
	return sq.Width * sq.Width
}

func (sq *Square) Perimeter() float64 {
	return sq.Width * 4
}

func ShareInfo(sp Shape) string {
	return fmt.Sprintf("Area:%v Perimeter:%v", sp.Area(), sp.Perimeter())
}
