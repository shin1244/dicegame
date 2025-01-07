package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	// g.SnackLadderPhase.SnLPhaseUpdate()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	g.SnackLadderPhase.SnLPhaseDraw(screen)
	g.drawDices(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(960, 640)
	ebiten.SetWindowTitle("Dice Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal("게임 실행 불가")
	}
}
