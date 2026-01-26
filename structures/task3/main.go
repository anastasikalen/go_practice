package main

import "fmt"

func main() {
	rec := Rectangle{Width: 10, Height: 5}
	rec2 := Rectangle{Width: 17, Height: 23}
	fmt.Println(rec.Area())
	fmt.Println(rec.Perimeter())
	fmt.Println(rec2.Area())
	fmt.Println(rec2.Perimeter())
}

type Rectangle struct {
	Width  float64
	Height float64
} //КОД ДНЯ 9426

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
