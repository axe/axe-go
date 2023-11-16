package ui_test

import (
	"fmt"
	"testing"

	"github.com/axe/axe-go/pkg/ui"
)

func TestColorModify(t *testing.T) {
	if ui.Alpha(1).HasAffect() {
		t.Errorf("Alpha(1).HasAffect")
	}
}

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

func TestFlagsTake(t *testing.T) {
	var flags ui.Flags = 2 | 4 | 16 | 128
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
}

func TestVertexBufferClear(t *testing.T) {
	vbs := ui.BufferPool.Get()
	vb := vbs.Buffer()
	vb.Add(ui.Vertex{X: 1}, ui.Vertex{X: 2}, ui.Vertex{X: 3})
	vbs.Clear()
	vb1 := vbs.Buffer()
	vb1.Clear()
}

const EPSILON = 0.00001

func floatEqual(a, b float32) bool {
	if a > b {
		return (a - b) < EPSILON
	} else {
		return (b - a) < EPSILON
	}
}

func coordEqual(a, b ui.Coord) bool {
	return floatEqual(a.X, b.X) && floatEqual(a.Y, b.Y)
}

func TestTranslate(t *testing.T) {
	cases := []struct {
		X, Y, Tx, Ty, Ex, Ey float32
	}{
		{X: 0, Y: 0, Tx: 1, Ty: 0, Ex: 1, Ey: 0},
	}

	for tcIndex, tc := range cases {
		t.Run(fmt.Sprintf("#%d", tcIndex), func(t *testing.T) {
			te := ui.Transform{}
			te.Identity()
			te.Translate(tc.Tx, tc.Ty)
			ax, ay := te.Transform(tc.X, tc.Y)

			if ax != tc.Ex || ay != tc.Ey {
				t.Errorf("Failed Identity/Translate")
			}

			ts := ui.Transform{}
			ts.SetTranslate(tc.Tx, tc.Ty)
			ax, ay = ts.Transform(tc.X, tc.Y)

			if ax != tc.Ex || ay != tc.Ey {
				t.Errorf("Failed SetTranslate")
			}
		})
	}
}

func TestRotate(t *testing.T) {
	type InputExpected struct {
		Input, Expected ui.Coord
	}

	cases := []struct {
		Build func() ui.Transform
		Tests []InputExpected
	}{{
		Build: func() ui.Transform {
			tr := ui.Transform{}
			tr.SetRotateDegrees(90)
			return tr
		},
		Tests: []InputExpected{{
			Input:    ui.Coord{X: 10, Y: 0},
			Expected: ui.Coord{X: 0, Y: 10},
		}},
	}}

	for tcIndex, tc := range cases {
		t.Run(fmt.Sprintf("#%d", tcIndex), func(t *testing.T) {
			tr := tc.Build()

			for _, test := range tc.Tests {
				actual := tr.TransformCoord(test.Input)
				if !coordEqual(actual, test.Expected) {
					t.Errorf("actual %+v != expected %+v", actual, test.Expected)
				}
			}
		})
	}
}
