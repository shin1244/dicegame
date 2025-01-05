package main

import (
	"dice-game/entities"
	"dice-game/spritesheet"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Dice *entities.Dice
}

func NewGame() *Game {
	diceImg, _, err := ebitenutil.NewImageFromFile("assets/image/Dice.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}

	game := &Game{
		Dice: &entities.Dice{
			Sprite: &entities.Sprite{
				Img: diceImg,
				X:   0,
				Y:   0,
			},
			Val:             0,
			IsPressed:       false,
			DiceSpriteSheet: spritesheet.NewSpriteSheet(6, 1, 96),
		},
	}

	return game
}
