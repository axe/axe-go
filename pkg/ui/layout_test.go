package ui_test

import (
	"fmt"
	"testing"

	"github.com/axe/axe-go/pkg/ui"
)

func TestLayoutGrid(t *testing.T) {
	parent := ui.Bounds{Right: 300, Bottom: 300}
	text := ui.MustTextToVisual("{h:0.5}{pv:0.5}Toggle FullHeight")
	base := &ui.Base{
		Placement: ui.Maximized(),
		Layout: ui.LayoutGrid{
			FullHeight:          false,
			FullWidth:           false,
			VerticalAlignment:   ui.AlignmentCenter,
			HorizontalAlignment: ui.AlignmentCenter,
			VerticalSpacing:     ui.Amount{Value: 10},
			HorizontalSpacing:   ui.Amount{Value: 10},
			Columns:             3,
			MinSize:             ui.Coord{X: 80, Y: 80},
			AspectRatio:         0,
		},
		Children: []*ui.Base{{
			Layers: []ui.Layer{{
				Placement: ui.Maximized().Shrink(8),
				Visual:    text,
			}},
		}},
	}

	ctx := getContextWithFont("roboto")
	base.Place(ctx, parent, true)

	fmt.Printf("Layoutable #1: %+v\n", base.Children[0].Bounds)
	fmt.Printf("Layoutable #1 Visual: %+v\n", text.Paragraphs.Measure(ctx))
}
