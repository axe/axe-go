package ui_test

import (
	"fmt"
	"testing"

	"github.com/axe/axe-go/pkg/ui"
)

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
