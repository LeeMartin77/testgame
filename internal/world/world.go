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

		player: &Player{
			Position: &geometry.Vec2{
				X: 0, Y: 0,
			},
			Size: 0.5,
			Tile: ass.PlayerShip(),
		},
	}
}

type Player struct {
	Position *geometry.Vec2
	Size     float64
	Tile     *ebiten.Image
}

type WorldState struct {
	assets assets.Provider

	camera *rendering.Camera

	player *Player
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

// Draw implements [World].
func (w *WorldState) Draw(screen *ebiten.Image) {
	w.camera.RenderToScreen(w.player, screen)
	return
}

// Update implements [World].
func (w *WorldState) Update(input *input.PlayerInput) error {
	w.player.Position.X += input.Vec.X
	w.player.Position.Y += input.Vec.Y
	return nil
}
