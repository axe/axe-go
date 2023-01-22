package axe

import (
	"github.com/axe/axe-go/pkg/ecs"
)

type World struct {
	ecs.World
}

var _ GameSystem = &World{}

func NewWorld(name string, settings ecs.WorldSettings) *World {
	return &World{World: *ecs.NewWorld(name, settings)}
}

func (w *World) Init(game *Game) error {
	return w.World.Init()
}

func (w *World) Update(game *Game) {
	w.World.Update()
}
