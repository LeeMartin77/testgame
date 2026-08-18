package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/input"
)

type World interface {
	Update(input *input.PlayerInput) error
	Draw(screen *ebiten.Image)
}

func NewWorld() World {
	return &WorldState{}
}

type WorldState struct {
}

// Draw implements [World].
func (w *WorldState) Draw(screen *ebiten.Image) {
	return
}

// Update implements [World].
func (w *WorldState) Update(input *input.PlayerInput) error {
	return nil
}
