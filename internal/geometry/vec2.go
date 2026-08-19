package geometry

type Vec2 struct {
	X, Y float64
}

func (v *Vec2) OffsetFrom(v2 *Vec2) Vec2 {
	return Vec2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}
