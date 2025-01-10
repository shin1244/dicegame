package main

import (
	"dice-game/entities"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	var err error
	RewardBoss, _, err = ebitenutil.NewImageFromFile("assets/images/BossReward.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
	RewardNormal, _, err = ebitenutil.NewImageFromFile("assets/images/reward.png")
	if err != nil {
		log.Fatal("이미지 불러오기 오류")
	}
}

type Reward struct {
	Chance []uint
	IsBoss bool
}

var RewardBoss *ebiten.Image
var RewardNormal *ebiten.Image
var RewardPhaseChance []uint = []uint{0, 0, 0, 1}

func (g *Game) RewardPhaseUpdate() {
	if g.Reward.IsBoss {
		if g.DiceManager.Select != -1 {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if g.DiceManager.Select == 1 {
					AddDice(&g.DiceManager.MyDices.FriendlyDice, *entities.DiceDefault())
					g.DiceManager.Select = -1
					g.Reward.IsBoss = false
				} else if g.DiceManager.Select == 2 {
					AddDice(&g.DiceManager.MyDices.EnemyDice, *entities.DiceDefault())
					g.DiceManager.Select = -1
					g.Reward.IsBoss = false
				} else {
					g.DiceManager.Select = -1
				}
			}
		}
	} else {
		if g.DiceManager.Select != -1 {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				if g.Reward.Chance[g.DiceManager.Select] > 0 {
					g.DiceManager.Clicked()
				} else {
					g.DiceManager.Select = -1
				}
				return
			}

			for _, dice := range g.DiceManager.MyDices.RewardDice {
				g.Gold += dice.Sides[dice.Val]
			}
			g.DiceManager.NotClicked()
			g.ChangeCursedTile()
			g.initSnakeLadderPhase()
		}
	}
}

func (g *Game) RewardPhaseDraw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, 64) // 상태창 높이만큼 아래로 이동

	if g.Reward.IsBoss {
		screen.DrawImage(RewardBoss, opt)
	} else {
		screen.DrawImage(RewardNormal, opt)
	}
}

func NewRewardPhase() *Reward {
	return &Reward{}
}
