package spritesheet

import "image"

type Spritesheet struct {
	SpriteWidth  int
	SpriteHeight int
	SpriteSize   int
}

func (s *Spritesheet) Rect(index int) image.Rectangle {
	x := (index % s.SpriteWidth) * s.SpriteSize
	y := (index / s.SpriteWidth) * s.SpriteSize

	return image.Rect(x, y, x+s.SpriteSize, y+s.SpriteSize)
}

func NewSpriteSheet(w, h, s int) *Spritesheet {
	return &Spritesheet{w, h, s}
}
