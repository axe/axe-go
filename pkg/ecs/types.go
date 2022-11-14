package ecs

import (
	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/ds"
)

type DataID uint8

type DataOffset uint32

type ArrangementID uint16

type DataIDs = ds.Bits64Indexed[DataID]

const DATA_MAX = 64

type SystemContext struct {
	World *World
	Game  *axe.Game
}

type System[D any] interface {
	OnStage(data *D, e *Entity, ctx SystemContext)
	OnLive(data *D, e *Entity, ctx SystemContext)
	OnRemove(data *D, e *Entity, ctx SystemContext)
	Init(ctx SystemContext) error
	Update(iter ds.Iterable[EntityData[*D]], ctx SystemContext)
	Destroy(ctx SystemContext)
}

type EntitySystem interface {
	OnStage(e *Entity, ctx SystemContext)
	OnLive(e *Entity, ctx SystemContext)
	OnDelete(e *Entity, ctx SystemContext)
	Init(ctx SystemContext) error
	Update(iter ds.Iterable[Entity], ctx SystemContext)
	Destroy(ctx SystemContext)
}

type EntityData[D any] struct {
	Entity *Entity
	Data   D
}

type valueStaging struct {
	valueID     DataID
	valueOffset DataOffset
}

type dataLocation struct {
	dataID     DataID
	valueID    DataID
	live       bool
	dataOffset DataOffset
}

type dataValuePair struct {
	dataID      DataID
	valueID     DataID
	live        bool
	offsetIndex uint8
}

type worldDataStore interface {
	getValueIDs() DataIDs
	getDataSize() uintptr
	add(e *Entity) DataOffset
	remove(e *Entity, offset DataOffset)
	destroy()
}

type worldValueStore interface {
	init(ctx SystemContext) error
	update(ctx SystemContext)
	destroy(ctx SystemContext)
	unstage(e *Entity, stageOffset DataOffset, target DataID, targetOffset DataOffset, ctx SystemContext)
	remove(e *Entity, data DataID, dataOffset DataOffset, live bool, ctx SystemContext)
	move(sourceID DataID, sourceOffset DataOffset, targetID DataID, targetOffset DataOffset)
}

type worldValueData[V any] interface {
	get(offset DataOffset, live bool) *V
	free(offset DataOffset, live bool)
}

type dataValueData[D any] interface {
	addTo(w *World, data *worldDatas[D])
}

type dataValue[D any, V any] struct {
	dataId  DataID
	valueID DataID
	get     func(data *D) *V
}
