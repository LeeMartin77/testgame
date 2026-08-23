package manager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/game"
	"github.com/leemartin77/testgame/internal/input"
)

func NewMenu() Manager {
	return &ManagerState{
		input: input.NewPlayerInput(),
	}
}

type Manager interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type ManagerState struct {
	game game.Game

	input *input.PlayerInput
}

// Draw implements [Manager].
func (g *ManagerState) Draw(screen *ebiten.Image) {
	if g.game != nil {
		g.game.Draw(screen)
		return
	}
	return
}

// Update implements [Manager].
func (g *ManagerState) Update() error {
	g.input.Update()
	if g.game != nil {
		return g.game.Update(g.input)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// start a game
		newgame, exitchan, err := game.Initialise()
		if err != nil {
			return err
		}
		g.game = newgame
		go func() {
			<-exitchan
			// we wait for it, if it appears, kill it
			g.game = nil
		}()
	}
	return nil
}

func (g *ManagerState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if g.game != nil {

		return g.game.Layout(outsideWidth, outsideHeight)
	}
	return 1280, 800
}
