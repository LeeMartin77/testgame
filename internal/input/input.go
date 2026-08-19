package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/geometry"
)

func NewPlayerInput() *PlayerInput {
	return &PlayerInput{
		Vec: &geometry.Vec2{
			X: 0,
			Y: 0,
		},
	}
}

type PlayerInput struct {
	Vec *geometry.Vec2
}

func (pi *PlayerInput) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		pi.Vec.Y = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {

		pi.Vec.Y = 1
	} else {
		pi.Vec.Y = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		pi.Vec.X = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {

		pi.Vec.X = 1
	} else {

		pi.Vec.X = 0
	}
}
