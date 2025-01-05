package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) { // 스페이스바로 주사위 굴리기
		g.Dice.Val = rand.Intn(6)
		g.Dice.IsPressed = true

		// 주사위 떨림 효과
		g.Dice.X = float64(rand.Intn(9) - 4)
		g.Dice.Y = float64(rand.Intn(9) - 4)
	} else {
		if g.Dice.IsPressed {
			g.Dice.IsPressed = false
		}
		// 주사위 떨림 초기화
		g.Dice.X = 0
		g.Dice.Y = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	options := &ebiten.DrawImageOptions{}                        // 주사위 이미지 옵션 생성
	options.GeoM.Translate(float64(g.Dice.X), float64(g.Dice.Y)) // 현재 위치
	screen.DrawImage(
		g.Dice.Img.SubImage(
			g.Dice.DiceSpriteSheet.Rect(g.Dice.Val),
		).(*ebiten.Image),
		options,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Dice Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal("게임 실행 불가")
	}
}
