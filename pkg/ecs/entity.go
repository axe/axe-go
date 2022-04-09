package ecs

type Entity struct {
	ID         uint32
	Components []uint32
	Has        uint64
	Flags      uint64
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
