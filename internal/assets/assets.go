package assets

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Provider interface {
	ShipsTiles() *ebiten.Image
	PlayerShip() *ebiten.Image
}

type providerState struct {
	ships *ebiten.Image
	tiles *ebiten.Image
}

// PlayerShip implements [Provider].
func (p *providerState) PlayerShip() *ebiten.Image {
	ps := p.ships.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
	return ps
}

// ShipsTiles implements [Provider].
func (p *providerState) ShipsTiles() *ebiten.Image {
	return p.ships
}

//go:embed kenny_pixel_shmup
var f embed.FS

func LoadAssets() (Provider, error) {
	ships, _, err := ebitenutil.NewImageFromFileSystem(f, "kenny_pixel_shmup/ships_packed.png")
	if err != nil {
		return nil, err
	}

	tiles, _, err := ebitenutil.NewImageFromFileSystem(f, "kenny_pixel_shmup/tiles_packed.png")
	if err != nil {
		return nil, err
	}

	return &providerState{
		ships: ships,
		tiles: tiles,
	}, nil
}
