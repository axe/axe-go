package ua

import "github.com/axe/axe-go/pkg/ui"

var (

	// Wiggle
	Wiggle = AnimationGen([]ui.BasicAnimationFrame{
		{Time: 0.000, Origin: OriginCenter, Rotate: 0},
		{Time: 0.125, Origin: OriginCenter, Rotate: -45},
		{Time: 0.375, Origin: OriginCenter, Rotate: 45},
		{Time: 0.583, Origin: OriginCenter, Rotate: -30},
		{Time: 0.750, Origin: OriginCenter, Rotate: 30},
		{Time: 0.875, Origin: OriginCenter, Rotate: -15},
		{Time: 0.958, Origin: OriginCenter, Rotate: 15},
		{Time: 1.000, Origin: OriginCenter, Rotate: 0},
	})

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
