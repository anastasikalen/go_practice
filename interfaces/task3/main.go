package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{
		23,
	}
	PrintArea(circle)
	rectangle := Rectangle{
		23, 34,
	}
	PrintArea(rectangle)

}
func PrintArea(s Shape) {
	fmt.Printf("area : %f\n", s.Area()) //КОД ДНЯ 1154
}

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	length, width float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (rect Rectangle) Area() float64 {
	return rect.length * rect.width
}
