package axe

import (
	"fmt"
	"math/bits"
	"reflect"
	"sort"

	"github.com/axe/axe-go/pkg/util"
)

type BaseComponent interface {
	ID() uint8
	GetName() string
}

type WorldSettings struct {
	MaxEntityInstances   uint32
	MaxDestroysPerUpdate uint32
	MaxNewPerUpdate      uint32
	RestructureOnChange  bool
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
	GetName() string
	Components() ComponentSet

	createData(capacity uint32, freeCapacity uint32) entityStorage
}

type Component[T any] struct {
	Name    string
	Initial T

	id         uint8
	components ComponentSet
}

var nextComponentId uint8

func DefineComponent[T any](name string, initial T) *Component[T] {
	id := nextComponentId
	nextComponentId++

	return &Component[T]{
		id:         id,
		components: ComponentSet(uint64(1) << id),
		Name:       name,
		Initial:    initial,
	}
}

var _ BaseComponent = &Component[int]{}
var _ EntityDataSource = &Component[int]{}

func (c Component[T]) ID() uint8 {
	return c.id
}

func (c Component[T]) GetName() string {
	return c.Name
}

func (c Component[T]) Components() ComponentSet {
	return c.components
}

func (c Component[T]) AddSystem(w *World, system EntityDataSystem[T]) {
	datas := w.ComponentDatas[c.id]
	if len(datas) == 0 {
		panic(fmt.Sprintf("Error adding system to %s, you must add components to the world before systems.", c.Name))
	}
}

func (c Component[T]) Get(w *World, e *Entity) *T {
	return c.get(w, e, false)
}

func (c Component[T]) Add(w *World, e *Entity) *T {
	return c.get(w, e, true)
}

func (c Component[T]) get(w *World, e *Entity, create bool) *T {
	if e.components.Has(c.id) {
		return c.getLive(w, e)
	} else if e.staging.Has(c.id) {
		return c.getStaging(w, e)
	} else if create {
		return c.addStaging(w, e)
	}
	return nil
}

func (c Component[T]) getLive(w *World, e *Entity) *T {
	for i := range e.links {
		data := w.Data[e.links[i].dataID]
		if data.getComponents().Has(c.id) {
			return data.get(c.id, e.links[i]).(*T)
		}
	}
	return nil
}

func (c Component[T]) getStaging(w *World, e *Entity) *T {
	data := w.Staging[c.id]
	link := e.linkFor(c.id)
	return data.get(c.id, link).(*T)
}

func (c Component[T]) addStaging(w *World, e *Entity) *T {
	data := w.Staging[c.id]
	e.links = append(e.links, entityDataLink{})
	e.staging.Set(c.id)
	link := &e.links[len(e.links)-1]
	data.add(e, link)
	return data.get(c.id, *link).(*T)
}

func (c Component[T]) createData(capacity uint32, freeCapacity uint32) entityStorage {
	data := newEntityStore(capacity, freeCapacity, c.Initial)
	data.components = c.components
	data.getters = make([]EntityDataGetter[T], c.id+1)
	data.getters[c.id] = func(data *T) any {
		return data
	}
	return &data
}

type EntityDataGetter[D any] func(data *D) any

type EntityType[D any] struct {
	Name    string
	Initial D

	id         uint8
	components ComponentSet
	getters    []EntityDataGetter[D]
}

var _ EntityDataSource = &EntityType[int]{}

var nextTypeId uint8

func DefineEntityType[D any](name string, initial D) *EntityType[D] {
	typeId := nextTypeId
	nextTypeId++

	return &EntityType[D]{
		id:      typeId,
		Name:    name,
		getters: make([]EntityDataGetter[D], MAX_DATA),
	}
}

func (et EntityType[D]) ID() uint8 {
	return et.id
}

func (et EntityType[T]) GetName() string {
	return et.Name
}

func (et *EntityType[D]) AddComponent(comp BaseComponent, field EntityDataGetter[D]) {
	if field != nil {
		et.components.Set(comp.ID())
	} else {
		et.components.Unset(comp.ID())
	}
	et.getters[comp.ID()] = field
}

func (et EntityType[D]) Components() ComponentSet {
	return et.components
}

func (et *EntityType[D]) createData(capacity uint32, freeCapacity uint32) entityStorage {
	data := newEntityStore(capacity, freeCapacity, et.Initial)
	data.components = et.components
	data.getters = et.getters
	return &data
}

// for lone component instances, all instances, or a type
type EntityDataSystem[T any] interface {
	OnStage(e *Entity, data *T)
	OnAdd(e *Entity, data *T)
	OnRemove(e *Entity, data *T)
	PreUpdate(game *Game) bool
	Update(game *Game, e *Entity, data *T) bool
	PostUpdate(game *Game)
}

// all entities, or components match
type EntitySystem interface {
	OnStage(e *Entity)
	OnNew(e *Entity)
	OnDestroy(e *Entity)
	PreUpdate(game *Game) bool
	Update(game *Game, e *Entity) bool
	PostUpdate(game *Game)
}

const MAX_DATA = 64

type World struct {
	// Settings for the data in this world.
	DataSettings WorldSettings
	// List of entities in this world.
	Entity sparseArray[Entity]
	// Stack of Staged entities.
	Staged stack[*Entity]
	// Stack of entities that have been freed that are ready for reuse.
	Free stack[*Entity]
	// All data storage, where the index is the generated dataId.
	Data []entityStorage
	// All staging components, where the index is the component id.
	Staging []entityStorage
	// Component storage, indexed by component id.
	ComponentToData []entityStorage
	// Component storage list, in order of being added to the world.
	ComponentStores []entityStorage
	// Stores by component id that contain the component.
	ComponentDatas [][]entityStorage
	// Type storage, indexed by type id.
	TypeToData []entityStorage
	// Type storage list, in order of being added to the world.
	TypeStores []entityStorage
	// All global entity systems.
	Systems []EntitySystem
}

func NewWorld(settings WorldSettings) *World {
	return &World{
		DataSettings:    settings,
		Entity:          newSparseArray[Entity](settings.MaxEntityInstances),
		Staged:          newStack[*Entity](settings.MaxNewPerUpdate),
		Free:            newStack[*Entity](settings.MaxDestroysPerUpdate),
		Data:            make([]entityStorage, 0, MAX_DATA),
		Staging:         make([]entityStorage, MAX_DATA),
		ComponentToData: make([]entityStorage, MAX_DATA),
		TypeToData:      make([]entityStorage, MAX_DATA),
		TypeStores:      make([]entityStorage, 0, MAX_DATA),
		ComponentDatas:  make([][]entityStorage, MAX_DATA),
		ComponentStores: make([]entityStorage, 0, MAX_DATA),
		Systems:         make([]EntitySystem, 0),
	}
}

func (w *World) Update(game *Game) {
	w.Stage()

	for _, system := range w.Systems {
		if system.PreUpdate(game) {
			count := w.Entity.count
			entities := w.Entity.items
			for i := uint32(0); i < count; i++ {
				e := &entities[i]
				if e.Live() && !system.Update(game, e) {
					break
				}
			}
			system.PostUpdate(game)
		}
	}

	// TODO other systems

	w.Shrink()
}

func (w *World) AddSystem(system EntitySystem) {
	w.Systems = append(w.Systems, system)
}

type entitySystemFiltered struct {
	match util.Match[ComponentSet]
	inner EntitySystem
}

func (sys entitySystemFiltered) OnStage(e *Entity) {
	if sys.match(e.components) {
		sys.inner.OnStage(e)
	}
}
func (sys entitySystemFiltered) OnNew(e *Entity) {
	if sys.match(e.components) {
		sys.inner.OnNew(e)
	}
}
func (sys entitySystemFiltered) OnDestroy(e *Entity) {
	if sys.match(e.components) {
		sys.inner.OnDestroy(e)
	}
}
func (sys entitySystemFiltered) PreUpdate(game *Game) bool {
	return sys.inner.PreUpdate(game)
}
func (sys entitySystemFiltered) Update(game *Game, e *Entity) bool {
	if sys.match(e.components) {
		return sys.inner.Update(game, e)
	}
	return true
}
func (sys entitySystemFiltered) PostUpdate(game *Game) {
	sys.inner.PostUpdate(game)
}

func (w *World) AddComponentsSystem(match util.Match[ComponentSet], system EntitySystem) {
	w.Systems = append(w.Systems, entitySystemFiltered{
		match: match,
		inner: system,
	})
}

func (w *World) Add(source EntityDataSource) {
	dataId := uint8(len(w.Data))

	data := source.createData(w.DataSettings.MaxEntityInstances, w.DataSettings.MaxDestroysPerUpdate)
	data.setID(dataId)
	w.Data = append(w.Data, data)

	components := source.Components()
	if components.Size() == 1 {
		componentId := components.Take()
		w.ComponentToData[componentId] = data
		w.ComponentStores = append(w.ComponentStores, data)
		w.addComponentData(componentId, data)

		staging := source.createData(w.DataSettings.MaxNewPerUpdate, w.DataSettings.MaxNewPerUpdate)
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

func (w *World) addComponentData(componentId uint8, data entityStorage) {
	if w.ComponentDatas[componentId] == nil {
		w.ComponentDatas[componentId] = make([]entityStorage, 0, MAX_DATA)
	}
	w.ComponentDatas[componentId] = append(w.ComponentDatas[componentId], data)
}

func (w *World) sortTypes() {
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

func (w *World) New() *Entity {
	var e *Entity
	if !w.Free.empty() {
		e = w.Free.pop()
	} else {
		e = w.Entity.take()
	}
	e.links = make([]entityDataLink, 0)
	for _, system := range w.Systems {
		system.OnNew(e)
	}
	w.Staged.push(e)
	return e
}

func (w *World) Destroy(e *Entity) {
	if e.Destroyed() {
		return
	}
	for _, system := range w.Systems {
		system.OnDestroy(e)
	}
	for _, link := range e.links {
		data := w.Data[link.dataID]
		data.remove(e, &link)
	}
	e.links = nil
	e.components = 0
	w.Free.push(e)
}

// Takes all staged entities and component values and puts them in place
func (w *World) Stage() {
	w.stageNewEntities()
	w.stageNewComponents()
}

func (w *World) Shrink() {
	w.Entity.shrink()
	for _, data := range w.Data {
		data.shrink()
	}
}

func (w *World) stageNewEntities() {
	for i := 0; i < w.Staged.count; i++ {
		staged := w.Staged.items[i]
		// setting the components to whats linked marks the entity as live
		for _, link := range staged.links {
			componentId := link.dataID
			staged.components.Set(componentId)
		}
		// based on the components, find the best way to store the data
		datas := w.getDataForComponents(staged.components)
		// for each staged component data, get the live storage and
		// copy the data over
		liveLinks := make([]entityDataLink, 0)
		for _, data := range datas {
			liveLinks = append(liveLinks, entityDataLink{})
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
		staged.staging = 0

		for _, system := range w.Systems {
			system.OnStage(staged)
		}
	}
	w.Staged.count = 0
}

func (w *World) stageNewComponents() {
	if w.DataSettings.RestructureOnChange && len(w.TypeStores) > 0 {
		// TODO
	} else {
		for _, componentStaging := range w.Staging {
			if componentStaging == nil {
				continue
			}
			pairs := componentStaging.getEntityOffsets()
			componentId := componentStaging.getID()

			for _, pair := range pairs {
				liveData := w.ComponentToData[componentId]
				liveLink := entityDataLink{}
				liveData.add(pair.entity, &liveLink)
				dataValue := liveData.get(componentId, liveLink)
				stageLink := entityDataLink{
					dataID:     componentId,
					dataOffset: pair.dataOffset,
				}
				stageValue := componentStaging.get(componentId, stageLink)
				copy(dataValue, stageValue)
				componentStaging.remove(pair.entity, &stageLink)
				pair.entity.components.Set(componentId)
				pair.entity.staging.Unset(componentId)
				pair.entity.removeLink(componentId, pair.dataOffset)
				pair.entity.links = append(pair.entity.links, liveLink)
			}
		}
	}
}

func (w *World) getDataForComponents(components ComponentSet) []entityStorage {
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

type Entity struct {
	components ComponentSet
	staging    ComponentSet
	links      []entityDataLink
}

func (e Entity) Destroyed() bool {
	return e.links == nil
}

func (e Entity) Staging() bool {
	return e.components.Empty()
}

func (e Entity) Live() bool {
	return !e.components.Empty()
}

func (e Entity) Has(comp BaseComponent) bool {
	return e.components.Has(comp.ID())
}

func (e Entity) Components() ComponentSet {
	return e.components
}

func (e Entity) StagingComponents() ComponentSet {
	return e.staging
}

func (e Entity) linkFor(dataID uint8) entityDataLink {
	for _, link := range e.links {
		if link.dataID == dataID {
			return link
		}
	}
	return entityDataLink{dataID: dataID + 1}
}

func (e *Entity) removeLink(dataID uint8, dataOffset uint32) {
	for i, link := range e.links {
		if link.dataID == dataID && link.dataOffset == dataOffset {
			e.links = sliceRemoveAt(e.links, i)
			return
		}
	}
}

func (e Entity) isEmpty() bool {
	return e.Destroyed()
}

type entityDataLink struct {
	dataID     uint8
	dataOffset uint32
}

type entityDataPair[D any] struct {
	data   D
	entity *Entity
}

func (pair entityDataPair[D]) isEmpty() bool {
	return pair.entity == nil
}

type entityOffsetPair struct {
	dataOffset uint32
	entity     *Entity
}

type entityStorage interface {
	add(e *Entity, link *entityDataLink)
	remove(e *Entity, link *entityDataLink)
	setInitial(initial any)
	setID(id uint8)
	getID() uint8
	getComponents() ComponentSet
	getDataSize() uintptr
	get(component uint8, link entityDataLink) any
	getEntityOffsets() []entityOffsetPair
	shrink()
}

type entityStore[D any] struct {
	data       sparseArray[entityDataPair[D]]
	dataSize   uintptr
	free       stack[uint32]
	initial    D
	id         uint8
	components ComponentSet
	getters    []EntityDataGetter[D]
}

var _ entityStorage = &entityStore[int]{}

func isEntityDataPairEmpty[D any](value entityDataPair[D]) bool {
	return value.entity == nil
}

func newEntityStore[D any](capacity uint32, freeCapacity uint32, initial D) entityStore[D] {
	return entityStore[D]{
		initial:  initial,
		data:     newSparseArray[entityDataPair[D]](capacity),
		dataSize: reflect.TypeOf(initial).Size(),
		free:     newStack[uint32](freeCapacity),
		getters:  make([]EntityDataGetter[D], MAX_DATA),
	}
}

func (ed *entityStore[D]) setID(id uint8) {
	ed.id = id
}

func (ed entityStore[D]) getID() uint8 {
	return ed.id
}

func (ed entityStore[D]) getDataSize() uintptr {
	return ed.dataSize
}

func (ed entityStore[D]) getEntityOffsets() []entityOffsetPair {
	pairs := make([]entityOffsetPair, 0, ed.data.count)
	for i := uint32(0); i < ed.data.count; i++ {
		data := ed.data.items[i]
		if data.entity != nil {
			pairs = append(pairs, entityOffsetPair{
				dataOffset: i,
				entity:     data.entity,
			})
		}
	}
	return pairs
}

func (ed *entityStore[D]) setInitial(initial any) {
	if i, ok := initial.(D); ok {
		ed.initial = i
	}
}

func (ed *entityStore[D]) shrink() {
	ed.data.shrink()
}

func (ed *entityStore[D]) add(e *Entity, link *entityDataLink) {
	if !ed.free.empty() {
		link.dataOffset = ed.free.pop()
	} else {
		link.dataOffset = ed.data.count
		ed.data.count++
	}
	item := &ed.data.items[link.dataOffset]
	item.data = ed.initial
	item.entity = e
	link.dataID = ed.id
}

func (ed *entityStore[D]) remove(e *Entity, link *entityDataLink) {
	if link.dataID != ed.id || ed.data.items[link.dataOffset].entity != e {
		return
	}
	ed.data.items[link.dataOffset].entity = nil
	ed.free.push(link.dataOffset)
	link.dataID = 0
}

func (ed entityStore[D]) getComponents() ComponentSet {
	return ed.components
}

func (ed *entityStore[D]) get(componentId uint8, link entityDataLink) any {
	data := &ed.data.items[link.dataOffset]
	getter := ed.getters[componentId]
	if getter != nil {
		return getter(&data.data)
	}
	return nil
}

func copy(dst any, src any) {
	d := reflect.ValueOf(dst)
	s := reflect.ValueOf(src)
	d.Elem().Set(s.Elem())
}

type stack[T any] struct {
	items []T
	count int
}

func newStack[T any](capacity uint32) stack[T] {
	return stack[T]{
		items: make([]T, capacity),
		count: 0,
	}
}

func (s stack[T]) empty() bool {
	return s.count == 0
}

func (s *stack[T]) push(value T) {
	if s.count == len(s.items) {
		s.items = append(s.items, value)
	} else {
		s.items[s.count] = value
	}
	s.count++
}

func (s *stack[T]) pop() T {
	var empty T
	if s.count == 0 {
		return empty
	}
	value := s.items[s.count]
	s.items[s.count] = empty
	s.count--
	return value
}

type sparseValue interface {
	isEmpty() bool
}

type sparseArray[T sparseValue] struct {
	items []T
	count uint32
}

func newSparseArray[T sparseValue](capacity uint32) sparseArray[T] {
	return sparseArray[T]{
		items: make([]T, capacity),
		count: 0,
	}
}

func (s *sparseArray[T]) take() *T {
	if s.count == uint32(len(s.items)) {
		var empty T
		s.items = append(s.items, empty)
	}
	value := &s.items[s.count]
	s.count++
	return value
}

func (s *sparseArray[T]) shrink() {
	for s.count > 0 && s.items[s.count-1].isEmpty() {
		s.count--
	}
}

func sliceRemoveAt[E any](slice []E, index int) []E {
	return append(slice[:index], slice[index+1:]...)
}
