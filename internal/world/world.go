package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/assets"
	"github.com/leemartin77/testgame/internal/geometry"
	"github.com/leemartin77/testgame/internal/input"
	"github.com/leemartin77/testgame/internal/rendering"
)

type World interface {
	Update(input *input.PlayerInput) error
	Draw(screen *ebiten.Image)
}

func NewWorld(ass assets.Provider) World {
	return &WorldState{
		assets: ass,

		camera: &rendering.Camera{
			Pos: &geometry.Vec2{
				X: 0, Y: 0,
			},
			Scale: 1,
		},

		player: NewPlayer(ass),
	}
}

type WorldState struct {
	assets assets.Provider

	camera *rendering.Camera

	player *Player
}

// Draw implements [World].
func (w *WorldState) Draw(screen *ebiten.Image) {
	w.camera.RenderToScreen(w.player, screen)
	return
}

// Update implements [World].
func (w *WorldState) Update(input *input.PlayerInput) error {
	w.camera.ApplyZoomChange(input.CamZoom)
	w.player.Update(input)

	ApplyVelocity(w.player)
	return nil
}
