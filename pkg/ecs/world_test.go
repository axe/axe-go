package ecs

import (
	"fmt"
	"testing"
)

type Transform struct {
	x, y float32
}

// === WORLD ===
var WORLD = NewWorld(WorldOptions{})

// === COMPONENTS ===
var TRANSFORM = DefineComponent[Transform](WORLD, "transform")
var TAG = DefineComponent[string](WORLD, "tag")

// === TYPES ===
var BULLET = EntityCreate{
	Flags:      3,
	Components: []uint8{TRANSFORM.id},
}

func TestECS(t *testing.T) {
	p0 := WORLD.Create()
	p0.Flags = 1
	TRANSFORM.Set(p0, Transform{0, 1})

	p1 := WORLD.Create()
	p1.Flags = 2
	TRANSFORM.Set(p1, Transform{2, 3})

	p2 := WORLD.Create()
	p2.Flags = 1

	results0 := []uint32{}
	WORLD.Search(EntitySearch{}, func(entity *Entity) {
		results0 = append(results0, entity.ID)
	})

	fmt.Println(results0)

	results1 := []uint32{}
	WORLD.Search(EntitySearch{Flags: 1}, func(entity *Entity) {
		results1 = append(results1, entity.ID)
	})

	fmt.Println(results1)

	results2 := []uint32{}
	WORLD.Search(EntitySearch{Components: []uint8{TRANSFORM.id}}, func(entity *Entity) {
		results2 = append(results2, entity.ID)
	})

	fmt.Println(results2)

	results3 := []uint32{}
	WORLD.Search(EntitySearch{Flags: 1, FlagMatch: MATCH_NONE}, func(entity *Entity) {
		results3 = append(results3, entity.ID)
	})

	fmt.Println(results3)

	WORLD.Destroy(p2)

	results4 := []uint32{}
	WORLD.Search(EntitySearch{}, func(entity *Entity) {
		results0 = append(results4, entity.ID)
	})

	fmt.Println(results4)

	p4 := WORLD.CreateWith(BULLET)

	fmt.Println(TRANSFORM.Get(p4))
}
