package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Reward struct {
	RewardType uint
	DiceChance []int
}

var RewardPhaseChance []uint = []uint{0, 0, 0, 1}

// 보상 획득 gold 증가

func (g *Game) RewardPhaseUpdate() {}

func (g *Game) RewardPhaseDraw(screen *ebiten.Image) {}

func NewRewardPhase() *Reward {
	return &Reward{}
}
