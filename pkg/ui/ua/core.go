package ua

import (
	"github.com/axe/axe-go/pkg/ease"
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/ui"
)

var (
	DefaultDuration float32 = 1.0
	DefaultSave             = true
	DefaultEasing           = ease.Linear

	Named = id.NewDenseKeyMap[ui.AnimationFactory, uint16, uint8](
		id.WithStringMap(map[string]ui.AnimationFactory{
			// animatecss
			"bounce":             Bounce,
			"rubberband":         RubberBand,
			"flash":              Flash,
			"pulse":              Pulse,
			"shake":              Shake,
			"swing":              Swing,
			"tada":               Tada,
			"wobble":             Wobble,
			"fadein":             FadeIn,
			"fadeout":            FadeOut,
			"bounceindown":       BounceInDown,
			"bounceinleft":       BounceInLeft,
			"bounceinright":      BounceInRight,
			"bounceinup":         BounceInUp,
			"bouncein":           BounceIn,
			"bounceoutdown":      BounceOutDown,
			"bounceoutleft":      BounceOutLeft,
			"bounceoutright":     BounceOutRight,
			"bounceoutup":        BounceOutUp,
			"bounceout":          BounceOut,
			"fadeindown":         FadeInDown,
			"fadeindownbig":      FadeInDownBig,
			"fadeinleft":         FadeInLeft,
			"fadeinleftbig":      FadeInLeftBig,
			"fadeinright":        FadeInRight,
			"fadeinrightbig":     FadeInRightBig,
			"fadeinup":           FadeInUp,
			"fadeinupbig":        FadeInUpBig,
			"fadeoutdown":        FadeOutDown,
			"fadeoutdownbig":     FadeOutDownBig,
			"fadeoutleft":        FadeOutLeft,
			"fadeoutleftbig":     FadeOutLeftBig,
			"fadeoutright":       FadeOutRight,
			"fadeoutrightbig":    FadeOutRightBig,
			"fadeoutup":          FadeOutUp,
			"fadeoutupbig":       FadeOutUpBig,
			"rotatein":           RotateIn,
			"rotateindownleft":   RotateInDownLeft,
			"rotateindownright":  RotateInDownRight,
			"rotateinupleft":     RotateInUpLeft,
			"rotateinupright":    RotateInUpRight,
			"rotateout":          RotateOut,
			"rotateoutdownleft":  RotateOutDownLeft,
			"rotateoutdownright": RotateOutDownRight,
			"rotateoutupleft":    RotateOutUpLeft,
			"rotateoutupright":   RotateOutUpRight,
			"hinge":              Hinge,
			"rollin":             RollIn,
			"rollout":            RollOut,
			"zoomin":             ZoomIn,
			"zoominleft":         ZoomInLeft,
			"zoominright":        ZoomInRight,
			"zoominup":           ZoomInUp,
			"zoomindown":         ZoomInDown,
			"zoomout":            ZoomOut,
			"zoomoutdown":        ZoomOutDown,
			"zoomoutleft":        ZoomOutLeft,
			"zoomoutright":       ZoomOutRight,
			"zoomoutup":          ZoomOutUp,
			"slideindown":        SlideInDown,
			"slideinleft":        SlideInLeft,
			"slideinright":       SlideInRight,
			"slideinup":          SlideInUp,
			"slideoutdown":       SlideOutDown,
			"slideoutleft":       SlideOutLeft,
			"slideoutright":      SlideOutRight,
			"slideoutup":         SlideOutUp,
			// anim8js
			"wiggle": Wiggle,
			// axe
			"explode":         Explode,
			"revealdown":      RevealDown,
			"revealup":        RevealUp,
			"revealupdown":    RevealUpDown,
			"revealleft":      RevealLeft,
			"revealright":     RevealRight,
			"revealleftright": RevealLeftRight,
		}),
	)

	Animations = ui.Animations{
		Named: Named,
	}

	AnimationGen = func(frames []ui.BasicAnimationFrame) ui.BasicAnimation {
		return ui.BasicAnimation{
			Duration: DefaultDuration,
			Save:     DefaultSave,
			Easing:   DefaultEasing,
			Frames:   frames,
		}
	}

	// Common
	OriginTop         = ui.NewAmountPointParent(0.5, 0)
	OriginCenter      = ui.NewAmountPointParent(0.5, 0.5)
	OriginBottom      = ui.NewAmountPointParent(0.5, 1.0)
	OriginTopLeft     = ui.NewAmountPointParent(0, 0)
	OriginTopRight    = ui.NewAmountPointParent(1, 0)
	OriginCenterRight = ui.NewAmountPointParent(1, 0.5)
	OriginBottomLeft  = ui.NewAmountPointParent(0, 1)
	OriginBottomRight = ui.NewAmountPointParent(1, 1)
	OriginCenterLeft  = ui.NewAmountPointParent(0, 0.5)
)
