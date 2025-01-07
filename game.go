package main

type Game struct {
	SnackLadderPhase *SnakeLadder
	DiceManager      *DiceManager
	Phase            uint
	Gold             int
}

func NewGame() *Game {
	return &Game{
		SnackLadderPhase: NewSnakeLadderPhase(),
		DiceManager:      NewDiceManager(),
		Phase:            0,
		Gold:             0,
	}
}

func resetDiceChance(setDiceChance []uint, diceChance *[]uint) {
	*diceChance = make([]uint, len(setDiceChance))
	copy(*diceChance, setDiceChance)
}
