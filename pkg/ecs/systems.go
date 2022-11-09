package ecs

import (
	"github.com/axe/axe-go/pkg/util"
)

type systemFiltered struct {
	match util.Match[ComponentSet]
	inner System
}

var _ System = &systemFiltered{}

func (sys systemFiltered) OnStage(e *Entity, ctx SystemContext) {
	if sys.match(e.components) {
		sys.inner.OnStage(e, ctx)
	}
}
func (sys systemFiltered) OnNew(e *Entity, ctx SystemContext) {
	if sys.match(e.components) {
		sys.inner.OnNew(e, ctx)
	}
}
func (sys systemFiltered) OnDestroy(e *Entity, ctx SystemContext) {
	if sys.match(e.components) {
		sys.inner.OnDestroy(e, ctx)
	}
}
func (sys systemFiltered) PreUpdate(ctx SystemContext) bool {
	return sys.inner.PreUpdate(ctx)
}
func (sys systemFiltered) Update(e *Entity, ctx SystemContext) bool {
	if sys.match(e.components) {
		return sys.inner.Update(e, ctx)
	}
	return true
}
func (sys systemFiltered) PostUpdate(ctx SystemContext) {
	sys.inner.PostUpdate(ctx)
}
func (sys systemFiltered) Init(ctx SystemContext) error {
	return sys.inner.Init(ctx)
}
func (sys systemFiltered) Destroy(ctx SystemContext) {
	sys.inner.Destroy(ctx)
}

type dataSystemAny[T any] struct {
	inner DataSystem[T]
}

var _ DataSystem[any] = &dataSystemAny[int]{}

func (sys dataSystemAny[T]) OnStage(data *any, e *Entity, ctx SystemContext) {
	sys.inner.OnStage((*data).(*T), e, ctx)
}
func (sys dataSystemAny[T]) OnAdd(data *any, e *Entity, ctx SystemContext) {
	sys.inner.OnAdd((*data).(*T), e, ctx)
}
func (sys dataSystemAny[T]) OnRemove(data *any, e *Entity, ctx SystemContext) {
	sys.inner.OnRemove((*data).(*T), e, ctx)
}
func (sys dataSystemAny[T]) PreUpdate(ctx SystemContext) bool {
	return sys.inner.PreUpdate(ctx)
}
func (sys dataSystemAny[T]) Update(data *any, e *Entity, ctx SystemContext) bool {
	return sys.inner.Update((*data).(*T), e, ctx)
}
func (sys dataSystemAny[T]) PostUpdate(ctx SystemContext) {
	sys.inner.PostUpdate(ctx)
}
func (sys dataSystemAny[T]) Init(ctx SystemContext) error {
	return sys.inner.Init(ctx)
}
func (sys dataSystemAny[T]) Destroy(ctx SystemContext) {
	sys.inner.Destroy(ctx)
}
