package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/assets"
	"github.com/leemartin77/testgame/internal/input"
	"github.com/leemartin77/testgame/internal/world"
)

func Initialise() Game {
	ass, err := assets.LoadAssets()
	if err != nil {
		panic(err)
	}
	return &GameState{
		assets: ass,
		input:  &input.PlayerInput{},
		world:  world.NewWorld(),
	}
}

type GameState struct {
	assets assets.Provider
	input  *input.PlayerInput
	world  world.World
}

type Game interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

func (g *GameState) Update() error {
	g.input.Update()

	if err := g.world.Update(g.input); err != nil {
		return err
	}

	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *GameState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
