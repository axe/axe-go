package ecs

import (
	"fmt"
	"reflect"
)

type BaseComponent interface {
	ID() uint8
	GetName() string

	getComponentInstance(w *World, e *Entity, create bool) any
}

func get[T any](e *Entity, create bool) *T {
	world := ActiveWorld()
	var empty T
	source := world.sourceByType[reflect.TypeOf(empty)]
	if source == nil {
		return nil
	}
	value := source.getComponentInstance(world, e, create)
	return value.(*T)
}

func Get[T any](e *Entity) *T {
	return get[T](e, false)
}

func Add[T any](e *Entity) *T {
	return get[T](e, true)
}

func Set[T any](e *Entity, value T) bool {
	ptr := get[T](e, true)
	if ptr != nil {
		*ptr = value
		return true
	}
	return false
}

type Component[T any] struct {
	Name    string
	Initial T

	id         uint8
	components ComponentSet
}

var nextComponentId uint8

func DefineComponent[T any](name string, initial T) *Component[T] {
	id := nextComponentId
	nextComponentId++

	return &Component[T]{
		id:         id,
		components: ComponentSet(uint64(1) << id),
		Name:       name,
		Initial:    initial,
	}
}

var _ BaseComponent = &Component[int]{}
var _ DataSource = &Component[int]{}

func (c Component[T]) ID() uint8 {
	return c.id
}

func (c Component[T]) GetName() string {
	return c.Name
}

func (c Component[T]) Components() ComponentSet {
	return c.components
}

func (c Component[T]) AddSystem(system DataSystem[T]) {
	datas := ActiveWorld().componentDatas[c.id]
	if len(datas) == 0 {
		panic(fmt.Sprintf("Error adding system to %s, you must add components to the world before systems.", c.Name))
	}
	for _, data := range datas {
		data.addComponentSystem(c.id, dataSystemAny[T]{inner: system})
	}
}

func (c *Component[T]) Enable() {
	ActiveWorld().Enable(c)
}

func (c Component[T]) Get(e *Entity) *T {
	return c.get(ActiveWorld(), e, false)
}

func (c Component[T]) Add(e *Entity) *T {
	return c.get(ActiveWorld(), e, true)
}

func (c Component[T]) Set(e *Entity, value T) {
	ptr := c.get(ActiveWorld(), e, true)
	*ptr = value
}

func (c Component[T]) getComponentInstance(w *World, e *Entity, create bool) any {
	return c.get(w, e, create)
}

func (c Component[T]) get(w *World, e *Entity, create bool) *T {
	if e.components.Has(c.id) {
		return c.getLive(w, e)
	} else if e.staging.Has(c.id) {
		return c.getStaging(w, e)
	} else if create {
		return c.addStaging(w, e)
	}
	return nil
}

func (c Component[T]) getLive(w *World, e *Entity) *T {
	for i := range e.links {
		link := e.links[i]
		if link.staged {
			continue
		}
		data := w.data[link.dataID]
		if data.getComponents().Has(c.id) {
			return data.get(c.id, link).(*T)
		}
	}
	return nil
}

func (c Component[T]) getStaging(w *World, e *Entity) *T {
	data := w.staging[c.id]
	link := e.linkFor(c.id, true)

	return data.get(c.id, link).(*T)
}

func (c Component[T]) addStaging(w *World, e *Entity) *T {
	data := w.staging[c.id]
	e.links = append(e.links, dataLink{staged: true})
	e.staging.Set(c.id)
	link := &e.links[len(e.links)-1]
	data.add(w.ctx, e, link)
	value := data.get(c.id, *link).(*T)

	for _, data := range w.componentDatas[c.id] {
		data.onStage(w.ctx, e, value)
	}

	return value
}

func (c Component[T]) createStorage(capacity uint32, freeCapacity uint32) storage {
	getters := make([]DataGetter[T], c.id+1)
	getters[c.id] = func(data *T) any {
		return data
	}

	data := newStore(capacity, freeCapacity, c.Initial, c.components, getters)
	return &data
}
