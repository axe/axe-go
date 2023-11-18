package ux

import "github.com/axe/axe-go/pkg/ui"

type Theme struct {
	Styles            Styles
	ScrollSensitivity ui.Coord
	Templates         map[Kind]*ui.Template
}
