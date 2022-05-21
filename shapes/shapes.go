package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base   float64
	A      float64
	B      float64
	C      float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2.0) * math.Pi
}

func (c Circle) Perimeter() float64 {
	return (c.Radius) * 2
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}
