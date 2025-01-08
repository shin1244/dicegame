package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.GameOver()
	g.DiceManager.HandleInput()

	switch g.Phase {
	case 0:
		g.SnakeLadderPhaseUpdate()
	case 1:
		g.BattlePhaseUpdate()
	case 2:
		g.RewardPhaseUpdate()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	g.DrawStatusBar(screen)

	switch g.Phase {
	case 0:
		g.SnakeLadderPhaseDraw(screen)
	case 1:
		g.BattlePhaseDraw(screen)
	}
	g.DiceManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(960, 704)
	ebiten.SetWindowTitle("Dice Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := NewGame()
	g.SnackLadder.Chance = []uint{1, 0, 0, 0}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal("게임 실행 불가")
	}
}
