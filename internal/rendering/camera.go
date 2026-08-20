package rendering

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/geometry"
)

type Camera struct {
	Pos   *geometry.Vec2
	Scale float64
}

type RenderableAsset interface {
	// Position in the world
	Pos() *geometry.Vec2
	// Scale of the image (distinct from position)
	Scale() float64
	// Source image
	Img() *ebiten.Image
	// rotation in radians
	Rot() float64
}

func (c *Camera) RenderToScreen(thing RenderableAsset, screen *ebiten.Image) {
	// screen size

	// align
	camoff := geometry.Vec2{
		X: float64(screen.Bounds().Dx()) / 2,
		Y: float64(screen.Bounds().Dy()) / 2,
	}

	pos := thing.Pos().OffsetFrom(c.Pos)
	geom := ebiten.GeoM{}
	scl := c.Scale * thing.Scale()
	w := float64(thing.Img().Bounds().Dx())
	h := float64(thing.Img().Bounds().Dy())
	geom.Translate(-w/2, -h/2)
	geom.Rotate(thing.Rot())
	geom.Scale(scl, scl)
	geom.Translate(
		camoff.X+(pos.X*c.Scale),
		camoff.Y+(pos.Y*c.Scale),
	)
	screen.DrawImage(thing.Img(), &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (c *Camera) InView(pos *geometry.Vec2) bool {
	return true
}
