package main

import (
	"dice-game/entities"
	"dice-game/spritesheet"
	"dice-game/tile"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player  *entities.Player
	Dice    *entities.Dice
	tilemap [10][10]tile.Tile
}

func NewGame() *Game {
	diceImg, _, err := ebitenutil.NewImageFromFile("assets/images/Dice.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
	PlayerImg, _, err := ebitenutil.NewImageFromFile("assets/images/Player.png")
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
		Player: &entities.Player{
			Sprite: &entities.Sprite{
				Img: PlayerImg,
				X:   0,
				Y:   0,
			},
			NowIndex: 0,
		},
		tilemap: tile.NewTileMap(),
	}
	return game
}
