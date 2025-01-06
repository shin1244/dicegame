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
			g.Player.NowIndex += g.Dice.Val + 1
			log.Println(g.Player.PassCheck[(g.Player.NowIndex+1)/10], (g.Player.NowIndex+1)/10)
			// if g.Player.PassCheck[(g.Player.NowIndex+1)/10+1] {

			// }
			g.Dice.IsPressed = false
		}
		// 주사위 떨림 초기화
		g.Dice.X = 0
		g.Dice.Y = 0

	}
	if g.Player.NowIndex >= 100 {
		g.Player.NowIndex = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	tileMap := g.tilemap
	dice := g.Dice
	player := g.Player

	// 타일 그리기
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			// 짝수 층에선 반대로 그리기
			var drawX int
			if y%2 == 0 {
				drawX = x
			} else {
				drawX = 9 - x
			}

			tile := tileMap[y][x]

			tileMap[y][x].X = float64(0 + drawX*64)
			tileMap[y][x].Y = float64(0 + (9-y)*64)

			tileOpt := &ebiten.DrawImageOptions{}
			tileOpt.GeoM.Translate(tileMap[y][x].X, tileMap[y][x].Y)

			screen.DrawImage(tile.Image, tileOpt)
		}
	}

	// 주사위 그리기
	diceOpt := &ebiten.DrawImageOptions{}
	diceOpt.GeoM.Translate(float64(dice.X), float64(dice.Y)) // 현재 위치
	screen.DrawImage(
		dice.Img.SubImage(
			dice.DiceSpriteSheet.Rect(dice.Val),
		).(*ebiten.Image),
		diceOpt,
	)

	PlayerOpt := &ebiten.DrawImageOptions{}
	PlayerIdx := player.NowIndex
	PlayerOpt.GeoM.Translate(float64(tileMap[PlayerIdx/10][PlayerIdx%10].X), float64(tileMap[PlayerIdx/10][PlayerIdx%10].Y)) // 현재 위치
	screen.DrawImage(
		g.Player.Img,
		PlayerOpt,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Dice Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal("게임 실행 불가")
	}
}
