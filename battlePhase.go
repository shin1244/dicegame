package main

import (
	"dice-game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Battle struct {
	BattleType uint
	DiceChance []uint
}

var BattlePhaseChance []uint = []uint{0, 1, 1, 0}

// 아군 주사위를 던지기 전 적 주사위가 남아있으면 안됨
func (g *Game) BattlePhaseUpdate() {}

func (g *Game) BattlePhaseDraw(screen *ebiten.Image) {}

func NewBattlePhase(diceManager *game.DiceManager) *Battle {
	return &Battle{}
}
