package main

import (
	"dice-game/entities"
	"dice-game/tile"

	"github.com/hajimehoshi/ebiten/v2"
)

var SnakeLadderPhaseChance []uint = []uint{1, 0, 0, 0}

type SnakeLadder struct {
	Player    *entities.Player
	Tilemap   [10][10]tile.Tile
	Chance    []uint
	CurseTile uint
}

func (g *Game) SnakeLadderPhaseUpdate() {
	// 주사위 선택 및 클릭 처리
	if g.DiceManager.Select != -1 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if g.SnackLadder.Chance[g.DiceManager.Select] > 0 {
				g.DiceManager.Clicked()
			} else {
				g.DiceManager.Select = -1
			}
			return
		}

		// 주사위 결과 처리
		diceVal := g.DiceManager.MyDices.SnackLadderDice[0].Val + 1
		movePlayer(g.SnackLadder.Player, diceVal)
		g.DiceManager.NotClicked()
		g.SnackLadder.Chance[0]--
	}

	// 타일 클릭 처리
	if g.SnackLadder.Chance[0] == 0 {
		handleTileClick(g)
	}
}

func movePlayer(player *entities.Player, steps int) {
	for i := 0; i < steps; i++ {
		player.NowIndex++
		if (player.NowIndex+1)%10 == 0 {
			break // 10의 배수에 도달하면 멈춤
		}
	}
}

func handleTileClick(g *Game) {
	x, y := ebiten.CursorPosition()
	nowTile := tile.NowTile(g.SnackLadder.Tilemap, g.SnackLadder.Player)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if float64(x) >= nowTile.X && float64(x) < nowTile.X+64 &&
			float64(y) >= nowTile.Y && float64(y) < nowTile.Y+64 {
			// 저주 타일일 경우
			if nowTile.Type == 6 {
				g.Cursed++
				if g.Cursed >= 10 {
					g.GameOver()
				}
			}
			g.initBattlePhase()
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
