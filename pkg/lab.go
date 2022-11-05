package axe

import (
	"fmt"
	"math/bits"
	"reflect"
	"sort"
)

type BaseComp interface {
	ID() uint8
}

type DataSettings struct {
	MaxEntityInstances   uint32
	MaxDestroysPerUpdate uint32
	MaxNewPerUpdate      uint32
}

type ComponentSet uint64

func (set ComponentSet) Max() int {
	return 64 - bits.LeadingZeros64(uint64(set))
}
func (set ComponentSet) Size() int {
	return bits.OnesCount64(uint64(set))
}
func (set ComponentSet) Empty() bool {
	return set == 0
}
func (set *ComponentSet) Set(index uint8) {
	*set |= 1 << index
}
func (set *ComponentSet) Unset(index uint8) {
	*set &= ^(1 << index)
}
func (set ComponentSet) Has(index uint8) bool {
	return set&(1<<index) != 0
}
func (set *ComponentSet) Take() uint8 {
	x := uint8(bits.TrailingZeros64(uint64(*set)))
	*set ^= 1 << x
	return x
}

type EntityDataSource interface {
	ID() uint8
	Components() ComponentSet
	CreateData(capacity uint32, freeCapacity uint32) entityStorage
}

type Comp[T any] struct {
	Name    string
	Initial T

	id         uint8
	components ComponentSet
}

var nextComponentId uint8

func DefineComp[T any](name string, initial T) *Comp[T] {
	id := nextComponentId
	nextComponentId++

	return &Comp[T]{
		id:         id,
		components: ComponentSet(uint64(1) << id),
		Name:       name,
		Initial:    initial,
	}
}

var _ BaseComp = &Comp[int]{}
var _ EntityDataSource = &Comp[int]{}

func (c Comp[T]) ID() uint8 {
	return c.id
}

func (c Comp[T]) Components() ComponentSet {
	return c.components
}

func (c Comp[T]) AddSystem(w *Wor, system EntityDataSystem[T]) {
	datas := w.ComponentDatas[c.id]
	if len(datas) == 0 {
		panic(fmt.Sprintf("Error adding system to %s, you must add components to the world before systems.", c.Name))
	}
}

func (c Comp[T]) Get(w *Wor, e *Ent) *T {
	return c.get(w, e, false)
}

func (c Comp[T]) Add(w *Wor, e *Ent) *T {
	return c.get(w, e, true)
}

func (c Comp[T]) get(w *Wor, e *Ent, create bool) *T {
	has := e.components&c.components != 0
	if !has && !create {
		return nil
	}
	var data entityStorage
	var link *entLink
	if has {
		for i := range e.links {
			data = w.Data[e.links[i].dataID]
			if data.getComponents().Has(c.id) {
				link = &e.links[i]
				break
			}
		}
		if link == nil {
			return nil
		}
	} else {
		data = w.Staging[c.id]
		e.links = append(e.links, entLink{})
		link = &e.links[len(e.links)-1]
		data.add(e, link)
	}
	return data.get(c.id, *link).(*T)
}

func (c Comp[T]) CreateData(capacity uint32, freeCapacity uint32) entityStorage {
	data := newEntityDataPairs(capacity, freeCapacity, c.Initial)
	data.components = c.components
	data.getters = make([]EntityDataGetter[T], c.id+1)
	data.getters[c.id] = func(data *T) any {
		return data
	}
	return &data
}

type EntityDataGetter[D any] func(data *D) any

type EntType[D any] struct {
	Name    string
	Initial D

	id         uint8
	components ComponentSet
	getters    []EntityDataGetter[D]
}

var _ EntityDataSource = &EntType[int]{}

var nextTypeId uint8

func DefineEntityType[D any](name string, initial D) *EntType[D] {
	typeId := nextTypeId
	nextTypeId++

	return &EntType[D]{
		id:      typeId,
		Name:    name,
		getters: make([]EntityDataGetter[D], MAX_DATA),
	}
}

func (et *EntType[D]) AddComponent(comp BaseComp, field EntityDataGetter[D]) {
	if field != nil {
		et.components.Set(comp.ID())
	} else {
		et.components.Unset(comp.ID())
	}
	et.getters[comp.ID()] = field
}

func (et EntType[D]) ID() uint8 {
	return et.id
}

func (et EntType[D]) Components() ComponentSet {
	return et.components
}

func (et *EntType[D]) CreateData(capacity uint32, freeCapacity uint32) entityStorage {
	data := newEntityDataPairs(capacity, freeCapacity, et.Initial)
	data.components = et.components
	data.getters = et.getters
	return &data
}

// for lone component instances, all instances, or a type
type EntityDataSystem[T any] interface {
	OnStage(e *Ent, data *T)
	OnAdd(e *Ent, data *T)
	OnRemove(e *Ent, data *T)
	PreUpdate(game *Game) bool
	Update(game *Game, e *Ent, data *T) bool
	PostUpdate(game *Game)
}

// all entities, or components match
type EntitySystem interface {
	OnStage(e *Ent)
	OnNew(e *Ent)
	OnDestroy(e *Ent)
	PreUpdate(game *Game) bool
	Update(game *Game, e *Ent) bool
	PostUpdate(game *Game)
}

const MAX_DATA = 64

type Wor struct {
	Entity          []Ent
	EntityCount     uint32
	Staged          []*Ent
	StagedCount     uint32
	Free            []*Ent
	FreeCount       uint32
	Data            []entityStorage
	Staging         []entityStorage
	DataSettings    DataSettings
	ComponentToData []entityStorage
	TypeToData      []entityStorage
	TypeStores      []entityStorage
	ComponentDatas  [][]entityStorage
	ComponentStores []entityStorage
	Systems         []EntitySystem
}

func NewWor(settings DataSettings) *Wor {
	return &Wor{
		Entity:          make([]Ent, settings.MaxEntityInstances),
		EntityCount:     0,
		Staged:          make([]*Ent, settings.MaxNewPerUpdate),
		StagedCount:     0,
		Free:            make([]*Ent, settings.MaxDestroysPerUpdate),
		FreeCount:       0,
		Data:            make([]entityStorage, 0, MAX_DATA),
		DataSettings:    settings,
		Staging:         make([]entityStorage, MAX_DATA),
		ComponentToData: make([]entityStorage, MAX_DATA),
		TypeToData:      make([]entityStorage, MAX_DATA),
		TypeStores:      make([]entityStorage, 0, MAX_DATA),
		ComponentDatas:  make([][]entityStorage, MAX_DATA),
		ComponentStores: make([]entityStorage, 0, MAX_DATA),
		Systems:         make([]EntitySystem, 0),
	}
}

func (w *Wor) Update(game *Game) {
	for _, system := range w.Systems {
		if system.PreUpdate(game) {
			for i := uint32(0); i < w.EntityCount; i++ {
				e := &w.Entity[i]
				if e.Live() && !system.Update(game, e) {
					break
				}
			}
			system.PostUpdate(game)
		}
	}
}

func (w *Wor) AddSystem(system EntitySystem) {
	w.Systems = append(w.Systems, system)
}

func (w *Wor) Add(source EntityDataSource) {
	dataId := uint8(len(w.Data))

	data := source.CreateData(w.DataSettings.MaxEntityInstances, w.DataSettings.MaxDestroysPerUpdate)
	data.setID(dataId)
	w.Data = append(w.Data, data)

	components := source.Components()
	if components.Size() == 1 {
		componentId := components.Take()
		w.ComponentToData[componentId] = data
		w.ComponentStores = append(w.ComponentStores, data)
		w.addComponentData(componentId, data)

		staging := source.CreateData(w.DataSettings.MaxNewPerUpdate, w.DataSettings.MaxNewPerUpdate)
		staging.setID(componentId)
		w.Staging[componentId] = staging
	} else {
		w.TypeToData[source.ID()] = data
		w.TypeStores = append(w.TypeStores, data)
		for !components.Empty() {
			componentId := components.Take()
			w.addComponentData(componentId, data)
		}
		w.sortTypes()
	}
}

func (w *Wor) addComponentData(componentId uint8, data entityStorage) {
	if w.ComponentDatas[componentId] == nil {
		w.ComponentDatas[componentId] = make([]entityStorage, 0, MAX_DATA)
	}
	w.ComponentDatas[componentId] = append(w.ComponentDatas[componentId], data)
}

func (w *Wor) sortTypes() {
	sort.Slice(w.TypeStores, func(i, j int) bool {
		a := w.TypeStores[i]
		b := w.TypeStores[j]
		d := b.getComponents().Size() - a.getComponents().Size()
		if d == 0 {
			d = int(b.getDataSize() - a.getDataSize())
		}
		return d < 0
	})
}

func (w *Wor) New() *Ent {
	var e *Ent
	if w.FreeCount > 0 {
		e = w.Free[w.FreeCount]
		w.Free[w.FreeCount] = nil
		w.FreeCount--
	} else {
		e = &w.Entity[w.EntityCount]
		w.EntityCount++
		if e.links == nil {
			e.links = make([]entLink, 0)
		}
	}
	for _, system := range w.Systems {
		system.OnNew(e)
	}
	w.Staged[w.StagedCount] = e
	w.StagedCount++
	return e
}

func (w *Wor) Destroy(e *Ent) {
	for _, system := range w.Systems {
		system.OnDestroy(e)
	}
	for _, link := range e.links {
		data := w.Data[link.dataID]
		data.remove(e, &link)
	}
	e.links = e.links[:0]
	e.components = 0
	w.Free[w.FreeCount] = e
	w.FreeCount++
}

// Takes all staged entities and component values and puts them in place
func (w *Wor) Stage() {
	w.stageNewEntities()
	w.stageNewComponents()
}

func (w *Wor) stageNewEntities() {
	for i := uint32(0); i < w.StagedCount; i++ {
		staged := w.Staged[i]
		// setting the components to whats linked marks the entity as live
		for _, link := range staged.links {
			componentId := link.dataID
			staged.components.Set(componentId)
		}
		// based on the components, find the best way to store the data
		datas := w.getDataForComponents(staged.components)
		// for each staged component data, get the live storage and
		// copy the data over
		liveLinks := make([]entLink, 0)
		for _, data := range datas {
			liveLinks = append(liveLinks, entLink{})
			liveLink := &liveLinks[len(liveLinks)-1]
			data.add(staged, liveLink)

			components := data.getComponents()
			for !components.Empty() {
				componentId := components.Take()
				componentStaging := w.Staging[componentId]
				stageLink := staged.linkFor(componentId)
				stageValue := componentStaging.get(componentId, stageLink)
				dataValue := data.get(componentId, *liveLink)
				copy(dataValue, stageValue)
				componentStaging.remove(staged, &stageLink)
			}
		}
		staged.links = liveLinks
	}
	w.StagedCount = 0
}

func copy(dst any, src any) {
	d := reflect.ValueOf(dst)
	s := reflect.ValueOf(src)
	d.Elem().Set(s.Elem())
}

func (w *Wor) stageNewComponents() {
	// for _, componentStaging := range w.Staging {
	// 	pairs := componentStaging.getEntityOffsets()
	// 	for _, pair := range pairs {

	// 	}
	// }
}

func (w *Wor) getDataForComponents(components ComponentSet) []entityStorage {
	data := make([]entityStorage, 0)
	remaining := components
	for _, typeData := range w.TypeStores {
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
		data = append(data, w.ComponentToData[compomentId])
	}
	return data
}

type Ent struct {
	components ComponentSet
	links      []entLink
}

func (e Ent) Destroyed() bool {
	return e.components.Empty() && len(e.links) == 0
}

func (e Ent) Staging() bool {
	return e.components.Empty() && len(e.links) > 0
}

func (e Ent) Live() bool {
	return !e.components.Empty()
}

func (e Ent) Has(comp BaseComp) bool {
	return e.components.Has(comp.ID())
}

func (e Ent) linkFor(dataID uint8) entLink {
	for _, link := range e.links {
		if link.dataID == dataID {
			return link
		}
	}
	return entLink{dataID: dataID + 1}
}

type entLink struct {
	dataID     uint8
	dataOffset uint32
}

type entityDataPair[D any] struct {
	Data   D
	Entity *Ent
}

type entityOffsetPair struct {
	dataOffset uint32
	entity     *Ent
}

type entityStorage interface {
	add(e *Ent, link *entLink)
	remove(e *Ent, link *entLink)
	setID(id uint8)
	getID() uint8
	getComponents() ComponentSet
	getDataSize() uintptr
	get(component uint8, link entLink) any
	getEntityOffsets() []entityOffsetPair
}

type entityDataPairs[D any] struct {
	data      []entityDataPair[D]
	dataCount uint32
	dataSize  uintptr
	free      []uint32
	freeCount uint32
	initial   D

	id         uint8
	components ComponentSet
	getters    []EntityDataGetter[D]
}

var _ entityStorage = &entityDataPairs[int]{}

func newEntityDataPairs[D any](capacity uint32, freeCapacity uint32, initial D) entityDataPairs[D] {
	return entityDataPairs[D]{
		initial:  initial,
		data:     make([]entityDataPair[D], capacity),
		dataSize: reflect.TypeOf(initial).Size(),
		free:     make([]uint32, freeCapacity),
		getters:  make([]EntityDataGetter[D], MAX_DATA),
	}
}

func (ed *entityDataPairs[D]) setID(id uint8) {
	ed.id = id
}

func (ed entityDataPairs[D]) getID() uint8 {
	return ed.id
}

func (ed entityDataPairs[D]) getDataSize() uintptr {
	return ed.dataSize
}

func (ed entityDataPairs[D]) getEntityOffsets() []entityOffsetPair {
	pairs := make([]entityOffsetPair, 0, ed.dataCount)
	for i := uint32(0); i < ed.dataCount; i++ {
		data := ed.data[i]
		if data.Entity != nil {
			pairs = append(pairs, entityOffsetPair{
				dataOffset: i,
				entity:     data.Entity,
			})
		}
	}
	return pairs
}

func (ed *entityDataPairs[D]) add(e *Ent, link *entLink) {
	if ed.freeCount > 0 {
		ed.freeCount--
		link.dataOffset = ed.free[ed.freeCount]
	} else {
		link.dataOffset = ed.dataCount
		ed.dataCount++
	}
	ed.data[link.dataOffset].Data = ed.initial
	ed.data[link.dataOffset].Entity = e
	link.dataID = ed.id
}

func (ed *entityDataPairs[D]) remove(e *Ent, link *entLink) {
	if link.dataID != ed.id {
		return
	}
	ed.data[link.dataOffset].Entity = nil
	ed.free[ed.freeCount] = link.dataOffset
	ed.freeCount++
	link.dataID = 0
}

func (ed entityDataPairs[D]) getComponents() ComponentSet {
	return ed.components
}

func (ed *entityDataPairs[D]) get(componentId uint8, link entLink) any {
	data := &ed.data[link.dataOffset]
	getter := ed.getters[componentId]
	if getter != nil {
		return getter(&data.Data)
	}
	return nil
}
