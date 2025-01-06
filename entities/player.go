package entities

type Player struct {
	*Sprite
	NowIndex  int
	PassCheck [10]bool
}
