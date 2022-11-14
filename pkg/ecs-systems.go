package axe

import (
	"github.com/axe/axe-go/pkg/ds"
)

type entitySystemFiltered struct {
	search WorldSearch
	inner  EntitySystem
}

var _ EntitySystem = &entitySystemFiltered{}

func NewEntitySystemFiltered(sys EntitySystem, filter WorldSearch) EntitySystem {
	return entitySystemFiltered{filter, sys}
}

func (sys entitySystemFiltered) OnStage(e *Entity, ctx EntityContext) {
	if sys.search.IsMatch(e) {
		sys.inner.OnStage(e, ctx)
	}
}

func (sys entitySystemFiltered) OnLive(e *Entity, ctx EntityContext) {
	if sys.search.IsMatch(e) {
		sys.inner.OnLive(e, ctx)
	}
}

func (sys entitySystemFiltered) OnDelete(e *Entity, ctx EntityContext) {
	if sys.search.IsMatch(e) {
		sys.inner.OnDelete(e, ctx)
	}
}

func (sys entitySystemFiltered) Update(iterable ds.Iterable[Entity], ctx EntityContext) {
	filteredIterable := ds.NewFilterIterable(iterable, sys.search.IsMatch)

	sys.inner.Update(filteredIterable, ctx)
}

func (sys entitySystemFiltered) Init(ctx EntityContext) error {
	return sys.inner.Init(ctx)
}

func (sys entitySystemFiltered) Destroy(ctx EntityContext) {
	sys.inner.Destroy(ctx)
}
