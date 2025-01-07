package main

import (
	"dice-game/entities"
	"dice-game/phase"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SnackLadderPhase *phase.SnakeLadder
	Phase            uint
	Dices            *entities.MyDices
	SelectedDice     []entities.Dice
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

func (g *Game) RollDice() {
	x, y := ebiten.CursorPosition()

	// 주사위 섹션의 경계 계산을 위한 상수
	startX := 640.0
	sectionHeight := 160.0

	// 커서가 주사위 영역 내에 있는지 확인 (x축)
	if float64(x) >= startX {
		// y축 위치에 따라 섹션 확인
		switch {
		case float64(y) < sectionHeight:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.SelectedDice == nil {
				fmt.Println("스네이크 래더 주사위 섹션 클릭")
				g.SelectedDice = g.Dices.SnackLadderDice
			}
		case float64(y) < sectionHeight*2:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.SelectedDice == nil {
				fmt.Println("아군 주사위 섹션 클릭")
				g.SelectedDice = g.Dices.FriendlyDice
			}
		case float64(y) < sectionHeight*3:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.SelectedDice == nil {
				fmt.Println("적군 주사위 섹션 클릭")
				g.SelectedDice = g.Dices.EnemyDice
			}
		default:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.SelectedDice == nil {
				fmt.Println("리워드 주사위 섹션 클릭")
				g.SelectedDice = g.Dices.RewardDice
			}
		}
	}
	if g.SelectedDice != nil && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		fmt.Println("주사위 클릭 종료")
		g.SelectedDice = nil
	}
}
