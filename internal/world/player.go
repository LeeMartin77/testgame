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
	Rotation float64
	Size     float64
	Tile     *ebiten.Image

	Velocity *geometry.Vec2

	ShipPower         float64
	ShipDrag          float64
	ShipRotationSpeed float64
}

func NewPlayer(ass assets.Provider) *Player {
	return &Player{
		Position: &geometry.Vec2{
			X: 0, Y: 0,
		},
		// in Radians
		Rotation: 0,
		Size:     0.5,
		Tile:     ass.PlayerShip(),

		Velocity: &geometry.Vec2{
			X: 0, Y: 0,
		},

		ShipPower:         0.5,
		ShipDrag:          0.03,
		ShipRotationSpeed: 1,
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

// Scale implements [rendering.RenderableAsset].
func (p *Player) Scale() float64 {
	return p.Size
}

// Scale implements [rendering.RenderableAsset].
func (p *Player) Rot() float64 {
	return p.Rotation
}

func (p *Player) Update(input *input.PlayerInput) {
	p.handleInput(input)

	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y

	dragScale := p.ShipDrag * (math.Abs(p.Velocity.X) + math.Abs(p.Velocity.Y))

	p.Velocity.X = p.Velocity.X + (-1 * dragScale * p.Velocity.X)
	p.Velocity.Y = p.Velocity.Y + (-1 * dragScale * p.Velocity.Y)

}

func (p *Player) handleInput(input *input.PlayerInput) {
	speedadd := float64(0)
	if input.Vec.Y > 0 {
		speedadd = input.Vec.Y * p.ShipPower
	}
	p.Velocity.X += speedadd * math.Sin(p.Rotation)
	p.Velocity.Y += -1 * speedadd * math.Cos(p.Rotation)

	p.Rotation += input.Rot * p.ShipRotationSpeed
}
