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
	Pos() *geometry.Vec2
	Scale() float64
	Img() *ebiten.Image
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
	geom.Scale(scl, scl)
	geom.Translate(
		camoff.X+(pos.X*c.Scale)-((float64((thing.Img().Bounds().Dx()))*scl)/2),
		camoff.Y+(pos.Y*c.Scale)-((float64((thing.Img().Bounds().Dy()))*scl)/2),
	)
	screen.DrawImage(thing.Img(), &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (c *Camera) InView(pos *geometry.Vec2) bool {
	return true
}
