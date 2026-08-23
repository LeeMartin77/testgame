package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/manager"
)

func main() {
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Testgame")
	if err := ebiten.RunGame(manager.NewMenu()); err != nil {
		log.Fatal(err)
	}
}
