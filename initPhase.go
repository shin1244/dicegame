package main

import (
	"math/rand"
)

func (g *Game) initBattlePhase() {
	initDiceChance(BattlePhaseChance, &g.Battle.Chance)
	g.Battle.FriendScore = 0
	g.Battle.EnemyScore = 0
	g.Battle.BattleType = uint(rand.Intn(2))
	g.Phase = 1
}

func (g *Game) initSnakeLadderPhase() {
	initDiceChance(SnakeLadderPhaseChance, &g.SnackLadder.Chance)
	g.Phase = 0
}

func (g *Game) initRewardPhase(win bool) {
	initDiceChance(RewardPhaseChance, &g.Reward.Chance)
	if !win {
		g.Life -= 1
	}
	g.Phase = 2
}

func initDiceChance(setDiceChance []uint, diceChance *[]uint) {
	*diceChance = make([]uint, len(setDiceChance))
	copy(*diceChance, setDiceChance)
}
