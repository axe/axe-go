package ecs

import (
	"reflect"

	"github.com/axe/axe-go/pkg/ds"
)

type store[D any] struct {
	data             ds.SparseList[entityDataPair[D]]
	initial          D
	id               uint8
	components       ComponentSet
	getters          []DataGetter[D]
	componentSystems []DataSystem[any]
	systems          []DataSystem[any]
	reflectType      reflect.Type
}

var _ storage = &store[int]{}

func newStore[D any](capacity uint32, freeCapacity uint32, initial D, components ComponentSet, getters []DataGetter[D]) store[D] {
	reflectType := reflect.TypeOf(initial)

	return store[D]{
		initial:          initial,
		data:             ds.NewSparseList[entityDataPair[D]](capacity),
		getters:          getters,
		components:       components,
		componentSystems: make([]DataSystem[any], components.Max()),
		systems:          make([]DataSystem[any], 0),
		reflectType:      reflectType,
	}
}

func (ed *store[D]) setID(id uint8) {
	ed.id = id
}

func (ed store[D]) getID() uint8 {
	return ed.id
}

func (ed store[D]) getType() reflect.Type {
	return ed.reflectType
}

func (ed store[D]) getEntityOffsets() []entityOffsetPair {
	pairs := make([]entityOffsetPair, 0, ed.data.Size())
	ed.data.Iterate(func(item *entityDataPair[D], index, liveIndex uint32) bool {
		pairs = append(pairs, entityOffsetPair{
			dataOffset: index,
			entity:     item.entity,
		})
		return true
	})
	return pairs
}

func (ed *store[D]) init(ctx SystemContext) error {
	for _, sys := range ed.systems {
		err := sys.Init(ctx)
		if err != nil {
			return err
		}
	}

	for _, componentSys := range ed.componentSystems {
		if componentSys == nil {
			continue
		}
		err := componentSys.Init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ed *store[D]) destroy(ctx SystemContext) {
	for _, sys := range ed.systems {
		sys.Destroy(ctx)
	}

	for _, componentSys := range ed.componentSystems {
		if componentSys == nil {
			continue
		}
		componentSys.Destroy(ctx)
	}

	// TODO call other system events per entity?
}

func (ed *store[D]) add(ctx SystemContext, e *Entity, link *dataLink) {
	item, index := ed.data.Take()
	item.data = ed.initial
	item.entity = e
	link.dataID = ed.id
	link.dataOffset = index

	itemData := any(&item.data)

	for _, sys := range ed.systems {
		sys.OnAdd(&itemData, e, ctx)
	}

	for componentId, componentSys := range ed.componentSystems {
		getter := ed.getters[componentId]
		if componentSys == nil || getter == nil {
			continue
		}
		componentInstance := getter(&item.data)
		componentSys.OnAdd(&componentInstance, e, ctx)
	}
}

func (ed *store[D]) remove(ctx SystemContext, e *Entity, link *dataLink) {
	pair := ed.data.At(link.dataOffset)
	if link.dataID != ed.id || pair.entity != e {
		return
	}

	itemData := any(&pair.data)

	for _, sys := range ed.systems {
		sys.OnRemove(&itemData, e, ctx)
	}

	for componentId, componentSys := range ed.componentSystems {
		getter := ed.getters[componentId]
		if componentSys == nil || getter == nil {
			continue
		}
		componentInstance := getter(&pair.data)
		componentSys.OnRemove(&componentInstance, e, ctx)
	}

	ed.data.Free(link.dataOffset)
	pair.entity = nil
	link.dataID = 0
}

func (ed store[D]) getComponents() ComponentSet {
	return ed.components
}

func (ed *store[D]) get(componentId uint8, link dataLink) any {
	data := ed.data.At(link.dataOffset)
	getter := ed.getters[componentId]
	if getter != nil {
		return getter(&data.data)
	}
	return nil
}

func (ed *store[D]) addComponentSystem(component uint8, system DataSystem[any]) {
	ed.componentSystems[component] = system
}

func (ed *store[D]) addSystem(system DataSystem[any]) {
	ed.systems = append(ed.systems, system)
}

func (ed *store[D]) update(ctx SystemContext) {
	for _, sys := range ed.systems {
		if sys.PreUpdate(ctx) {
			ed.data.Iterate(func(item *entityDataPair[D], index, liveIndex uint32) bool {
				itemData := any(&item.data)

				return sys.Update(&itemData, item.entity, ctx)
			})

			sys.PostUpdate(ctx)
		}
	}

	for componentId, componentSys := range ed.componentSystems {
		if componentSys == nil {
			continue
		}
		getter := ed.getters[componentId]
		if getter == nil {
			continue
		}
		if componentSys.PreUpdate(ctx) {
			ed.data.Iterate(func(item *entityDataPair[D], index, liveIndex uint32) bool {
				componentInstance := getter(&item.data)

				return componentSys.Update(&componentInstance, item.entity, ctx)
			})

			componentSys.PostUpdate(ctx)
		}
	}
}

func (ed *store[D]) onStage(ctx SystemContext, e *Entity, data any) {
	for _, componentSys := range ed.componentSystems {
		if componentSys == nil {
			continue
		}
		componentSys.OnStage(&data, e, ctx)
	}
}
