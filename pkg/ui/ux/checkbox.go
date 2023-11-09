package ux

import "github.com/axe/axe-go/pkg/ui"

type CheckboxStyles struct {
	CheckboxSize    ui.Coord
	CheckedTile     ui.Tile
	UncheckedTile   ui.Tile
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
