package manager

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/leemartin77/testgame/internal/game"
	"github.com/leemartin77/testgame/internal/input"
)

func NewMenu() Manager {
	return &ManagerState{
		input: input.NewPlayerInput(),

		menuselected: "playsolo",

		menu:    homeoptions,
		options: defaultoptions,

		waitingforinput:     false,
		waitingforinputchan: make(chan ebiten.Key),
	}
}

var homeoption = MenuOption{"Home", func(ms *ManagerState) error {
	if input.Valid(ms.input.Controls) {
		ms.menuselected = "playsolo"
		ms.menu = homeoptions
	}
	return nil
}}

var settingsoption = MenuOption{"Settings", func(ms *ManagerState) error {
	ms.menuselected = "home"
	ms.menu = settingoptions
	return nil
}}

var defaultoptions = map[string]MenuOption{
	"playsolo": {"Play Solo", func(ms *ManagerState) error {
		return ms.startgame()
	}},
	"settings": settingsoption,
	"home":     homeoption,
	"configurecontrols": {"Configure Controls", func(ms *ManagerState) error {
		ms.menuselected = input.ControlList[0]
		ms.menu = input.ControlList
		ms.menu = append(ms.menu, "settings")

		for _, k := range input.ControlList {
			ms.options[k] = MenuOption{
				label: "Set " + input.ControlLabels[k],
				action: func(mss *ManagerState) error {
					mss.waitingforinput = true
					go func() {
						newkey := <-mss.waitingforinputchan
						input.Update(mss.input.Controls, k, newkey)
						mss.waitingforinput = false
					}()
					return nil
				},
			}
		}

		ms.options["settings"] = settingsoption
		return nil
	}},
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

	options map[string]MenuOption

	waitingforinput     bool
	waitingforinputchan chan ebiten.Key

	lastinput input.PlayerInput
}

var homeoptions = []string{"playsolo", "settings"}
var settingoptions = []string{"configurecontrols", "home"}

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
		menustring += g.options[opt].label + "\n"
	}
	ebitenutil.DebugPrint(screen, menustring)
	ebitenutil.DebugPrintAt(screen, g.controlsummary(), 200, 40)
	return
}

func (g *ManagerState) controlsummary() string {
	str := ""
	for _, ctrl := range input.ControlList {
		keys := g.input.Controls[ctrl]
		str += input.ControlLabels[ctrl] + ": "
		if len(keys) == 0 {
			str += "none! (this is a problem)"
		} else {
			for i, k := range keys {
				if i > 0 {
					str += ", "
				}
				str += k.String()
			}
		}
		str += "\n"
	}
	if !input.Valid(g.input.Controls) {
		str += "\n\n Controls invalid - you cannot leave til all controls have buttons"
	}
	return str
}

// Update implements [Manager].
func (g *ManagerState) Update() error {
	g.input.Update()
	if g.game != nil {
		return g.game.Update(g.input)
	}

	// special state where we're just waiting for the next input
	if g.waitingforinput {
		justpressed := inpututil.AppendJustPressedKeys([]ebiten.Key{})
		if len(justpressed) > 0 {
			g.waitingforinput = false
			g.waitingforinputchan <- justpressed[0]
		}
	}

	// we should only do something if the input state changes
	// (trying to capture the idea of "just pressed")
	if !g.input.HasChangedFrom(&g.lastinput) {
		return nil
	}
	g.lastinput = g.input.CopyOf()

	if g.input.Rot > 0 {
		opt, ok := g.options[g.menuselected]
		if ok {
			return opt.action(g)
		}
	}
	menuidx := slices.Index(g.menu, g.menuselected)
	menulen := len(g.menu)
	if g.input.CamZoom > 0 {
		menuidx -= 1
	} else if g.input.CamZoom < 0 {
		menuidx += 1
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
