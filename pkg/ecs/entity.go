package ecs

type Entity struct {
	id         uint32
	components []int
	has        uint64
	Flags      uint64
}

func (e Entity) Id() uint32 {
	return e.id
}
func (e Entity) Has(component BaseComponent) bool {
	return component.Has(&e)
}

type EntityCreate struct {
	Components []uint8
	Flags      uint64
}

type EntitySearch struct {
	Flags      uint64
	FlagMatch  FlagMatch
	Components []uint8
}
