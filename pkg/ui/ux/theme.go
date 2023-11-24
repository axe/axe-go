package ux

import "github.com/axe/axe-go/pkg/ui"

type Theme struct {
	Settings          Settings
	ScrollSensitivity ui.Coord
	Templates         map[Kind]*ui.Template
}
