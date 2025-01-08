package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Reward struct {
	Chance []uint
}

var RewardPhaseChance []uint = []uint{0, 0, 0, 1}

func (g *Game) RewardPhaseUpdate() {
	if g.DiceManager.Select != -1 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if g.Reward.Chance[g.DiceManager.Select] > 0 {
				g.DiceManager.Clicked()
			} else {
				g.DiceManager.Select = -1
			}
			return
		}

		// 클릭이 없을 때의 처리
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
