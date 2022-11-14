package axe

import (
	"github.com/axe/axe-go/pkg/ds"
)

type EntityDataID uint8

type EntityDataOffset uint32

type EntityArrangementID uint16

type EntityDataIDs = ds.Bits64Indexed[EntityDataID]

const ENTITY_DATA_MAX = 64

type EntityContext struct {
	World *World
	Game  *Game
}

type EntityDataSystem[D any] interface {
	OnStage(data *D, e *Entity, ctx EntityContext)
	OnLive(data *D, e *Entity, ctx EntityContext)
	OnRemove(data *D, e *Entity, ctx EntityContext)
	Init(ctx EntityContext) error
	Update(iter ds.Iterable[EntityValue[*D]], ctx EntityContext)
	Destroy(ctx EntityContext)
}

type EntitySystem interface {
	OnStage(e *Entity, ctx EntityContext)
	OnLive(e *Entity, ctx EntityContext)
	OnDelete(e *Entity, ctx EntityContext)
	Init(ctx EntityContext) error
	Update(iter ds.Iterable[Entity], ctx EntityContext)
	Destroy(ctx EntityContext)
}

type EntityValue[D any] struct {
	Data   D
	Entity *Entity
}

type entityValueStaging struct {
	valueID     EntityDataID
	valueOffset EntityDataOffset
}

type entityDataLocation struct {
	dataID     EntityDataID
	valueID    EntityDataID
	live       bool
	dataOffset EntityDataOffset
}

type entityDataValuePair struct {
	dataID      EntityDataID
	valueID     EntityDataID
	live        bool
	offsetIndex uint8
}

type worldDataStore interface {
	getValueIDs() EntityDataIDs
	getDataSize() uintptr
	add(e *Entity) EntityDataOffset
	remove(e *Entity, offset EntityDataOffset)
	destroy()
}

type worldValueStore interface {
	init(ctx EntityContext) error
	update(ctx EntityContext)
	destroy(ctx EntityContext)
	unstage(e *Entity, stageOffset EntityDataOffset, target EntityDataID, targetOffset EntityDataOffset, ctx EntityContext)
	remove(e *Entity, data EntityDataID, dataOffset EntityDataOffset, live bool, ctx EntityContext)
	move(sourceID EntityDataID, sourceOffset EntityDataOffset, targetID EntityDataID, targetOffset EntityDataOffset)
}

type worldValueData[V any] interface {
	get(offset EntityDataOffset, live bool) *V
	free(offset EntityDataOffset, live bool)
}

type entityDataValueData[D any] interface {
	addTo(w *World, data *worldDatas[D])
}

type entityDataValue[D any, V any] struct {
	dataId  EntityDataID
	valueID EntityDataID
	get     func(data *D) *V
}
