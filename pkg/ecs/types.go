package ecs

import (
	"reflect"

	axe "github.com/axe/axe-go/pkg"
)

type SystemContext struct {
	Game  *axe.Game
	World *World
}

type DataSource interface {
	ID() uint8
	GetName() string
	Components() ComponentSet

	createStorage(capacity uint32, freeCapacity uint32) storage
	getComponentInstance(w *World, e *Entity, create bool) any
}

type DataGetter[D any] func(data *D) any

// for lone component instances, all instances, or a type
type DataSystem[T any] interface {
	OnStage(data *T, e *Entity, ctx SystemContext)
	OnAdd(data *T, e *Entity, ctx SystemContext)
	OnRemove(data *T, e *Entity, ctx SystemContext)
	PreUpdate(ctx SystemContext) bool
	Update(data *T, e *Entity, ctx SystemContext) bool
	PostUpdate(ctx SystemContext)
	Init(ctx SystemContext) error
	Destroy(ctx SystemContext)
}

// all entities, or components match
type System interface {
	OnStage(e *Entity, ctx SystemContext)
	OnNew(e *Entity, ctx SystemContext)
	OnDestroy(e *Entity, ctx SystemContext)
	PreUpdate(ctx SystemContext) bool
	Update(e *Entity, ctx SystemContext) bool
	PostUpdate(ctx SystemContext)
	Init(ctx SystemContext) error
	Destroy(ctx SystemContext)
}

type dataLink struct {
	dataID     uint8
	dataOffset uint32
	staged     bool
}

type entityDataPair[D any] struct {
	data   D
	entity *Entity
}

type entityOffsetPair struct {
	dataOffset uint32
	entity     *Entity
}

type storage interface {
	init(ctx SystemContext) error
	destroy(ctx SystemContext)
	add(ctx SystemContext, e *Entity, link *dataLink)
	remove(ctx SystemContext, e *Entity, link *dataLink)
	update(ctx SystemContext)
	onStage(ctx SystemContext, e *Entity, data any)
	setID(id uint8)
	getID() uint8
	getComponents() ComponentSet
	get(component uint8, link dataLink) any
	addComponentSystem(component uint8, system DataSystem[any])
	addSystem(system DataSystem[any])
	getEntityOffsets() []entityOffsetPair
	getType() reflect.Type
}
