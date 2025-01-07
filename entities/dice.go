package entities

import (
	"dice-game/spritesheet"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Dice struct {
	*Sprite
	Sides           []int
	Val             int
	DiceSpriteSheet *spritesheet.Spritesheet
}

var diceImg *ebiten.Image

func init() {
	var err error
	diceImg, _, err = ebitenutil.NewImageFromFile("assets/images/Dice64.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
}

type MyDices struct {
	SnackLadderDice []Dice
	FriendlyDice    []Dice
	EnemyDice       []Dice
	RewardDice      []Dice
}

func DiceDefault() *Dice {
	return &Dice{
		Sprite:          &Sprite{Img: diceImg},
		Sides:           []int{1, 2, 3, 4, 5, 6},
		DiceSpriteSheet: spritesheet.NewSpriteSheet(6, 1, 64),
	}
}

func Dice456() *Dice {
	return &Dice{
		Sprite:          &Sprite{Img: diceImg},
		Sides:           []int{4, 5, 6, 4, 5, 6},
		DiceSpriteSheet: spritesheet.NewSpriteSheet(6, 1, 64),
	}
}
