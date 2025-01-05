package entities

import "dice-game/spritesheet"

type Dice struct {
	*Sprite
	Val             int
	IsPressed       bool
	DiceSpriteSheet *spritesheet.Spritesheet
}
