package main

type Battle struct {
	BattleType uint
}

var BattlePhaseChance []int = []int{0, 1, 1, 0}

// 적 주사위가 남아있으면 안됨
