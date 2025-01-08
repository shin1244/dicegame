package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.GameOver()
	g.DiceManager.HandleInput()
	if g.Phase == 0 {
		g.SnakeLadderPhaseUpdate()
	} else if g.Phase == 1 {
		g.BattlePhaseUpdate()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	if g.Phase == 0 {
		g.SnakeLadderPhaseDraw(screen)
	}
	g.DiceManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(960, 640)
	ebiten.SetWindowTitle("Dice Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := NewGame()
	g.SnackLadderPhase.DiceChance = []uint{1, 0, 0, 0}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal("게임 실행 불가")
	}
}
