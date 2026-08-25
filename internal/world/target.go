package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/assets"
	"github.com/leemartin77/testgame/internal/geometry"
)

func NewTarget(x, y, vx, vy, rad, rotspd float64, ass assets.Provider) RenderablePhysicsObject {
	return &Target{
		Position: &geometry.Vec2{
			X: x, Y: y,
		},
		Radius: rad,
		Velocity: &geometry.Vec2{
			X: vx, Y: vy,
		},
		RotationSpeed: rotspd,
		Rotation:      0,

		Tile: ass.Target(),
	}
}

type Target struct {
	Position *geometry.Vec2
	Radius   float64

	Velocity      *geometry.Vec2
	RotationSpeed float64
	Rotation      float64

	Tile *ebiten.Image
}

// Img implements [RenderablePhysicsObject].
func (t *Target) Img() *ebiten.Image {
	return t.Tile
}

// Pos implements [PhysicsObject].
func (t *Target) Pos() *geometry.Vec2 {
	return t.Position
}

// Rad implements [PhysicsObject].
func (t *Target) Rad() float64 {
	return t.Radius
}

// Rot implements [PhysicsObject].
func (t *Target) Rot() float64 {
	return t.Rotation
}

// RotSpd implements [PhysicsObject].
func (t *Target) RotSpd() float64 {
	return t.RotationSpeed
}

// SetRot implements [PhysicsObject].
func (t *Target) SetRot(nr float64) {
	t.Rotation = nr
}

// Vel implements [PhysicsObject].
func (t *Target) Vel() *geometry.Vec2 {
	return t.Velocity
}
