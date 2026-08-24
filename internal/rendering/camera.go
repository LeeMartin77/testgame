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
	// Radius of the image (distinct from position)
	Rad() float64
	// Source image
	Img() *ebiten.Image
	// rotation in radians
	Rot() float64
}

func (c *Camera) ApplyZoomChange(zc float64) {
	c.Scale += zc
	if c.Scale <= 0.25 {
		c.Scale = 0.25
	}
}

func (c *Camera) RenderToScreen(thing RenderableAsset, screen *ebiten.Image) {
	geom := ebiten.GeoM{}

	c.localTransformationOfAssetImage(thing, &geom)

	c.globalTransformOfAssetImage(thing, screen, &geom)

	screen.DrawImage(thing.Img(), &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (c *Camera) localTransformationOfAssetImage(thing RenderableAsset, geom *ebiten.GeoM) {
	w := float64(thing.Img().Bounds().Dx())
	h := float64(thing.Img().Bounds().Dy())
	geom.Translate(-w/2, -h/2)
	geom.Rotate(thing.Rot())

	// we'll always work with squares but just for safety
	scalar := w
	if h > scalar {
		scalar = h
	}

	op := thing.Rad() / scalar

	scl := c.Scale * op

	geom.Scale(scl, scl)
}

func (c *Camera) globalTransformOfAssetImage(thing RenderableAsset, screen *ebiten.Image, geom *ebiten.GeoM) {
	camoff := geometry.Vec2{
		X: float64(screen.Bounds().Dx()) / 2,
		Y: float64(screen.Bounds().Dy()) / 2,
	}
	pos := thing.Pos().OffsetFrom(c.Pos)
	geom.Translate(
		camoff.X+(pos.X*c.Scale),
		camoff.Y+(pos.Y*c.Scale),
	)
}

func (c *Camera) InView(pos *geometry.Vec2) bool {
	return true
}
