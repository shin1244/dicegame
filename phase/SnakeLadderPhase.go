package phase

import (
	"dice-game/entities"
	"dice-game/tile"

	"github.com/hajimehoshi/ebiten/v2"
)

type SnakeLadder struct {
	Player  *entities.Player
	Tilemap [10][10]tile.Tile
}

func (s *SnakeLadder) SnLPhaseUpdate() {
	// 스페이스바로 주사위 굴리기
	// if ebiten.IsKeyPressed(ebiten.KeySpace) {
	// 	s.Dices. = rand.Intn(6) // 수정해야함
	// 	s.Dice.IsPressed = true
	// 	// 주사위 떨림 효과
	// 	s.Dice.X = float64(rand.Intn(9) - 4)
	// 	s.Dice.Y = float64(rand.Intn(9) - 4)
	// } else {
	// 	if s.Dice.IsPressed {
	// 		diceVal := s.Dice.Val + 1
	// 		for i := 0; i < diceVal; i++ {
	// 			s.Player.NowIndex++
	// 			if (s.Player.NowIndex+1)%10 == 0 {
	// 				break // 10의 배수에 도달하면 멈춤
	// 			}
	// 		}
	// 		s.Dice.IsPressed = false
	// 	}
	// 	// 주사위 떨림 초기화
	// 	s.Dice.X = 0
	// 	s.Dice.Y = 0
	// }
}

func (s *SnakeLadder) SnLPhaseDraw(screen *ebiten.Image) {
	tileMap := s.Tilemap
	player := s.Player

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

	PlayerOpt := &ebiten.DrawImageOptions{}
	PlayerIdx := player.NowIndex
	PlayerOpt.GeoM.Translate(float64(tileMap[PlayerIdx/10][PlayerIdx%10].X), float64(tileMap[PlayerIdx/10][PlayerIdx%10].Y)) // 현재 위치
	screen.DrawImage(
		s.Player.Img,
		PlayerOpt,
	)
}

func NewSnLPhase() *SnakeLadder {
	SnL := &SnakeLadder{
		Player:  entities.NewPlayer(),
		Tilemap: tile.NewTileMap(),
	}
	return SnL
}
