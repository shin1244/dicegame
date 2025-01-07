package main

import (
	"dice-game/entities"
	"dice-game/game"
	"dice-game/tile"

	"github.com/hajimehoshi/ebiten/v2"
)

var SnakeLadderPhaseChance []uint = []uint{1, 0, 0, 0}

type SnakeLadder struct {
	Player     *entities.Player
	Tilemap    [10][10]tile.Tile
	DiceChance []uint
}

func (g *Game) SnakeLadderPhaseUpdate() {
	resetDiceChance(SnakeLadderPhaseChance, &g.SnackLadderPhase.DiceChance)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.DiceManager.SelectedDices != -1 {
		if g.SnackLadderPhase.DiceChance[g.DiceManager.SelectedDices] > 0 {
			g.DiceManager.Clicked()
		} else {
			g.DiceManager.SelectedDices = -1
		}
	} else if g.DiceManager.SelectedDices != -1 {
		diceVal := g.DiceManager.MyDices.SnackLadderDice[0].Val + 1
		for i := 0; i < diceVal; i++ {
			g.SnackLadderPhase.Player.NowIndex++
			if (g.SnackLadderPhase.Player.NowIndex+1)%10 == 0 {
				break // 10의 배수에 도달하면 멈춤
			}
		}
		g.DiceManager.NotClicked()
		g.SnackLadderPhase.DiceChance[0] -= 1
	}
}

func (g *Game) SnakeLadderPhaseDraw(screen *ebiten.Image) {
	tileMap := g.SnackLadderPhase.Tilemap
	player := g.SnackLadderPhase.Player

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
		player.Img,
		PlayerOpt,
	)
}

func NewSnakeLadderPhase(diceManager *game.DiceManager) *SnakeLadder {
	return &SnakeLadder{
		Player:  entities.NewPlayer(),
		Tilemap: tile.NewTileMap(),
	}
}
