package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Redius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Redius, 2)
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

// Old funcations without interface
func Perimeter(ractangle Rectangle) float64 {
	return 2 * (ractangle.Width + ractangle.Height)
}

func Area(ractangle Rectangle) float64 {
	return ractangle.Width * ractangle.Height
}
