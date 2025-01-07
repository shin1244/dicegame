package main

import (
	"dice-game/game"
)

type Game struct {
	SnackLadderPhase *SnakeLadder
	DiceManager      *game.DiceManager
	Phase            uint
	Gold             int
}

func NewGame() *Game {
	diceManager := game.NewDiceManager()
	return &Game{
		SnackLadderPhase: NewSnakeLadderPhase(diceManager),
		DiceManager:      diceManager,
		Phase:            0,
		Gold:             0,
	}
}

func resetDiceChance(setDiceChance []uint, diceChance *[]uint) {
	*diceChance = make([]uint, len(setDiceChance))
	copy(*diceChance, setDiceChance)
}
