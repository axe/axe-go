package ecs

import "github.com/axe/axe-go/pkg/util"

type Entity struct {
	id         uint32
	components ComponentSet
	staging    ComponentSet
	links      []dataLink
}

func NewEntity() *Entity {
	return ActiveWorld().New()
}

func (e Entity) ID() uint32 {
	return e.id
}

func (e Entity) Deleted() bool {
	return e.links == nil
}

func (e Entity) Staging() bool {
	return e.components.Empty()
}

func (e Entity) Live() bool {
	return !e.components.Empty()
}

func (e Entity) Has(comp BaseComponent) bool {
	return (e.components | e.staging).Has(comp.ID())
}

func (e Entity) HasLive(comp BaseComponent) bool {
	return e.components.Has(comp.ID())
}

func (e Entity) HasStaging(comp BaseComponent) bool {
	return e.staging.Has(comp.ID())
}

func (e Entity) Components() ComponentSet {
	return e.components
}

func (e Entity) StagingComponents() ComponentSet {
	return e.staging
}

func (e *Entity) Get(comp BaseComponent) any {
	return comp.getComponentInstance(ActiveWorld(), e, false)
}

func (e *Entity) Add(comp BaseComponent) any {
	return comp.getComponentInstance(ActiveWorld(), e, true)
}

func (e *Entity) Set(comp BaseComponent, value any) {
	ptr := comp.getComponentInstance(ActiveWorld(), e, true)
	if ptr != nil {
		util.Copy(ptr, &value)
	}
}

func (e *Entity) Delete() {
	ActiveWorld().Delete(e)
}

func (e Entity) linkFor(dataID uint8, staged bool) dataLink {
	for _, link := range e.links {
		if link.dataID == dataID && link.staged == staged {
			return link
		}
	}
	return dataLink{dataID: dataID + 1}
}

func (e *Entity) removeLink(match dataLink) {
	for i, link := range e.links {
		if link == match {
			e.links = util.SliceRemoveAt(e.links, i)
			return
		}
	}
}
