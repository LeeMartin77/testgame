package manager

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/leemartin77/testgame/internal/game"
	"github.com/leemartin77/testgame/internal/input"
)

func NewMenu() Manager {
	return &ManagerState{
		input: input.NewPlayerInput(),

		menuselected: "playsolo",

		menu: homeoptions,
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

	menuselected string
	menu         []string

	lastinput input.PlayerInput
}

var options = map[string]MenuOption{
	"playsolo": {"Play Solo", func(ms *ManagerState) error {
		return ms.startgame()
	}},
	"settings": {"Settings", func(ms *ManagerState) error {
		ms.menuselected = "home"
		ms.menu = settingoptions
		return nil
	}},
	"home": {"Home", func(ms *ManagerState) error {
		ms.menuselected = "playsolo"
		ms.menu = homeoptions
		return nil
	}},
}

var homeoptions = []string{"playsolo", "settings"}
var settingoptions = []string{"home"}

type MenuOption struct {
	label  string
	action func(ms *ManagerState) error
}

// Draw implements [Manager].
func (g *ManagerState) Draw(screen *ebiten.Image) {
	if g.game != nil {
		g.game.Draw(screen)
		return
	}
	menustring := "TestGame \n\n"
	for _, opt := range g.menu {
		if g.menuselected == opt {
			menustring += "> "
		} else {
			menustring += "  "
		}
		menustring += options[opt].label + "\n"
	}
	ebitenutil.DebugPrint(screen, menustring)
	return
}

// Update implements [Manager].
func (g *ManagerState) Update() error {
	g.input.Update()
	if g.game != nil {
		return g.game.Update(g.input)
	}

	// we should only do something if the input state changes
	// (trying to capture the idea of "just pressed")
	if !g.input.HasChangedFrom(&g.lastinput) {
		return nil
	}
	g.lastinput = g.input.CopyOf()

	if g.input.Rot > 0 {
		opt, ok := options[g.menuselected]
		if ok {
			return opt.action(g)
		}
	}
	menuidx := slices.Index(g.menu, g.menuselected)
	menulen := len(g.menu)
	if g.input.CamZoom > 0 {
		menuidx += 1
	} else if g.input.CamZoom < 0 {
		menuidx -= 1
	}

	if menuidx > menulen-1 {
		menuidx = 0
	}
	if menuidx < 0 {
		menuidx = menulen - 1
	}

	g.menuselected = g.menu[menuidx]

	return nil
}

func (g *ManagerState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if g.game != nil {
		return g.game.Layout(outsideWidth, outsideHeight)
	}
	return 1280, 800
}

func (g *ManagerState) startgame() error {
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
	return nil
}
