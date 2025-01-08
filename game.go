package main

import "log"

type Game struct {
	SnackLadderPhase *SnakeLadder
	BattlePhase      *Battle
	DiceManager      *DiceManager
	Phase            uint
	Gold             int
	Life             int
}

func NewGame() *Game {
	return &Game{
		SnackLadderPhase: NewSnakeLadderPhase(),
		BattlePhase:      NewBattlePhase(),
		DiceManager:      NewDiceManager(),
		Phase:            0,
		Gold:             0,
		Life:             2,
	}
}

func resetDiceChance(setDiceChance []uint, diceChance *[]uint) {
	*diceChance = make([]uint, len(setDiceChance))
	copy(*diceChance, setDiceChance)
}

func (g *Game) GameOver() {
	if g.Life == 0 {
		log.Fatal("게임 오버")
	}
}
