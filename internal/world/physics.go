package world

import "github.com/leemartin77/testgame/internal/geometry"

// All things shalt be perfect circles
// in classic physics simplification
type PhysicsObject interface {
	// Position in the world
	Pos() *geometry.Vec2
	// Velocity
	Vel() *geometry.Vec2
	// Radius
	Rad() float64
}

func ApplyVelocity(po PhysicsObject) {
	po.Pos().X += po.Vel().X
	po.Pos().Y += po.Vel().Y
}
