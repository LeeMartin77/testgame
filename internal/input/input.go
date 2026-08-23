package input

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/geometry"
)

type PlayerInput struct {
	Vec *geometry.Vec2
	Rot float64

	CamZoom float64

	Exit bool

	Controls ConfiguredControls
}

func (pi *PlayerInput) CopyOf() PlayerInput {
	return PlayerInput{
		Vec:     pi.Vec.Clone(),
		Rot:     pi.Rot,
		CamZoom: pi.CamZoom,
		Exit:    pi.Exit,
	}
}

func (pi *PlayerInput) HasChangedFrom(pi2 *PlayerInput) bool {
	if pi2 == nil {
		// it's changed if it's never happened, duh
		return true
	}
	return !pi.Vec.Equal(pi2.Vec) || pi.Rot != pi2.Rot || pi.CamZoom != pi2.CamZoom || pi.Exit != pi2.Exit
}

func NewPlayerInput() *PlayerInput {
	return &PlayerInput{
		Vec: &geometry.Vec2{
			X: 0,
			Y: 0,
		},
		Rot:     0,
		CamZoom: 0,
		Exit:    false,

		Controls: ConfiguredControls{
			"accelerate":  {ebiten.KeyW},
			"decelerate":  {ebiten.KeyS},
			"track_left":  {ebiten.KeyA},
			"track_right": {ebiten.KeyD},
			"spin_ccw":    {ebiten.KeyLeft},
			"spin_cw":     {ebiten.KeyRight},

			"up":   {ebiten.KeyUp},
			"down": {ebiten.KeyDown},

			"exitgame": {ebiten.KeyEscape},
		},
	}
}

type ConfiguredControls = map[string][]ebiten.Key

func Valid(cnf ConfiguredControls) bool {
	for _, cl := range ControlList {
		vl, ok := cnf[cl]
		if !ok || len(vl) == 0 {
			return false
		}
	}
	return true
}

func Update(cnf ConfiguredControls, cmd string, newkey ebiten.Key) ConfiguredControls {
	// prune from any existing controls
	for k, v := range cnf {
		i := slices.Index(v, newkey)
		if i > -1 {
			cnf[k] = append(v[:i], v[i+1:]...)
		}
	}
	// add to control
	cnf[cmd] = append(cnf[cmd], newkey)
	return cnf
}

var ControlList = []string{
	"accelerate",
	"decelerate",
	"track_left",
	"track_right",
	"spin_ccw",
	"spin_cw",

	"up",
	"down",

	"exitgame",
}

var ControlLabels = map[string]string{
	"accelerate":  "Accelerate",
	"decelerate":  "Decelerate",
	"track_left":  "Track Left",
	"track_right": "Track Right",
	"spin_ccw":    "Spin Counterclockwise",
	"spin_cw":     "Spin Clockwise",

	"up":   "Cam Zoom In",
	"down": "Cam Zoom Out",

	"exitgame": "Exit Game",
}

var controlactions = map[string]func(pi *PlayerInput){
	"accelerate": func(pi *PlayerInput) {
		pi.Vec.Y += 1
	},
	"decelerate": func(pi *PlayerInput) {
		pi.Vec.Y -= 1
	},
	"track_left": func(pi *PlayerInput) {
		pi.Vec.X -= 1
	},
	"track_right": func(pi *PlayerInput) {
		pi.Vec.X += 1
	},
	"spin_ccw": func(pi *PlayerInput) {
		pi.Rot -= 1

	},
	"spin_cw": func(pi *PlayerInput) {
		pi.Rot += 1
	},

	"up": func(pi *PlayerInput) {
		pi.CamZoom = 0.25
	},
	"down": func(pi *PlayerInput) {
		pi.CamZoom = -0.25
	},
	"exitgame": func(pi *PlayerInput) {
		pi.Exit = true
	},
}

func (pi *PlayerInput) reset() {
	// makes it easier
	pi.Exit = false
	pi.Vec.Y = 0
	pi.Vec.X = 0
	pi.Rot = 0
	pi.CamZoom = 0
}

func (pi *PlayerInput) Update() {
	pi.reset()

	for _, ctrl := range ControlList {
		btns := pi.Controls[ctrl]
		for _, btn := range btns {
			if ebiten.IsKeyPressed(btn) {
				controlactions[ctrl](pi)
				break // we only register once otherwise hell breaks out
			}
		}
	}
}
