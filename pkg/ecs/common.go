package ecs

import (
	"reflect"
)

func get[T any](e *Entity, create bool) *T {
	world := ActiveWorld()
	var empty T
	valueID := world.typeMap[reflect.TypeOf(empty)]
	datasBase := world.datasBase[valueID]
	if datasBase == nil {
		return nil
	}
	datas := datasBase.(*Data[T])
	return datas.get(e, create)
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

func DefineComponent[C any](name string, initial C) *Data[C] {
	return newData(name, initial)
}

func DefineType[T any](name string, initial T, with func(t *Data[T])) *Data[T] {
	data := newData(name, initial)
	with(data)
	return data
}

func DefineTypeComponent[T any, C any](t *Data[T], c *Data[C], get func(data *T) *C) {
	t.values[c.id] = &dataValue[T, C]{
		dataId:  t.id,
		valueID: c.id,
		get:     get,
	}
}
