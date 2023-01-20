package ecs

import (
	"github.com/axe/axe-go/pkg/ds"
)

type systemFiltered struct {
	search WorldSearch
	inner  System
}

var _ System = &systemFiltered{}

func NewSystemFiltered(sys System, filter WorldSearch) System {
	return systemFiltered{filter, sys}
}

func (sys systemFiltered) OnStage(e *Entity, ctx Context) {
	if sys.search.IsMatch(e) {
		sys.inner.OnStage(e, ctx)
	}
}

func (sys systemFiltered) OnLive(e *Entity, ctx Context) {
	if sys.search.IsMatch(e) {
		sys.inner.OnLive(e, ctx)
	}
}

func (sys systemFiltered) OnDelete(e *Entity, ctx Context) {
	if sys.search.IsMatch(e) {
		sys.inner.OnDelete(e, ctx)
	}
}

func (sys systemFiltered) Update(iterable ds.Iterable[Entity], ctx Context) {
	filteredIterable := ds.NewFilterIterable(iterable, sys.search.IsMatch)

	sys.inner.Update(filteredIterable, ctx)
}

func (sys systemFiltered) Init(ctx Context) error {
	return sys.inner.Init(ctx)
}

func (sys systemFiltered) Destroy(ctx Context) {
	sys.inner.Destroy(ctx)
}
