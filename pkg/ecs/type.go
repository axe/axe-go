package ecs

import "fmt"

type Type[D any] struct {
	Name    string
	Initial D

	id         uint8
	components ComponentSet
	getters    []DataGetter[D]
}

var _ DataSource = &Type[int]{}

var nextTypeId uint8

func DefineType[D any](name string, initial D) *Type[D] {
	typeId := nextTypeId
	nextTypeId++

	return &Type[D]{
		id:      typeId,
		Name:    name,
		getters: make([]DataGetter[D], MAX_DATA),
	}
}

func (et Type[D]) ID() uint8 {
	return et.id
}

func (et Type[T]) GetName() string {
	return et.Name
}

func (et *Type[D]) AddComponent(comp BaseComponent, field DataGetter[D]) *Type[D] {
	if field != nil {
		et.components.Set(comp.ID())
	} else {
		et.components.Unset(comp.ID())
	}
	et.getters[comp.ID()] = field

	return et
}

func (et Type[D]) Components() ComponentSet {
	return et.components
}

func (et Type[T]) AddSystem(system DataSystem[T]) {
	data := ActiveWorld().typeToData[et.id]
	if data == nil {
		panic(fmt.Sprintf("Error adding system to %s, you must add types to the world before systems.", et.Name))
	}
	data.addSystem(dataSystemAny[T]{inner: system})
}

func (et *Type[T]) Enable() {
	ActiveWorld().Enable(et)
}

func (et *Type[D]) getComponentInstance(w *World, e *Entity, create bool) any {
	data := w.typeToData[et.id]
	if data == nil {
		return nil
	}
	link := e.linkFor(data.getID(), false)
	if link.dataID != data.getID() {
		return nil
	}
	if store, ok := data.(*store[D]); ok {
		value := store.data.At(link.dataOffset)

		return value
	}
	return nil
}

func (et *Type[D]) createStorage(capacity uint32, freeCapacity uint32) storage {
	data := newStore(capacity, freeCapacity, et.Initial, et.components, et.getters)
	return &data
}
