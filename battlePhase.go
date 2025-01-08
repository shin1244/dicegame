package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	var err error
	BattleType0, _, err = ebitenutil.NewImageFromFile("assets/images/B0.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
	BattleType1, _, err = ebitenutil.NewImageFromFile("assets/images/B1.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
}

type Battle struct {
	BattleType  uint
	Chance      []uint
	FriendScore int
	EnemyScore  int
}

var BattleType0 *ebiten.Image
var BattleType1 *ebiten.Image
var BattlePhaseChance []uint = []uint{0, 1, 1, 0}

// 아군 주사위를 던지기 전 적 주사위가 남아있으면 안됨
func (g *Game) BattlePhaseUpdate() {
	if g.DiceManager.Select != -1 {
		handleDiceSelection(g)
		return
	}

	if isBattleComplete(g.Battle.Chance) {
		handleBattleResult(g)
	}
}

func handleDiceSelection(g *Game) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		handleMouseClick(g)
		return
	}

	updateBattleScore(g)
	g.DiceManager.NotClicked()
}

func handleMouseClick(g *Game) {
	if g.Battle.Chance[g.DiceManager.Select] > 0 {
		if g.DiceManager.Select == 1 && g.Battle.Chance[2] != 0 {
			g.DiceManager.Select = -1
		} else {
			g.DiceManager.Clicked()
		}
	} else {
		g.DiceManager.Select = -1
	}
}

func updateBattleScore(g *Game) {
	g.Battle.Chance[g.DiceManager.Select] -= 1
	if g.DiceManager.Select == 1 {
		for _, dice := range g.DiceManager.MyDices.FriendlyDice {
			g.Battle.FriendScore += dice.Sides[dice.Val]
		}
	} else {
		for _, dice := range g.DiceManager.MyDices.EnemyDice {
			g.Battle.EnemyScore += dice.Sides[dice.Val]
		}
	}
}

func isBattleComplete(chance []uint) bool {
	return chance[1] == 0 && chance[2] == 0
}

func handleBattleResult(g *Game) {
	switch g.Battle.BattleType {
	case 0:
		g.initRewardPhase(g.Battle.FriendScore > g.Battle.EnemyScore)
	case 1:
		g.initRewardPhase(g.Battle.FriendScore < g.Battle.EnemyScore)
	}
}

func (g *Game) BattlePhaseDraw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, 64) // 상태창 높이만큼 아래로 이동

	if g.Battle.BattleType == 0 {
		screen.DrawImage(BattleType0, opt)
	} else {
		screen.DrawImage(BattleType1, opt)
	}
}

func NewBattlePhase() *Battle {
	return &Battle{
		BattleType: 0,
		Chance:     BattlePhaseChance,
	}
}
