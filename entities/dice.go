package entities

import (
	"dice-game/spritesheet"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Dice struct {
	*Sprite
	Sides           []int
	Val             int
	DiceSpriteSheet *spritesheet.Spritesheet
}

type MyDices struct {
	SnackLadderDice []Dice
	FriendlyDice    []Dice
	EnemyDice       []Dice
	RewardDice      []Dice
}

func DefaultDice() *Dice {
	diceImg, _, err := ebitenutil.NewImageFromFile("assets/images/Dice64.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
	return &Dice{
		Sprite: &Sprite{
			Img: diceImg,
			X:   0,
			Y:   0,
		},
		Sides:           []int{1, 2, 3, 4, 5, 6},
		DiceSpriteSheet: spritesheet.NewSpriteSheet(6, 1, 64),
	}
}

func NewDices() *MyDices {
	return &MyDices{
		SnackLadderDice: make([]Dice, 0),
		FriendlyDice:    make([]Dice, 0),
		EnemyDice:       make([]Dice, 0),
		RewardDice:      make([]Dice, 0),
	}
}
