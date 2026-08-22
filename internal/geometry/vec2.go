package geometry

import "math"

type Vec2 struct {
	X, Y float64
}

func (v *Vec2) OffsetFrom(v2 *Vec2) Vec2 {
	return Vec2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

// get total amount of vector as a single line. Always positive
func (v *Vec2) Magnitude() float64 {
	r := math.Abs(v.X) + math.Abs(v.Y)
	if r < 0 {
		return 0
	}
	return r
}
