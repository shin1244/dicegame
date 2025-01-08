package main

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Battle struct {
	BattleType  uint
	DiceChance  []uint
	FriendScore int
	EnemyScore  int
}

var BattlePhaseChance []uint = []uint{0, 1, 1, 0}

// 아군 주사위를 던지기 전 적 주사위가 남아있으면 안됨
func (g *Game) BattlePhaseUpdate() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.DiceManager.SelectedDices != -1 {
		if g.BattlePhase.DiceChance[2] > 0 {
			if g.DiceManager.SelectedDices == 2 {
				g.DiceManager.Clicked()
			} else {
				g.DiceManager.SelectedDices = -1
			}
		} else if g.BattlePhase.DiceChance[1] > 0 {
			if g.DiceManager.SelectedDices == 1 {
				g.DiceManager.Clicked()
			} else {
				g.DiceManager.SelectedDices = -1
			}
		}
	} else if g.DiceManager.SelectedDices != -1 {
		g.BattlePhase.DiceChance[g.DiceManager.SelectedDices] -= 1
		if g.DiceManager.SelectedDices == 1 {
			for _, dice := range g.DiceManager.MyDices.FriendlyDice {
				g.BattlePhase.FriendScore += dice.Sides[dice.Val]
			}
		} else {
			for _, dice := range g.DiceManager.MyDices.EnemyDice {
				g.BattlePhase.EnemyScore += dice.Sides[dice.Val]
			}
		}
		g.DiceManager.NotClicked()
		if g.BattlePhase.BattleType == 1 {
			fmt.Println("보다 작으면 아군 승리")
		} else {
			fmt.Println("보다 크면 아군 승리")
		}
	}
	if g.BattlePhase.DiceChance[1] == 0 && g.BattlePhase.DiceChance[2] == 0 {
		switch g.BattlePhase.BattleType {
		case 0:
			if g.BattlePhase.FriendScore > g.BattlePhase.EnemyScore {
				g.initSnakeLadderPhase(true)
			} else {
				g.initSnakeLadderPhase(false)
			}
		case 1:
			if g.BattlePhase.FriendScore < g.BattlePhase.EnemyScore {
				g.initSnakeLadderPhase(true)
			} else {
				g.initSnakeLadderPhase(false)
			}
		}
	}
}

func (g *Game) BattlePhaseDraw(screen *ebiten.Image) {}

func NewBattlePhase() *Battle {
	return &Battle{
		BattleType: 0,
		DiceChance: BattlePhaseChance,
	}
}

func (g *Game) initBattlePhase() {
	resetDiceChance(BattlePhaseChance, &g.BattlePhase.DiceChance)
	g.BattlePhase.FriendScore = 0
	g.BattlePhase.EnemyScore = 0
	g.BattlePhase.BattleType = uint(rand.Intn(2))
	g.Phase = 1
}
