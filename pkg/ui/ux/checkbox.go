package ux

import (
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/ui"
)

type CheckboxSettings struct {
	CheckboxSize    gfx.Coord
	CheckedTile     gfx.Tile
	UncheckedTile   gfx.Tile
	CheckedLayers   []ui.Layer
	UncheckedLayers []ui.Layer
}

type Checkbox struct {
	Label        string
	LabelValue   Value[string]
	Checked      bool
	CheckedValue Value[bool]

	OnCheck   func()
	OnUncheck func()
	OnChange  func(checked bool)
}
