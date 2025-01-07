package game

import (
	"dice-game/entities"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DiceStartX    = 640.0 // 게임판 오른쪽 시작 위치
	DiceStartY    = 160.0 // 게임판 오른쪽 시작 위치
	DiceSize      = 64.0  // 주사위 크기
	SectionHeight = 160.0 // 섹션 높이
)

type DiceManager struct {
	MyDices       entities.MyDices
	SelectedDices int
}

func NewDiceManager() *DiceManager {
	dm := &DiceManager{
		MyDices: entities.MyDices{
			SnackLadderDice:       make([]entities.Dice, 0),
			SnackLadderDiceChange: 0,
			FriendlyDice:          make([]entities.Dice, 0),
			FriendlyDiceChange:    0,
			EnemyDice:             make([]entities.Dice, 0),
			EnemyDiceChange:       0,
			RewardDice:            make([]entities.Dice, 0),
			RewardDiceChange:      0,
		},
		SelectedDices: -1,
	}

	// 초기 주사위 설정
	dm.MyDices.SnackLadderDice = append(dm.MyDices.SnackLadderDice, *entities.DefaultDice())
	dm.MyDices.FriendlyDice = append(dm.MyDices.FriendlyDice, *entities.DefaultDice(), *entities.DefaultDice())
	dm.MyDices.EnemyDice = append(dm.MyDices.EnemyDice, *entities.DefaultDice(), *entities.DefaultDice())
	dm.MyDices.RewardDice = append(dm.MyDices.RewardDice, *entities.DefaultDice())

	return dm
}

func (dm *DiceManager) Draw(screen *ebiten.Image) {
	// 스네이크 주사위 섹션
	for i, dice := range dm.MyDices.SnackLadderDice {
		dm.drawDiceInSection(screen, &dice, 0, i)
	}

	// 아군 주사위 섹션
	for i, dice := range dm.MyDices.FriendlyDice {
		dm.drawDiceInSection(screen, &dice, 1, i)
	}

	// 적군 주사위 섹션
	for i, dice := range dm.MyDices.EnemyDice {
		dm.drawDiceInSection(screen, &dice, 2, i)
	}

	// 리워드 주사위 섹션
	for i, dice := range dm.MyDices.RewardDice {
		dm.drawDiceInSection(screen, &dice, 3, i)
	}
}

// 주사위 그리는 함수
func (dm *DiceManager) drawDiceInSection(screen *ebiten.Image, dice *entities.Dice, section, index int) {
	if dice.Img == nil {
		return
	}

	opt := &ebiten.DrawImageOptions{}
	x := DiceStartX + float64(index)*(DiceSize)
	y := float64(section) * SectionHeight
	opt.GeoM.Translate(x+dice.X, y+dice.Y)

	screen.DrawImage(
		dice.Img.SubImage(
			dice.DiceSpriteSheet.Rect(dice.Val),
		).(*ebiten.Image),
		opt,
	)
}

// 주사위 클릭 인식 함수
func (dm *DiceManager) HandleInput() {
	x, y := ebiten.CursorPosition()

	if float64(x) >= DiceStartX {
		switch {
		case float64(y) < SectionHeight:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && dm.SelectedDices == -1 {
				dm.SelectedDices = 0
			}
		case float64(y) < SectionHeight*2:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && dm.SelectedDices == -1 {
				dm.SelectedDices = 1
			}
		case float64(y) < SectionHeight*3:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && dm.SelectedDices == -1 {
				dm.SelectedDices = 2
			}
		default:
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && dm.SelectedDices == -1 {
				dm.SelectedDices = 3
			}
		}
	}
}

// 클릭된 주사위 돌리기
func (dm *DiceManager) Clicked() {
	if dm.SelectedDices == -1 {
		return
	}
	mouseX, mouseY := ebiten.CursorPosition()

	// 선택된 주사위 정보 가져오기
	var dices *[]entities.Dice
	switch dm.SelectedDices {
	case 0:
		dices = &dm.MyDices.SnackLadderDice
	case 1:
		dices = &dm.MyDices.FriendlyDice
	case 2:
		dices = &dm.MyDices.EnemyDice
	case 3:
		dices = &dm.MyDices.RewardDice
	}

	// 선택된 주사위 돌리기
	for i := range *dices {
		(*dices)[i].Val = rand.Intn(6)
		(*dices)[i].X = float64(mouseX) - DiceStartX
		(*dices)[i].Y = float64(mouseY) - DiceStartY*float64(dm.SelectedDices)
	}
}

func (dm *DiceManager) NotClicked() {
	// 모든 주사위 위치 초기화
	allDices := [][]entities.Dice{
		dm.MyDices.SnackLadderDice,
		dm.MyDices.FriendlyDice,
		dm.MyDices.EnemyDice,
		dm.MyDices.RewardDice,
	}

	for _, diceSet := range allDices {
		for i := range diceSet {
			diceSet[i].X = 0
			diceSet[i].Y = 0
		}
	}
	// 선택된 주사위 초기화
	dm.SelectedDices = -1
}
