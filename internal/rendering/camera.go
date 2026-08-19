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
	//scl := c.Scale * thing.Scale()
	geom := ebiten.GeoM{}
	geom.Translate(
		camoff.X+pos.X-float64(thing.Img().Bounds().Dx()/2),
		camoff.Y+pos.Y-float64(thing.Img().Bounds().Dy()/2),
	)
	screen.DrawImage(thing.Img(), &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (c *Camera) InView(pos *geometry.Vec2) bool {
	return true
}
