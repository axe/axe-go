package ecs

import "github.com/axe/axe-go/pkg/ds"

type BaseComponent interface {
	Id() uint8
	Name() string
	Has(entity *Entity) bool

	free(index uint32)
	add(entity *Entity)
}

type Component[T any] struct {
	id        uint8
	name      string
	instances *ds.SparseList[T]
}

func (this *Component[T]) Id() uint8 {
	return this.id
}

func (this *Component[T]) Name() string {
	return this.name
}

func (this *Component[T]) free(index uint32) {
	this.instances.Free(index)
}

func (this *Component[T]) add(entity *Entity) {
	var value T
	this.Set(entity, value)
}

func (this *Component[T]) Get(entity *Entity) *T {
	if this.Has(entity) {
		return this.instances.At(entity.components[this.id])
	}
	return nil
}

func (this *Component[T]) Has(entity *Entity) bool {
	return (entity.has & (1 << this.id)) != 0
}

func (this *Component[T]) Set(entity *Entity, value T) bool {
	if this.Has(entity) {
		ref := this.instances.At(entity.components[this.id])
		*ref = value

		return true
	} else {
		if len(entity.components) <= int(this.id) {
			entity.components = entity.components[:(this.id + 1)]
		}
		entity.components[this.id] = uint32(this.instances.Add(value))
		entity.has |= (1 << this.id)

		return false
	}
}

func (this *Component[T]) Remove(entity *Entity) bool {
	if this.Has(entity) {
		this.instances.Free(entity.components[this.id])
		entity.has &= ^(1 << this.id)
		return true
	} else {
		return false
	}
}

func DefineComponent[T any](world *World, name string) *Component[T] {
	component := new(Component[T])
	component.id = uint8(len(world.components))
	component.name = name
	component.instances = ds.NewSparseList[T](world.componentInstanceSize, world.componentFreeSize)

	world.components = append(world.components, component)

	return component
}
