package axe

import (
	"github.com/axe/axe-go/pkg/ecs"
)

type World struct {
	Impl *ecs.World
}

var _ GameSystem = &World{}

func NewWorld(name string, settings ecs.WorldSettings) *World {
	return &World{Impl: ecs.NewWorld(name, settings)}
}

func (w *World) Init(game *Game) error {
	return w.Impl.Init()
}

func (w *World) Update(game *Game) {
	w.Impl.Update()
}

func (w *World) Destroy() {
	w.Impl.Destroy()
}
