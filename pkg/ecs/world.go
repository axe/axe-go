package ecs

import (
	"reflect"
	"sort"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

type WorldSettings struct {
	MaxEntityInstances   uint32
	MaxDestroysPerUpdate uint32
	MaxNewPerUpdate      uint32
	RestructureOnChange  bool
}

const MAX_DATA = 64

type World struct {
	// Settings for the data in this world.
	DataSettings WorldSettings
	// List of entities in this world.
	Entity ds.SparseList[Entity]
	// Stack of staged entities.
	staged ds.Stack[*Entity]
	// All data storage, where the index is the generated dataId.
	data []storage
	// All staging components, where the index is the component id.
	staging []storage
	// Component storage, indexed by component id.
	componentToData []storage
	// Component storage list, in order of being added to the world.
	componentStores []storage
	// Stores by component id that contain the component.
	componentDatas [][]storage
	// Type storage, indexed by type id.
	typeToData []storage
	// Type storage list, in order of being added to the world.
	typeStores []storage
	// ecs.Get[Transform](e)
	sourceByType map[reflect.Type]DataSource
	// All global entity systems.
	systems []System
	// The cached system context for this world.
	ctx SystemContext
}

var _ axe.GameSystem = &World{}

type WorldSearch struct {
	Match                   util.Match[ComponentSet]
	IncludeStagedComponents bool
	IncludeStaged           bool
}

func NewWorld(settings WorldSettings) *World {
	world := &World{
		DataSettings:    settings,
		Entity:          ds.NewSparseList[Entity](settings.MaxEntityInstances),
		staged:          ds.NewStack[*Entity](settings.MaxNewPerUpdate),
		data:            make([]storage, 0, MAX_DATA),
		staging:         make([]storage, MAX_DATA),
		componentToData: make([]storage, MAX_DATA),
		typeToData:      make([]storage, MAX_DATA),
		typeStores:      make([]storage, 0, MAX_DATA),
		componentDatas:  make([][]storage, MAX_DATA),
		componentStores: make([]storage, 0, MAX_DATA),
		sourceByType:    make(map[reflect.Type]DataSource, MAX_DATA),
		systems:         make([]System, 0),
	}

	world.ctx.World = world
	activeWorld = world

	return world
}

var activeWorld *World

func ActiveWorld() *World {
	if activeWorld == nil {
		panic("There is no active world, you must create one or activate one")
	}
	return activeWorld
}

func (w *World) IsActive() bool {
	return w == activeWorld
}

func (w *World) Activate() {
	activeWorld = w
}

func (w *World) Init(game *axe.Game) error {
	w.Activate()

	w.ctx.Game = game

	for _, sys := range w.systems {
		err := sys.Init(w.ctx)
		if err != nil {
			return err
		}
	}

	for _, data := range w.data {
		err := data.init(w.ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *World) Destroy() {
	w.Activate()

	for _, sys := range w.systems {
		sys.Destroy(w.ctx)
	}

	for _, data := range w.data {
		data.destroy(w.ctx)
	}

	if w.IsActive() {
		activeWorld = nil
	}
}

func (w *World) Update(game *axe.Game) {
	w.Activate()

	w.Stage()

	for _, system := range w.systems {
		if system.PreUpdate(w.ctx) {
			w.Entity.Iterate(func(item *Entity, index, liveIndex uint32) bool {
				if item.Live() {
					return system.Update(item, w.ctx)
				}
				return true
			})
			system.PostUpdate(w.ctx)
		}
	}

	for _, data := range w.data {
		data.update(w.ctx)
	}

	w.Shrink()
}

func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
}

func (w *World) AddComponentsSystem(match util.Match[ComponentSet], system System) {
	w.systems = append(w.systems, systemFiltered{
		match: match,
		inner: system,
	})
}

func (w *World) Enable(source DataSource) {
	dataId := uint8(len(w.data))

	data := source.createStorage(w.DataSettings.MaxEntityInstances, w.DataSettings.MaxDestroysPerUpdate)
	data.setID(dataId)
	w.data = append(w.data, data)

	components := source.Components()
	if components.Size() == 1 {
		componentId := components.Take()
		w.componentToData[componentId] = data
		w.componentStores = append(w.componentStores, data)
		w.addComponentData(componentId, data)

		staging := source.createStorage(w.DataSettings.MaxNewPerUpdate, w.DataSettings.MaxNewPerUpdate)
		staging.setID(componentId)
		w.staging[componentId] = staging
	} else {
		w.typeToData[source.ID()] = data
		w.typeStores = append(w.typeStores, data)
		for !components.Empty() {
			componentId := components.Take()
			w.addComponentData(componentId, data)
		}
		w.sortTypes()
	}

	w.sourceByType[data.getType()] = source
}

func (w *World) Search(search WorldSearch, found func(e *Entity, index uint32) bool) {
	w.Activate()

	w.Entity.Iterate(func(item *Entity, index, liveIndex uint32) bool {
		components := item.components
		if search.IncludeStagedComponents {
			components |= item.staging
		}
		if search.Match == nil || search.Match(components) {
			if item.Live() || search.IncludeStaged {
				return found(item, liveIndex)
			}
		}
		return true
	})
}

func (w *World) New() *Entity {
	w.Activate()

	e, id := w.Entity.Take()
	e.id = id
	e.links = make([]dataLink, 0)
	for _, system := range w.systems {
		system.OnNew(e, w.ctx)
	}
	w.staged.Push(e)
	return e
}

func (w *World) Delete(e *Entity) {
	if e.Deleted() {
		return
	}
	w.Activate()

	for _, system := range w.systems {
		system.OnDestroy(e, w.ctx)
	}
	for _, link := range e.links {
		data := w.data[link.dataID]
		data.remove(w.ctx, e, &link)
	}
	e.links = nil
	e.components = 0
	w.Entity.Free(e.id)
}

// Takes all staged entities and component values and puts them in place
func (w *World) Stage() {
	w.Activate()
	w.stageNewEntities()
	w.stageNewComponents()
}

func (w *World) Shrink() {
	// w.Entity.Compress()
	// for _, data := range w.data {
	// data.shrink()
	// }
}

func (w *World) stageNewEntities() {
	for !w.staged.IsEmpty() {
		staged := w.staged.Pop()
		// setting the components to whats linked marks the entity as live
		for _, link := range staged.links {
			componentId := link.dataID
			staged.components.Set(componentId)
		}

		// based on the components, find the best way to store the data
		datas := w.getDataForComponents(staged.components)

		// for each staged component data, get the live storage and
		// copy the data over
		liveLinks := make([]dataLink, 0)
		for _, data := range datas {
			liveLinks = append(liveLinks, dataLink{})
			liveLink := &liveLinks[len(liveLinks)-1]
			data.add(w.ctx, staged, liveLink)

			components := data.getComponents()
			for !components.Empty() {
				componentId := components.Take()
				componentStaging := w.staging[componentId]
				stageLink := staged.linkFor(componentId, true)
				stageValue := componentStaging.get(componentId, stageLink)
				dataValue := data.get(componentId, *liveLink)
				util.Copy(dataValue, stageValue)
				componentStaging.remove(w.ctx, staged, &stageLink)
			}
		}
		staged.links = liveLinks
		staged.staging = 0

		for _, system := range w.systems {
			system.OnStage(staged, w.ctx)
		}
	}
}

func (w *World) stageNewComponents() {
	if w.DataSettings.RestructureOnChange && len(w.typeStores) > 0 {
		// TODO
	} else {
		for _, componentStaging := range w.staging {
			if componentStaging == nil {
				continue
			}
			pairs := componentStaging.getEntityOffsets()
			componentId := componentStaging.getID()

			for _, pair := range pairs {
				liveData := w.componentToData[componentId]
				liveLink := dataLink{}
				liveData.add(w.ctx, pair.entity, &liveLink)
				dataValue := liveData.get(componentId, liveLink)
				stageLink := dataLink{
					dataID:     componentId,
					dataOffset: pair.dataOffset,
					staged:     true,
				}
				stageValue := componentStaging.get(componentId, stageLink)
				util.Copy(dataValue, stageValue)
				componentStaging.remove(w.ctx, pair.entity, &stageLink)
				pair.entity.components.Set(componentId)
				pair.entity.staging.Unset(componentId)
				pair.entity.removeLink(stageLink)
				pair.entity.links = append(pair.entity.links, liveLink)
			}
		}
	}
}

func (w *World) getDataForComponents(components ComponentSet) []storage {
	data := make([]storage, 0)
	remaining := components
	for _, typeData := range w.typeStores {
		typeComponents := typeData.getComponents()
		if typeComponents&remaining == typeComponents {
			remaining ^= typeComponents
			data = append(data, typeData)
			if remaining.Empty() {
				break
			}
		}
	}
	for !remaining.Empty() {
		compomentId := remaining.Take()
		data = append(data, w.componentToData[compomentId])
	}
	return data
}

func (w *World) addComponentData(componentId uint8, data storage) {
	if w.componentDatas[componentId] == nil {
		w.componentDatas[componentId] = make([]storage, 0, MAX_DATA)
	}
	w.componentDatas[componentId] = append(w.componentDatas[componentId], data)
}

func (w *World) sortTypes() {
	sort.Slice(w.typeStores, func(i, j int) bool {
		a := w.typeStores[i]
		b := w.typeStores[j]
		d := b.getComponents().Size() - a.getComponents().Size()
		if d == 0 {
			d = int(b.getType().Size() - a.getType().Size())
		}
		return d < 0
	})
}
