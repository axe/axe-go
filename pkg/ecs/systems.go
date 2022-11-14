package ecs

import (
	"github.com/axe/axe-go/pkg/ds"
)

type systemFiltered struct {
	search WorldSearch
	inner  EntitySystem
}

var _ EntitySystem = &systemFiltered{}

func (sys systemFiltered) OnStage(e *Entity, ctx SystemContext) {
	if sys.search.IsMatch(e) {
		sys.inner.OnStage(e, ctx)
	}
}

func (sys systemFiltered) OnLive(e *Entity, ctx SystemContext) {
	if sys.search.IsMatch(e) {
		sys.inner.OnLive(e, ctx)
	}
}

func (sys systemFiltered) OnDelete(e *Entity, ctx SystemContext) {
	if sys.search.IsMatch(e) {
		sys.inner.OnDelete(e, ctx)
	}
}

func (sys systemFiltered) Update(iterable ds.Iterable[Entity], ctx SystemContext) {
	filteredIterable := ds.NewFilterIterable(iterable, sys.search.IsMatch)

	sys.inner.Update(filteredIterable, ctx)
}

func (sys systemFiltered) Init(ctx SystemContext) error {
	return sys.inner.Init(ctx)
}

func (sys systemFiltered) Destroy(ctx SystemContext) {
	sys.inner.Destroy(ctx)
}
