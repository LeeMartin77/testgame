package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/assets"
)

func Initialise() Game {
	ass, err := assets.LoadAssets()
	if err != nil {
		panic(err)
	}
	return &GameState{
		assets: ass,
	}
}

type Game interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type GameState struct {
	assets assets.Provider
}

func (g *GameState) Update() error {
	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.assets.ShipsTiles(), &ebiten.DrawImageOptions{})
}

func (g *GameState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
