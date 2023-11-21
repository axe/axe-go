package ua

import "github.com/axe/axe-go/pkg/ui"

var (
	// Explode
	ExplodeGen = func(scale float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0, Transparency: 0, Scale: &ui.Coord{X: 1, Y: 1}, Origin: OriginCenter},
			{Time: 1, Transparency: 1, Scale: &ui.Coord{X: scale, Y: scale}, Origin: OriginCenter},
		})
	}
	Explode = ExplodeGen(4)

	// Reveal
	RevealGen = func(startX, startY float32, origin ui.AmountPoint) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0, Scale: &ui.Coord{X: startX, Y: startY}, Origin: origin},
			{Time: 1, Scale: &ui.Coord{X: 1, Y: 1}, Origin: origin},
		})
	}
	RevealDown      = RevealGen(1, 0, OriginTop)
	RevealUp        = RevealGen(1, 0, OriginBottom)
	RevealUpDown    = RevealGen(1, 0, OriginCenter)
	RevealLeft      = RevealGen(0, 1, OriginCenterLeft)
	RevealRight     = RevealGen(0, 1, OriginCenterRight)
	RevealLeftRight = RevealGen(0, 1, OriginCenter)
)
