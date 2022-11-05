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
	POSITION_SCALE := DefineEntityType("position+scale", PositionScale{}, map[string]BaseComp{
		"Position": POSITION,
		"Scale":    SCALE,
	})

	w := NewWor(DataSettings{
		MaxEntityInstances:   64,
		MaxDestroysPerUpdate: 32,
		MaxNewPerUpdate:      32,
	})
	w.Add(POSITION)
	w.Add(SCALE)
	w.Add(POSITION_SCALE)

	e := w.New()
	pos := POSITION.Add(w, e)
	pos.X = 1
	pos.Y = 2
	fmt.Println("after position, before scale")
	scale := SCALE.Add(w, e)
	*scale = 4

	fmt.Println("before stage")
	w.Stage()
	fmt.Println("aftaer stage")
}
