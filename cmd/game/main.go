package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leemartin77/testgame/internal/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Testgame")
	if err := ebiten.RunGame(game.Initialise()); err != nil {
		log.Fatal(err)
	}
}
