package main

import (
	"dice-game/tile"
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
		g.Life--
	}
	if tile.NowTile(g.SnackLadder.Tilemap, g.SnackLadder.Player).Type == 5 {
		g.Reward.IsBoss = true
	}
	g.Phase = 2
}

func initDiceChance(setDiceChance []uint, diceChance *[]uint) {
	*diceChance = make([]uint, len(setDiceChance))
	copy(*diceChance, setDiceChance)
}

func (g *Game) ChangeCursedTile() {
	for i := 0; i < 3; i++ {
		tileIndex := g.SnackLadder.CurseTile + uint(i)
		row := tileIndex / 10
		col := tileIndex % 10

		g.SnackLadder.Tilemap[row][col].Image = tile.CursedTile.Image
		g.SnackLadder.Tilemap[row][col].Type = tile.CursedTile.Type
	}
	g.SnackLadder.CurseTile += 3
}
