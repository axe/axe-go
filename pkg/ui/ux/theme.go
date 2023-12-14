package ux

import (
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/ui"
)

type Theme struct {
	Settings          Settings
	ScrollSensitivity gfx.Coord
	Templates         map[Kind]*ui.Template
}
