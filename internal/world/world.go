package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/assets"
	"github.com/leemartin77/testgame/internal/input"
)

type World interface {
	Update(input *input.PlayerInput) error
	Draw(screen *ebiten.Image)
}

func NewWorld(ass assets.Provider) World {
	return &WorldState{
		assets: ass,
	}
}

type WorldState struct {
	assets assets.Provider
}

// Draw implements [World].
func (w *WorldState) Draw(screen *ebiten.Image) {
	return
}

// Update implements [World].
func (w *WorldState) Update(input *input.PlayerInput) error {
	return nil
}
