package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Reward struct {
	Chance []uint
}

var RewardPhaseChance []uint = []uint{0, 0, 0, 1}

// 보상 획득 gold 증가

func (g *Game) RewardPhaseUpdate() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.DiceManager.Select != -1 {
		if g.Reward.Chance[g.DiceManager.Select] > 0 {
			g.DiceManager.Clicked()
		} else {
			g.DiceManager.Select = -1
		}
	} else if g.DiceManager.Select != -1 {
		for _, dice := range g.DiceManager.MyDices.RewardDice {
			g.Gold += dice.Sides[dice.Val]
		}
		g.DiceManager.NotClicked()
		g.initSnakeLadderPhase()
	}
}

func (g *Game) RewardPhaseDraw(screen *ebiten.Image) {}

func NewRewardPhase() *Reward {
	return &Reward{}
}
