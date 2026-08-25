package world

import (
	"github.com/leemartin77/testgame/internal/geometry"
	"github.com/leemartin77/testgame/internal/rendering"
)

// All things shalt be perfect circles
// in classic physics simplification
type PhysicsObject interface {
	// Position in the world
	Pos() *geometry.Vec2
	// Velocity
	Vel() *geometry.Vec2
	// Radius
	Rad() float64

	RotSpd() float64
	Rot() float64
	SetRot(float64)
}

type RenderablePhysicsObject interface {
	rendering.RenderableAsset
	PhysicsObject
}

func ApplyPhysics(po PhysicsObject) {
	ApplyVelocity(po)
	ApplyRotation(po)
}

func ApplyVelocity(po PhysicsObject) {
	po.Pos().X += po.Vel().X
	po.Pos().Y += po.Vel().Y
}

func ApplyRotation(po PhysicsObject) {
	po.SetRot(po.Rot() + po.RotSpd())
}
