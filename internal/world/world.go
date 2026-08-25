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
		targets: []RenderablePhysicsObject{
			NewTarget(46, -54, 0.5, -0.3, 20, -0.004, ass),
			NewTarget(-62, -54, -0.7, 0.4, 30, 0.009, ass),
			NewTarget(37, 0, 0.3, 0.2, 24, -0.007, ass),
		},
	}
}

type WorldState struct {
	assets assets.Provider

	camera *rendering.Camera

	player *Player

	targets []RenderablePhysicsObject
}

// Draw implements [World].
func (w *WorldState) Draw(screen *ebiten.Image) {
	for _, obj := range w.targets {
		w.camera.RenderToScreen(obj, screen)
	}
	w.camera.RenderToScreen(w.player, screen)
	return
}

// Update implements [World].
func (w *WorldState) Update(input *input.PlayerInput) error {
	w.camera.ApplyZoomChange(input.CamZoom)
	w.player.Update(input)

	ApplyPhysics(w.player)
	for _, obj := range w.targets {
		ApplyPhysics(obj)
	}
	return nil
}
