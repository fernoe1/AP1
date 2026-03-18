package model

import "math"

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Area() float64 {
	sP := (t.A + t.B + t.C) / 2.0
	return math.Sqrt(sP * (sP - t.A) * (sP - t.B) * (sP - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}
