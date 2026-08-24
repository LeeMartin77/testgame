package world

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/assets"
	"github.com/leemartin77/testgame/internal/geometry"
	"github.com/leemartin77/testgame/internal/input"
)

type Player struct {
	Position *geometry.Vec2
	Radius   float64

	Velocity      *geometry.Vec2
	RotationSpeed float64

	Rotation float64
	Tile     *ebiten.Image

	ShipPower      float64
	ShipDrag       float64
	ShipSpeedLimit float64

	ShipRotationPower      float64
	ShipRotationDrag       float64
	ShipRotationSpeedLimit float64
}

func NewPlayer(ass assets.Provider) *Player {
	return &Player{
		Position: &geometry.Vec2{
			X: 0, Y: 0,
		},
		Radius: 30,
		// in Radians
		Rotation: 0,

		Velocity: &geometry.Vec2{
			X: 0, Y: 0,
		},
		RotationSpeed: 0,

		Tile: ass.PlayerShip(),

		ShipRotationPower:      0.001,
		ShipRotationDrag:       3,
		ShipRotationSpeedLimit: 1,

		ShipPower:      0.5,
		ShipDrag:       0.1,
		ShipSpeedLimit: 10,
	}
}

// Img implements [rendering.RenderableAsset].
func (p *Player) Img() *ebiten.Image {
	return p.Tile
}

// Pos implements [rendering.RenderableAsset].
func (p *Player) Pos() *geometry.Vec2 {
	return p.Position
}

func (p *Player) Vel() *geometry.Vec2 {
	return p.Velocity
}

// Scale implements [rendering.RenderableAsset].
func (p *Player) Rad() float64 {
	return p.Radius
}

// Scale implements [rendering.RenderableAsset].
func (p *Player) Rot() float64 {
	return p.Rotation
}

// Scale implements [world.PhysicsObject].
func (p *Player) RotSpd() float64 {
	return p.RotationSpeed
}

func (p *Player) SetRot(n float64) {
	p.Rotation = n
}

// Remember - 60 Tick rate
func (p *Player) Update(input *input.PlayerInput) {

	// Thrust or drag
	if input.Vec.Y > 0 {
		speedadd := input.Vec.Y * p.ShipPower

		speedadd = (1 - ILerp(0, p.ShipSpeedLimit, p.Velocity.Magnitude())) * speedadd

		p.Velocity.X += speedadd * math.Sin(p.Rotation)
		p.Velocity.Y += -1 * speedadd * math.Cos(p.Rotation)
	} else {
		dragScale := ILerp(0, p.ShipSpeedLimit, p.Velocity.Magnitude()) * p.ShipDrag

		p.Velocity.X = p.Velocity.X + (-1 * dragScale * p.Velocity.X)
		p.Velocity.Y = p.Velocity.Y + (-1 * dragScale * p.Velocity.Y)
	}

	// Rotate or drag
	if input.Rot != 0 {
		rotspdadd := input.Rot * p.ShipRotationPower
		rotspdadd = (1 - ILerp(0, p.ShipRotationSpeedLimit, p.RotationSpeed)) * rotspdadd
		p.RotationSpeed += rotspdadd
	} else {
		dragScale := ILerp(0, p.ShipRotationSpeedLimit, math.Abs(p.RotationSpeed)) * p.ShipRotationDrag
		p.RotationSpeed *= 1 - dragScale
		p.RotationSpeed = math.Max(-p.ShipRotationSpeedLimit, math.Min(p.RotationSpeed, p.ShipRotationSpeedLimit))
		if math.Abs(p.RotationSpeed) < 1e-6 {
			p.RotationSpeed = 0
		}
	}
}

// Get the interpolant within a range
func ILerp(a, b, v float64) float64 {
	// https://en.wikipedia.org/wiki/Linear_interpolation
	return (v - a) / (b - a)
}
