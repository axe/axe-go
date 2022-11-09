package ecs

import (
	"fmt"
	"testing"

	"github.com/axe/axe-go/pkg/geom"
)

/*
1. Define components
2. Define types
3. Create world
4. Enable components in world
5. Enable types in world
6. Add systems
7. Populate & update world
*/

func TestWorld(t *testing.T) {

	type PositionScale struct {
		Position geom.Vec3f
		Scale    float32
	}

	POSITION := DefineComponent("position", geom.Vec3f{})
	SCALE := DefineComponent("scale", float32(0))

	POSITION_SCALE := DefineType("position+scale", PositionScale{})
	POSITION_SCALE.AddComponent(POSITION, func(data *PositionScale) any { return &data.Position })
	POSITION_SCALE.AddComponent(SCALE, func(data *PositionScale) any { return &data.Scale })

	w := NewWorld(WorldSettings{
		MaxEntityInstances:   64,
		MaxDestroysPerUpdate: 32,
		MaxNewPerUpdate:      32,
	})
	w.Enable(POSITION)
	w.Enable(SCALE)
	w.Enable(POSITION_SCALE)

	e := w.New()
	fmt.Printf("entity %s\n", statusOf(e))
	pos := POSITION.Add(e)
	pos.X = 1
	pos.Y = 2
	scale := SCALE.Add(e)
	*scale = 4

	fmt.Printf("entity %s\n", statusOf(e))
	w.Stage()

	fmt.Printf("entity %s\n", statusOf(e))
	fmt.Printf("entity pos %+v\n", *POSITION.Get(e))
	fmt.Printf("entity scale %+v\n", *SCALE.Get(e))

	e2 := w.New()
	pos2 := POSITION.Add(e2)
	pos2.X = 34
	w.Stage()
	fmt.Printf("entity2 pos %+v\n", *POSITION.Get(e2))

	scale2 := SCALE.Add(e2)
	*scale2 = 86
	fmt.Printf("entity2 staging components %+v\n", e2.StagingComponents())
	w.Stage()
	fmt.Printf("entity2 scale %+v\n", *SCALE.Get(e2))
	fmt.Printf("entity2 staging components %+v\n", e2.StagingComponents())
}

func statusOf(e *Entity) string {
	if e.Deleted() {
		return "destroyed"
	} else if e.Staging() {
		return "staging"
	} else if e.Live() {
		return "live"
	} else {
		return "unknown"
	}
}
