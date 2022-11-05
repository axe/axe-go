package axe

import (
	"fmt"
	"testing"
)

/*
1. Define components
2. Define types
3. Create world
4. Add components to world
5. Add types to world
6. Add systems
7. Populate & update world
*/

func TestWor(t *testing.T) {

	type PositionScale struct {
		Position Vec3f
		Scale    float32
	}

	POSITION := DefineComp("position", Vec3f{})
	SCALE := DefineComp("scale", float32(0))

	POSITION_SCALE := DefineEntityType("position+scale", PositionScale{})
	POSITION_SCALE.AddComponent(POSITION, func(data *PositionScale) any { return &data.Position })
	POSITION_SCALE.AddComponent(SCALE, func(data *PositionScale) any { return &data.Scale })

	w := NewWor(DataSettings{
		MaxEntityInstances:   64,
		MaxDestroysPerUpdate: 32,
		MaxNewPerUpdate:      32,
	})
	w.Add(POSITION)
	w.Add(SCALE)
	w.Add(POSITION_SCALE)

	e := w.New()
	fmt.Printf("entity %s\n", statusOf(e))
	pos := POSITION.Add(w, e)
	pos.X = 1
	pos.Y = 2
	scale := SCALE.Add(w, e)
	*scale = 4

	fmt.Printf("entity %s\n", statusOf(e))
	w.Stage()

	fmt.Printf("entity %s\n", statusOf(e))
	fmt.Printf("entity pos %+v\n", *POSITION.Get(w, e))
	fmt.Printf("entity scale %+v\n", *SCALE.Get(w, e))

	e2 := w.New()
	pos2 := POSITION.Add(w, e2)
	pos2.X = 34
	w.Stage()
	fmt.Printf("entity2 pos %+v\n", *POSITION.Get(w, e2))

	scale2 := SCALE.Add(w, e2)
	*scale2 = 86
	w.Stage()
	fmt.Printf("entity2 scale %+v\n", *SCALE.Get(w, e2))
}

func statusOf(e *Ent) string {
	if e.Destroyed() {
		return "destroyed"
	} else if e.Staging() {
		return "staging"
	} else if e.Live() {
		return "live"
	} else {
		return "unknown"
	}
}
