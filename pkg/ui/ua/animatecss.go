package ua

import (
	"github.com/axe/axe-go/pkg/ease"
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/ui"
)

var (
	// Bounce
	bounceEasing0 = ease.NewBezier(0.215, 0.610, 0.355, 1.000)
	bounceEasing1 = ease.NewBezier(0.755, 0.050, 0.855, 0.060)
	bounceEasing2 = ease.NewBezier(0.755, 0.050, 0.855, 0.060)

	BounceGen = func(x, y float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Easing: bounceEasing0},
			{Time: 0.20, Easing: bounceEasing0},
			{Time: 0.40, Easing: bounceEasing1, Translate: ui.NewAmountPoint(x, y)},
			{Time: 0.43, Easing: bounceEasing1, Translate: ui.NewAmountPoint(x, y)},
			{Time: 0.53, Easing: bounceEasing0},
			{Time: 0.70, Easing: bounceEasing2, Translate: ui.NewAmountPoint(x*0.5, y*0.5)},
			{Time: 0.80, Easing: bounceEasing0},
			{Time: 0.90, Easing: bounceEasing0, Translate: ui.NewAmountPoint(x*0.13, y*0.13)},
			{Time: 1.00, Easing: bounceEasing0},
		})
	}
	Bounce = BounceGen(0, -30)

	// RubberBand
	RubberBandGen = func(s float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Scale: &gfx.Coord{X: 1.00, Y: 1.00}, Origin: OriginCenter},
			{Time: 0.30, Scale: &gfx.Coord{X: 1 + s, Y: 1 - s}, Origin: OriginCenter},
			{Time: 0.40, Scale: &gfx.Coord{X: 1 - s, Y: 1 + s}, Origin: OriginCenter},
			{Time: 0.50, Scale: &gfx.Coord{X: 1 + (s * 0.6), Y: 1 - (s * 0.6)}, Origin: OriginCenter},
			{Time: 0.65, Scale: &gfx.Coord{X: 1 - (s * 0.2), Y: 1 + (s * 0.2)}, Origin: OriginCenter},
			{Time: 0.75, Scale: &gfx.Coord{X: 1 + (s * 0.2), Y: 1 - (s * 0.2)}, Origin: OriginCenter},
			{Time: 1.00, Scale: &gfx.Coord{X: 1.00, Y: 1.00}, Origin: OriginCenter},
		})
	}
	RubberBand = RubberBandGen(0.25)

	// Flash
	Flash = AnimationGen([]ui.BasicAnimationFrame{
		{Time: 0.00},
		{Time: 0.25, Transparency: 1},
		{Time: 0.50},
		{Time: 0.75, Transparency: 1},
		{Time: 1.00},
	})

	// Pulse
	PulseGen = func(x, y float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Scale: &gfx.Coord{X: 1.00, Y: 1.00}, Origin: OriginCenter},
			{Time: 0.25, Scale: &gfx.Coord{X: x, Y: y}, Origin: OriginCenter},
			{Time: 1.00, Scale: &gfx.Coord{X: 1.00, Y: 1.00}, Origin: OriginCenter},
		})
	}
	Pulse = PulseGen(1.05, 1.05)

	// Shake
	ShakeGen = func(dx, dy float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00},
			{Time: 0.10, Translate: ui.NewAmountPoint(-dx, -dy)},
			{Time: 0.20, Translate: ui.NewAmountPoint(dx, dy)},
			{Time: 0.30, Translate: ui.NewAmountPoint(-dx, -dy)},
			{Time: 0.40, Translate: ui.NewAmountPoint(dx, dy)},
			{Time: 0.50, Translate: ui.NewAmountPoint(-dx, -dy)},
			{Time: 0.60, Translate: ui.NewAmountPoint(dx, dy)},
			{Time: 0.70, Translate: ui.NewAmountPoint(-dx, -dy)},
			{Time: 0.80, Translate: ui.NewAmountPoint(dx, dy)},
			{Time: 0.90, Translate: ui.NewAmountPoint(-dx, -dy)},
			{Time: 1.00},
		})
	}
	Shake = ShakeGen(10, 0)

	// Swing
	SwingGen = func(originX, originY, degrees float32) ui.BasicAnimation {
		origin := ui.NewAmountPointParent(originX, originY)
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Origin: origin},
			{Time: 0.20, Origin: origin, Rotate: degrees},
			{Time: 0.40, Origin: origin, Rotate: -degrees * 0.6},
			{Time: 0.60, Origin: origin, Rotate: degrees * 0.3},
			{Time: 0.80, Origin: origin, Rotate: -degrees * 0.3},
			{Time: 1.00, Origin: origin},
		})
	}
	Swing = SwingGen(0.5, 0, 15)

	// Tada
	TadaGen = func(scaleUp, rotate float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Scale: &gfx.Coord{X: 1, Y: 1}, Rotate: 0, Origin: OriginCenter},
			{Time: 0.10, Scale: &gfx.Coord{X: 1 - scaleUp, Y: 1 - scaleUp}, Rotate: -rotate, Origin: OriginCenter},
			{Time: 0.20, Scale: &gfx.Coord{X: 1 - scaleUp, Y: 1 - scaleUp}, Rotate: -rotate, Origin: OriginCenter},
			{Time: 0.30, Scale: &gfx.Coord{X: 1 + scaleUp, Y: 1 + scaleUp}, Rotate: rotate, Origin: OriginCenter},
			{Time: 0.40, Scale: &gfx.Coord{X: 1 + scaleUp, Y: 1 + scaleUp}, Rotate: -rotate, Origin: OriginCenter},
			{Time: 0.50, Scale: &gfx.Coord{X: 1 + scaleUp, Y: 1 + scaleUp}, Rotate: rotate, Origin: OriginCenter},
			{Time: 0.60, Scale: &gfx.Coord{X: 1 + scaleUp, Y: 1 + scaleUp}, Rotate: -rotate, Origin: OriginCenter},
			{Time: 0.70, Scale: &gfx.Coord{X: 1 + scaleUp, Y: 1 + scaleUp}, Rotate: rotate, Origin: OriginCenter},
			{Time: 0.80, Scale: &gfx.Coord{X: 1 + scaleUp, Y: 1 + scaleUp}, Rotate: -rotate, Origin: OriginCenter},
			{Time: 0.90, Scale: &gfx.Coord{X: 1 + scaleUp, Y: 1 + scaleUp}, Rotate: rotate, Origin: OriginCenter},
			{Time: 1.00, Scale: &gfx.Coord{X: 1, Y: 1}, Rotate: 0, Origin: OriginCenter},
		})
	}
	Tada = TadaGen(0.1, 3)

	// Wobble
	WobbleGen = func(rotate, dx, dy float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Rotate: 0},
			{Time: 0.15, Rotate: -rotate, Translate: ui.NewAmountPointParent(dx, dy)},
			{Time: 0.30, Rotate: rotate * 0.6, Translate: ui.NewAmountPointParent(-dx*0.8, -dy*0.8)},
			{Time: 0.45, Rotate: -rotate * 0.6, Translate: ui.NewAmountPointParent(dx*0.6, dy*0.6)},
			{Time: 0.60, Rotate: rotate * 0.4, Translate: ui.NewAmountPointParent(-dx*0.4, -dy*0.4)},
			{Time: 0.75, Rotate: -rotate * 0.2, Translate: ui.NewAmountPointParent(dx*0.2, dy*0.2)},
			{Time: 1.00, Rotate: 0},
		})
	}
	Wobble = WobbleGen(.5, -2.5, 0)

	// Fade
	FadeGen = func(start, startTime, end, endTime float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: startTime, Transparency: start},
			{Time: endTime, Transparency: end},
		})
	}
	FadeOut    = FadeGen(0, 0, 1, 1)
	FadeIn     = FadeGen(1, 0, 0, 1)
	FadeOutEnd = AnimationGen([]ui.BasicAnimationFrame{
		{Time: 0, Transparency: 0},
		{Time: 1, Transparency: 0},
		{Time: 1, Transparency: 1},
	})

	// BounceIn
	BounceInGen = func(dx, dy float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Transparency: 1, Translate: ui.NewAmountPoint(dx*100, dy*100)},
			{Time: 0.60, Transparency: 0, Translate: ui.NewAmountPoint(-dx*0.83, -dy*0.83)},
			{Time: 0.75, Transparency: 0, Translate: ui.NewAmountPoint(dx*0.33, dy*0.33)},
			{Time: 0.90, Transparency: 0, Translate: ui.NewAmountPoint(-dx*0.16, -dy*0.16)},
			{Time: 1.00, Transparency: 0},
		}).WithDuration(2.0).WithEasing(ease.NewBezier(0.215, 0.610, 0.355, 1.000))
	}
	BounceInDown  = BounceInGen(0, -30)
	BounceInLeft  = BounceInGen(-30, 0)
	BounceInRight = BounceInGen(30, 0)
	BounceInUp    = BounceInGen(0, 30)
	BounceIn      = AnimationGen([]ui.BasicAnimationFrame{
		{Time: 0.00, Transparency: 1.0, Scale: &gfx.Coord{X: 0.3, Y: 0.3}, Origin: OriginCenter},
		{Time: 0.20, Transparency: 0.66, Scale: &gfx.Coord{X: 1.1, Y: 1.1}, Origin: OriginCenter},
		{Time: 0.40, Transparency: 0.33, Scale: &gfx.Coord{X: 0.9, Y: 0.9}, Origin: OriginCenter},
		{Time: 0.60, Transparency: 0, Scale: &gfx.Coord{X: 1.03, Y: 1.03}, Origin: OriginCenter},
		{Time: 0.80, Transparency: 0, Scale: &gfx.Coord{X: 0.97, Y: 0.97}, Origin: OriginCenter},
		{Time: 1.00, Transparency: 0, Scale: &gfx.Coord{X: 1, Y: 1}, Origin: OriginCenter},
	}).WithEasing(ease.NewBezier(0.215, 0.610, 0.355, 1.000))

	// BounceOut
	BounceOutDown  = BounceInUp.Reverse()
	BounceOutLeft  = BounceInRight.Reverse()
	BounceOutRight = BounceInLeft.Reverse()
	BounceOutUp    = BounceInDown.Reverse()
	BounceOut      = BounceIn.Reverse()

	// Translate
	TranslateGen = func(start, end ui.AmountPoint) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0, Translate: start},
			{Time: 1, Translate: end},
		})
	}

	// FadeIn
	FadeInDown     = TranslateGen(ui.NewAmountPointParent(0, -1), ui.AmountPoint{}).Merge(FadeIn)
	FadeInDownBig  = TranslateGen(ui.NewAmountPoint(0, -2000), ui.AmountPoint{}).Merge(FadeIn)
	FadeInLeft     = TranslateGen(ui.NewAmountPointParent(-1, 0), ui.AmountPoint{}).Merge(FadeIn)
	FadeInLeftBig  = TranslateGen(ui.NewAmountPoint(-2000, 0), ui.AmountPoint{}).Merge(FadeIn)
	FadeInRight    = TranslateGen(ui.NewAmountPointParent(1, 0), ui.AmountPoint{}).Merge(FadeIn)
	FadeInRightBig = TranslateGen(ui.NewAmountPoint(2000, 0), ui.AmountPoint{}).Merge(FadeIn)
	FadeInUp       = TranslateGen(ui.NewAmountPointParent(0, 1), ui.AmountPoint{}).Merge(FadeIn)
	FadeInUpBig    = TranslateGen(ui.NewAmountPoint(0, 2000), ui.AmountPoint{}).Merge(FadeIn)

	// FadeOut
	FadeOutDown     = FadeInUp.Reverse()
	FadeOutDownBig  = FadeInUpBig.Reverse()
	FadeOutLeft     = FadeInRight.Reverse()
	FadeOutLeftBig  = FadeInRightBig.Reverse()
	FadeOutRight    = FadeInLeft.Reverse()
	FadeOutRightBig = FadeInLeftBig.Reverse()
	FadeOutUp       = FadeInDown.Reverse()
	FadeOutUpBig    = FadeInDownBig.Reverse()

	// Rotate
	RotateGen = func(start, end, originX, originY float32) ui.BasicAnimation {
		center := ui.NewAmountPointParent(originX, originY)
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0, Rotate: start, Origin: center},
			{Time: 1, Rotate: end, Origin: center},
		})
	}
	RotateIn           = RotateGen(-200, 0, 0.5, 0.5).Merge(FadeIn)
	RotateInDownLeft   = RotateGen(-45, 0, 0, 1).Merge(FadeIn)
	RotateInDownRight  = RotateGen(45, 0, 1, 1).Merge(FadeIn)
	RotateInUpLeft     = RotateGen(45, 0, 0, 0).Merge(FadeIn)
	RotateInUpRight    = RotateGen(-90, 0, 1, 1).Merge(FadeIn)
	RotateOut          = RotateGen(0, 200, 0.5, 0.5).Merge(FadeOut)
	RotateOutDownLeft  = RotateGen(0, 45, 0, 1).Merge(FadeOut)
	RotateOutDownRight = RotateGen(0, -45, 1, 1).Merge(FadeOut)
	RotateOutUpLeft    = RotateGen(0, -45, 0, 1).Merge(FadeOut)
	RotateOutUpRight   = RotateGen(0, -90, 1, 1).Merge(FadeOut)

	// Hinge
	Hinge = AnimationGen([]ui.BasicAnimationFrame{
		{Time: 0.00, Rotate: 0},
		{Time: 0.20, Rotate: 80},
		{Time: 0.40, Rotate: 60},
		{Time: 0.60, Rotate: 80},
		{Time: 0.80, Rotate: 60},
		{Time: 1.00, Rotate: 0, Transparency: 1, Translate: ui.NewAmountPoint(0, 700)},
	}).WithDuration(2.0).WithEasing(ease.CssEaseInOut)

	// Roll
	RollIn  = RotateGen(-120, 0, 0.5, 0.5).Merge(TranslateGen(ui.NewAmountPointParent(-1, 0), ui.AmountPoint{})).Merge(FadeIn)
	RollOut = RotateGen(0, 120, 0.5, 0.5).Merge(TranslateGen(ui.AmountPoint{}, ui.NewAmountPointParent(1, 0))).Merge(FadeOut)

	// ZoomIn
	zoomBezier0 = ease.NewBezier(0.550, 0.055, 0.675, 0.190)
	zoomBezier1 = ease.NewBezier(0.175, 0.885, 0.320, 1.000)

	ZoomInGen = func(bigX, bigY, smallX, smallY float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Scale: &gfx.Coord{X: 0.1, Y: 0.1}, Transparency: 1, Easing: zoomBezier0, Translate: ui.NewAmountPoint(bigX, bigY), Origin: OriginCenter},
			{Time: 0.60, Scale: &gfx.Coord{X: 0.475, Y: 0.475}, Easing: zoomBezier1, Translate: ui.NewAmountPoint(smallX, smallY), Origin: OriginCenter},
			{Time: 1.00, Scale: &gfx.Coord{X: 1, Y: 1}, Origin: OriginCenter},
		})
	}
	ZoomInLeft  = ZoomInGen(-1000, 0, 10, 0)
	ZoomInRight = ZoomInGen(1000, 0, -10, 0)
	ZoomInUp    = ZoomInGen(0, 1000, 0, -60)
	ZoomInDown  = ZoomInGen(0, -1000, 0, 60)
	ZoomIn      = AnimationGen([]ui.BasicAnimationFrame{
		{Time: 0.0, Scale: &gfx.Coord{X: 0.3, Y: 0.3}, Transparency: 1, Origin: OriginCenter},
		{Time: 0.5, Scale: &gfx.Coord{X: 0.5, Y: 0.5}, Origin: OriginCenter},
		{Time: 1.0, Scale: &gfx.Coord{X: 1.0, Y: 1.0}, Origin: OriginCenter},
	})

	// ZoomOut
	ZoomOutGen = func(bigX, bigY, smallX, smallY float32) ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00, Scale: &gfx.Coord{X: 1, Y: 1}, Origin: OriginCenter},
			{Time: 0.40, Scale: &gfx.Coord{X: 0.475, Y: 0.475}, Easing: zoomBezier1, Translate: ui.NewAmountPoint(smallX, smallY), Origin: ui.NewAmountPointParent(0.5, 0.75)},
			{Time: 1.00, Scale: &gfx.Coord{X: 0.1, Y: 0.1}, Transparency: 1, Translate: ui.NewAmountPoint(bigX, bigY), Origin: OriginBottom},
		})
	}
	ZoomOut = AnimationGen([]ui.BasicAnimationFrame{
		{Time: 0.0, Scale: &gfx.Coord{X: 1.0, Y: 1.0}, Origin: OriginCenter},
		{Time: 1.0, Scale: &gfx.Coord{X: 0.3, Y: 0.3}, Transparency: 1, Origin: OriginCenter},
	})
	ZoomOutDown  = ZoomOutGen(0, 2000, 0, -60)
	ZoomOutLeft  = ZoomOutGen(-2000, 0, 42, 0)
	ZoomOutRight = ZoomOutGen(2000, 0, -42, 0)
	ZoomOutUp    = ZoomOutGen(0, -2000, 0, 60)

	// Slide
	SlideInDown   = TranslateGen(ui.NewAmountPointParent(0, -1), ui.AmountPoint{})
	SlideInLeft   = TranslateGen(ui.NewAmountPointParent(-1, 0), ui.AmountPoint{})
	SlideInRight  = TranslateGen(ui.NewAmountPointParent(1, 0), ui.AmountPoint{})
	SlideInUp     = TranslateGen(ui.NewAmountPointParent(0, 1), ui.AmountPoint{})
	SlideOutDown  = TranslateGen(ui.AmountPoint{}, ui.NewAmountPointParent(0, 1)).Merge(FadeOutEnd)
	SlideOutLeft  = TranslateGen(ui.AmountPoint{}, ui.NewAmountPointParent(-1, 0)).Merge(FadeOutEnd)
	SlideOutRight = TranslateGen(ui.AmountPoint{}, ui.NewAmountPointParent(1, 0)).Merge(FadeOutEnd)
	SlideOutUp    = TranslateGen(ui.AmountPoint{}, ui.NewAmountPointParent(0, -1)).Merge(FadeOutEnd)

	/* Template
	TemplateGen = func() ui.BasicAnimation {
		return AnimationGen([]ui.BasicAnimationFrame{
			{Time: 0.00},
			{Time: 1.00},
		})
	}
	Template = TemplateGen()
	*/
)
