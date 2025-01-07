package main

import (
	"dice-game/game"
)

type Game struct {
	SnackLadderPhase *SnakeLadder
	DiceManager      *game.DiceManager
	Phase            uint
}

func NewGame() *Game {
	diceManager := game.NewDiceManager()
	return &Game{
		SnackLadderPhase: NewSnakeLadderPhase(diceManager),
		DiceManager:      diceManager,
		Phase:            0,
	}
}

func resetDiceChance(diceChance []int, diceChance2 *[]int) {
	*diceChance2 = make([]int, len(diceChance))
	copy(*diceChance2, diceChance)
}
