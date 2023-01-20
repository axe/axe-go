package ecs

import (
	"github.com/axe/axe-go/pkg/ds"
)

type DataID uint8

type DataOffset uint32

type ArrangementID uint16

type DataIDs = ds.Bits64Indexed[DataID]

const DATA_MAX = 64

type Context struct {
	World *World
	// Game  *Game
}

type DataSystem[D any] interface {
	OnStage(data *D, e *Entity, ctx Context)
	OnLive(data *D, e *Entity, ctx Context)
	OnRemove(data *D, e *Entity, ctx Context)
	Init(ctx Context) error
	Update(iter ds.Iterable[Value[*D]], ctx Context)
	Destroy(ctx Context)
}

type System interface {
	OnStage(e *Entity, ctx Context)
	OnLive(e *Entity, ctx Context)
	OnDelete(e *Entity, ctx Context)
	Init(ctx Context) error
	Update(iter ds.Iterable[Entity], ctx Context)
	Destroy(ctx Context)
}

type Value[D any] struct {
	Data D
	ID   ID
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
	init(ctx Context) error
	update(ctx Context)
	destroy(ctx Context)
	unstage(e *Entity, stageOffset DataOffset, target DataID, targetOffset DataOffset, ctx Context)
	remove(e *Entity, data DataID, dataOffset DataOffset, live bool, ctx Context)
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
