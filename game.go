package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type Game struct {
	SnackLadder *SnakeLadder
	Battle      *Battle
	Reward      *Reward
	DiceManager *DiceManager
	Phase       uint
	Gold        int
	Life        uint
}

func NewGame() *Game {
	return &Game{
		SnackLadder: NewSnakeLadderPhase(),
		Battle:      NewBattlePhase(),
		Reward:      NewRewardPhase(),
		DiceManager: NewDiceManager(),
		Phase:       0,
		Gold:        0,
		Life:        2,
	}
}

func (g *Game) GameOver() {
	if g.Life == 0 {
		log.Fatal("게임 오버")
	}
}

func (g *Game) DrawStatusBar(screen *ebiten.Image) {
	statusBar := createStatusBar()
	drawStatusBar(screen, statusBar)
	drawStatusText(screen, g.Gold, g.Life)
}

func createStatusBar() *ebiten.Image {
	statusBar := ebiten.NewImage(960, 64)
	statusBar.Fill(colornames.Gray)
	return statusBar
}

func drawStatusBar(screen, statusBar *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	screen.DrawImage(statusBar, opt)
}

func drawStatusText(screen *ebiten.Image, gold int, life uint) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Gold: %d", gold), 20, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Life: %d", life), 20, 50)
}
