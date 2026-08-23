package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/assets"
	"github.com/leemartin77/testgame/internal/input"
	"github.com/leemartin77/testgame/internal/world"
)

func Initialise() (Game, chan struct{}, error) {
	ass, err := assets.LoadAssets()
	if err != nil {
		return nil, nil, err
	}
	exitchan := make(chan struct{})
	return &GameState{
		assets: ass,
		input:  input.NewPlayerInput(),
		world:  world.NewWorld(ass),

		exit: exitchan,
	}, exitchan, nil
}

type GameState struct {
	assets assets.Provider
	input  *input.PlayerInput
	world  world.World

	exit chan struct{}
}

type Game interface {
	Update(*input.PlayerInput) error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

func (g *GameState) Update(input *input.PlayerInput) error {
	if input.Exit {
		// this should end up vaguely blocking but it's a good thing
		g.exit <- struct{}{}
		return nil
	}
	if err := g.world.Update(input); err != nil {
		return err
	}

	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *GameState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 800
}
