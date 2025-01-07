package main

import (
	"dice-game/entities"
	"dice-game/phase"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SnackLadderPhase *phase.SnakeLadder
	Phase            uint
	Dices            *entities.MyDices
}

func NewGame() *Game {
	game := &Game{
		SnackLadderPhase: phase.NewSnLPhase(),
		Dices:            entities.NewDices(),
		Phase:            0,
	}
	game.Dices.SnackLadderDice = append(game.Dices.SnackLadderDice, *entities.DefaultDice())
	game.Dices.FriendlyDice = append(game.Dices.FriendlyDice, *entities.DefaultDice())
	game.Dices.FriendlyDice = append(game.Dices.FriendlyDice, *entities.DefaultDice())
	game.Dices.EnemyDice = append(game.Dices.EnemyDice, *entities.DefaultDice())
	game.Dices.EnemyDice = append(game.Dices.EnemyDice, *entities.DefaultDice())
	game.Dices.RewardDice = append(game.Dices.RewardDice, *entities.DefaultDice())
	return game
}

func (g *Game) drawDices(screen *ebiten.Image) {
	diceSize := 64.0
	startX := 640.0

	// 섹션별 시작 위치 계산을 위한 상수
	sectionHeight := 160.0

	// 스네이크 주사위 섹션
	for i, dice := range g.Dices.SnackLadderDice {
		diceOpt := &ebiten.DrawImageOptions{}
		diceOpt.GeoM.Translate(startX+dice.X+float64(i)*diceSize, float64(i)*diceSize)
		screen.DrawImage(
			dice.Img.SubImage(
				dice.DiceSpriteSheet.Rect(dice.Val),
			).(*ebiten.Image),
			diceOpt,
		)
	}

	// 아군 주사위 섹션
	for i, dice := range g.Dices.FriendlyDice {
		diceOpt := &ebiten.DrawImageOptions{}
		diceOpt.GeoM.Translate(startX+dice.X+float64(i)*diceSize, sectionHeight)
		screen.DrawImage(
			dice.Img.SubImage(
				dice.DiceSpriteSheet.Rect(dice.Val),
			).(*ebiten.Image),
			diceOpt,
		)
	}

	// 적군 주사위 섹션
	for i, dice := range g.Dices.EnemyDice {
		diceOpt := &ebiten.DrawImageOptions{}
		diceOpt.GeoM.Translate(startX+dice.X+float64(i)*diceSize, sectionHeight*2)
		screen.DrawImage(
			dice.Img.SubImage(
				dice.DiceSpriteSheet.Rect(dice.Val),
			).(*ebiten.Image),
			diceOpt,
		)
	}

	// 리워드 주사위 섹션
	for i, dice := range g.Dices.RewardDice {
		diceOpt := &ebiten.DrawImageOptions{}
		diceOpt.GeoM.Translate(startX+dice.X+float64(i)*diceSize, sectionHeight*3+float64(i)*diceSize)
		screen.DrawImage(
			dice.Img.SubImage(
				dice.DiceSpriteSheet.Rect(dice.Val),
			).(*ebiten.Image),
			diceOpt,
		)
	}
}

func RollDice()
