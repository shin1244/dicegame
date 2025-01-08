package main

import (
	"dice-game/entities"
	"dice-game/tile"

	"github.com/hajimehoshi/ebiten/v2"
)

var SnakeLadderPhaseChance []uint = []uint{1, 0, 0, 0}

type SnakeLadder struct {
	Player  *entities.Player
	Tilemap [10][10]tile.Tile
	Chance  []uint
}

func (g *Game) SnakeLadderPhaseUpdate() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.DiceManager.Select != -1 {
		// 현재 선택한 주사위 기회가 0 이상이면
		if g.SnackLadder.Chance[g.DiceManager.Select] > 0 {
			g.DiceManager.Clicked()
		} else {
			g.DiceManager.Select = -1
		}
	} else if g.DiceManager.Select != -1 {
		diceVal := g.DiceManager.MyDices.SnackLadderDice[0].Val + 1
		for i := 0; i < diceVal; i++ {
			g.SnackLadder.Player.NowIndex++
			if (g.SnackLadder.Player.NowIndex+1)%10 == 0 {
				break // 10의 배수에 도달하면 멈춤
			}
		}
		g.DiceManager.NotClicked()
		g.SnackLadder.Chance[0] -= 1
	}
	if g.SnackLadder.Chance[0] == 0 {
		playerIdx := g.SnackLadder.Player.NowIndex
		x, y := ebiten.CursorPosition()
		tileX := g.SnackLadder.Tilemap[playerIdx/10][playerIdx%10].X
		tileY := g.SnackLadder.Tilemap[playerIdx/10][playerIdx%10].Y

		// 마우스 클릭 이벤트와 타일 영역 체크
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if float64(x) >= tileX && float64(x) < tileX+64 &&
				float64(y) >= tileY && float64(y) < tileY+64 {
				g.initBattlePhase()
			}
		}
	}
}

func (g *Game) SnakeLadderPhaseDraw(screen *ebiten.Image) {
	tileMap := g.SnackLadder.Tilemap
	player := g.SnackLadder.Player
	// 타일 그리기
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			tile := tileMap[y][x]

			tileOpt := &ebiten.DrawImageOptions{}
			tileOpt.GeoM.Translate(tile.X, tile.Y)

			screen.DrawImage(tile.Image, tileOpt)
		}
	}
	PlayerOpt := &ebiten.DrawImageOptions{}
	PlayerIdx := player.NowIndex
	PlayerOpt.GeoM.Translate(tileMap[PlayerIdx/10][PlayerIdx%10].X, tileMap[PlayerIdx/10][PlayerIdx%10].Y)
	screen.DrawImage(
		player.Img,
		PlayerOpt,
	)
}

func NewSnakeLadderPhase() *SnakeLadder {
	sl := &SnakeLadder{
		Player:  entities.NewPlayer(),
		Tilemap: tile.NewTileMap(),
	}
	sl.InitializeTilePositions()
	return sl
}

func (s *SnakeLadder) InitializeTilePositions() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			var drawX int
			if y%2 == 0 {
				drawX = x
			} else {
				drawX = 9 - x
			}

			s.Tilemap[y][x].X = float64(0 + drawX*64)
			s.Tilemap[y][x].Y = float64(64 + (9-y)*64)
		}
	}
}
