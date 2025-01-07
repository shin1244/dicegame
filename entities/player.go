package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	*Sprite
	NowIndex  int
	PassCheck [10]bool
}

func NewPlayer() *Player {
	PlayerImg, _, err := ebitenutil.NewImageFromFile("assets/images/Player.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
	return &Player{
		Sprite: &Sprite{
			Img: PlayerImg,
			X:   0,
			Y:   0,
		},
		NowIndex: 0,
	}
}
