package animations

type Animations struct {
	First        int
	Last         int
	Step         int
	SpeedInTPS   float32
	frameCounter float32
	frame        int
}

func (a *Animations) Update() {
	a.frameCounter -= 1.0
	if a.frameCounter < 0.0 {
		a.frameCounter = a.SpeedInTPS
		a.frame += a.Step
		if a.frame > a.Last {
			a.frame = a.First
		}
	}
}

func (a *Animations) Frame() int {
	return a.frame
}

func NewAnimation(first, last, step int, speed float32) *Animations {
	return &Animations{
		first,
		last,
		step,
		speed,
		speed,
		first,
	}
}
