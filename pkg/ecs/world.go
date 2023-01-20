package ecs

import (
	"reflect"
	"sort"

	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

type WorldSettings struct {
	EntityCapacity            uint32
	EntityStageCapacity       uint32
	AverageComponentPerEntity uint32
	DeleteOnDestroy           bool
}

func (ws WorldSettings) IsZero() bool {
	return ws.EntityCapacity == 0 && ws.EntityStageCapacity == 0 && ws.AverageComponentPerEntity == 0 && !ws.DeleteOnDestroy
}

type WorldSearch struct {
	Match               util.Match[DataIDs]
	IncludeStagedValues bool
	IncludeStaged       bool
}

func (search WorldSearch) IsMatch(e *Entity) bool {
	if !e.Live() && !search.IncludeStaged {
		return false
	}
	if search.Match == nil {
		return true
	}
	values := e.LiveValues()
	if search.IncludeStagedValues {
		values |= e.StagingValues()
	}
	return search.Match(values)
}

type World struct {
	Settings WorldSettings
	Name     string

	ctx               Context
	datas             [DATA_MAX]worldDataStore // []worldDatas[D]
	datasSorted       []DataID
	datasBase         [DATA_MAX]DataBase
	values            [DATA_MAX]worldValueStore // []worldValues[V]
	valuesSorted      []worldValueStore
	arrangements      []arrangement
	arrangementMap    map[DataIDs]int
	entity            ds.SparseList[Entity]
	staging           ds.Stack[ID]
	stagingComponents ds.Stack[ID]
	typeMap           map[reflect.Type]DataID
	systems           []System
}

// var _ GameSystem = &World{}

func NewWorld(name string, settings WorldSettings) *World {
	world := &World{
		Name:     name,
		Settings: settings,

		datasSorted:       make([]DataID, 0),
		valuesSorted:      make([]worldValueStore, 0),
		arrangements:      make([]arrangement, 0),
		arrangementMap:    make(map[DataIDs]int),
		entity:            ds.NewSparseList[Entity](settings.EntityCapacity),
		staging:           ds.NewStack[ID](settings.EntityStageCapacity),
		stagingComponents: ds.NewStack[ID](settings.EntityStageCapacity),
		typeMap:           make(map[reflect.Type]DataID),
		systems:           make([]System, 0),
	}

	world.ctx.World = world
	activeWorld = world
	return world
}

var activeWorld *World

func ActiveWorld() *World {
	util.Assert(activeWorld != nil, "There is no active world, you must create one or activate one")
	return activeWorld
}

func HasActiveWorld() bool {
	return activeWorld != nil
}

func (w *World) IsActive() bool {
	return w == activeWorld
}

func (w *World) Activate() {
	activeWorld = w
}

func (w *World) getArrangement(values DataIDs) *arrangement {
	if arrangement, ok := w.arrangementMap[values]; ok {
		return &w.arrangements[arrangement]
	}
	id := ArrangementID(len(w.arrangements))
	pairs := make([]dataValuePair, values.LastOn()+1)
	remaining := values
	offsetIndex := uint8(0)
	datasChosen := DataIDs(0)
	datasOrdered := make([]DataID, 0)
	for _, dataID := range w.datasSorted {
		data := w.datas[dataID]
		if data == nil {
			continue
		}
		dataValues := data.getValueIDs()
		if remaining.All(dataValues) {
			remaining.Flip(dataValues)
			datasChosen.Set(dataID, true)

			for dataValues != 0 {
				valueID := dataValues.TakeFirst()
				pairs[valueID] = dataValuePair{
					live:        true,
					dataID:      dataID,
					valueID:     DataID(valueID),
					offsetIndex: offsetIndex,
				}
			}
			datasOrdered = append(datasOrdered, dataID)
			offsetIndex++
			if remaining == 0 {
				break
			}
		}
	}
	arrangementIndex := len(w.arrangements)
	w.arrangements = append(w.arrangements, arrangement{
		id:           id,
		pairs:        pairs,
		datas:        datasChosen,
		datasOrdered: datasOrdered,
		values:       values,
	})
	arrangement := &w.arrangements[id]
	w.arrangementMap[values] = arrangementIndex
	return arrangement
}

func (w *World) sortDatas() {
	sort.Slice(w.datasSorted, func(i, j int) bool {
		a := w.datas[w.datasSorted[i]]
		b := w.datas[w.datasSorted[j]]
		if a.getValueIDs().Ons() > b.getValueIDs().Ons() {
			return true
		}
		if a.getDataSize() > b.getDataSize() {
			return true
		}
		return false
	})
}

func (w *World) Enable(settings DataSettings, datas ...DataBase) {
	w.Activate()

	for _, data := range datas {
		data.Enable(settings)
	}
}

func (w *World) New() *Entity {
	w.Activate()

	e, id := w.entity.Take()
	e.id = ID(id)
	e.offsets = nil
	e.staging = make([]valueStaging, 0, w.Settings.AverageComponentPerEntity)
	w.staging.Push(e.id)
	for _, sys := range w.systems {
		sys.OnStage(e, w.ctx)
	}
	return e
}

func (w *World) Delete(e *Entity) {
	if e.Deleted() {
		return
	}
	w.Activate()

	for _, sys := range w.systems {
		sys.OnDelete(e, w.ctx)
	}

	if e.offsets != nil {
		arrangement := e.getArrangement(w)
		for valueID, dataValue := range arrangement.pairs {
			if !dataValue.live {
				continue
			}
			w.values[valueID].remove(e, dataValue.dataID, e.offsets[dataValue.offsetIndex], true, w.ctx)
		}
		for offsetIndex, dataID := range arrangement.datasOrdered {
			w.datas[dataID].remove(e, e.offsets[offsetIndex])
		}
	}

	if e.staging != nil {
		for _, stagingValue := range e.staging {
			w.values[stagingValue.valueID].remove(e, stagingValue.valueID, stagingValue.valueOffset, false, w.ctx)
		}
	}

	e.offsets = nil
	e.staging = nil
	w.entity.Free(uint32(e.id))
}

func (w *World) Stage() {
	w.Activate()

	w.stageEntity()
	w.stageValuesRestructure()
}

func (w *World) stageEntity() {
	for !w.staging.IsEmpty() {
		e := w.staging.Pop().Entity()

		stagingValues := e.StagingValues()

		arrangement := w.getArrangement(stagingValues)

		e.arrangement = arrangement.id
		e.offsets = make([]DataOffset, len(arrangement.datasOrdered))

		for indexOffset, dataID := range arrangement.datasOrdered {
			dataOffset := w.datas[dataID].add(e)
			e.offsets[indexOffset] = dataOffset
		}

		w.stageArrangementValues(e, arrangement)

		e.staging = nil
	}
}

func (w *World) stageValuesRestructure() {
	for !w.stagingComponents.IsEmpty() {
		e := w.stagingComponents.Pop().Entity()

		stagingValues := e.StagingValues()
		existingArrangement := e.getArrangement(w)
		removeDatas := existingArrangement.datas
		movingValues := existingArrangement.values

		arrangement := w.getArrangement(stagingValues | existingArrangement.values)

		offsets := make([]DataOffset, len(arrangement.datasOrdered))

		// for the new arrangement add new data or get offsets for unchanged
		for indexOffset, dataID := range arrangement.datasOrdered {
			if existingArrangement.datas.Get(dataID) {
				dataPair := existingArrangement.getDataPair(dataID)
				offsets[indexOffset] = e.offsets[dataPair.offsetIndex]
				dataValues := w.datas[dataID].getValueIDs()
				movingValues.Flip(dataValues)
			} else {
				dataOffset := w.datas[dataID].add(e)
				offsets[indexOffset] = dataOffset
			}
			removeDatas.Set(dataID, false)
		}

		// stage new values
		w.stageArrangementValues(e, arrangement)

		// handle existing values moving between data
		for !movingValues.IsEmpty() {
			valueID := DataID(movingValues.TakeFirst())
			sourcePair := existingArrangement.getValuePair(valueID)
			targetPair := arrangement.getValuePair(valueID)
			if sourcePair.dataID != targetPair.dataID {
				w.values[valueID].move(sourcePair.dataID, e.offsets[sourcePair.offsetIndex], targetPair.dataID, offsets[targetPair.offsetIndex])
			}
		}

		// remove any data that is no longer needed
		for !removeDatas.IsEmpty() {
			dataID := DataID(removeDatas.TakeFirst())
			dataPair := existingArrangement.getDataPair(dataID)
			offset := e.offsets[dataPair.offsetIndex]
			w.datas[dataID].remove(e, offset)
		}

		e.arrangement = arrangement.id
		e.offsets = offsets
		e.staging = nil
	}
}

func (w *World) stageArrangementValues(e *Entity, a *arrangement) {
	for _, stagingValue := range e.staging {
		data := a.pairs[stagingValue.valueID]
		stageValues := w.values[stagingValue.valueID]
		stageOffset := stagingValue.valueOffset
		targetID := data.dataID
		targetOffset := e.offsets[data.offsetIndex]
		stageValues.unstage(e, stageOffset, targetID, targetOffset, w.ctx)
	}
}

func (w *World) Init( /*game *Game*/ ) error {
	w.Activate()
	// w.ctx.Game = game

	for _, sys := range w.systems {
		err := sys.Init(w.ctx)
		if err != nil {
			return err
		}
	}
	for _, values := range w.valuesSorted {
		err := values.init(w.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *World) Update( /*game *Game*/ ) {
	w.Stage()

	for _, sys := range w.systems {
		sys.Update(&w.entity, w.ctx)
	}

	for _, values := range w.valuesSorted {
		values.update(w.ctx)
	}
}

func (w *World) Destroy() {
	w.Activate()

	if w.Settings.DeleteOnDestroy {
		entities := w.entity.Iterator()
		for entities.HasNext() {
			e := entities.Next()
			w.Delete(e)
		}
	}
	for _, sys := range w.systems {
		sys.Destroy(w.ctx)
	}
	w.systems = w.systems[:0]
	for _, values := range w.valuesSorted {
		values.destroy(w.ctx)
	}
	for _, dataID := range w.datasSorted {
		w.datas[dataID].destroy()
	}
	w.entity.Clear()

	if w.IsActive() {
		activeWorld = nil
	}
}

func (w *World) AddSystem(sys System) {
	w.systems = append(w.systems, sys)
}

func (w *World) Iterable() ds.Iterable[Entity] {
	w.Activate()

	return &w.entity
}

func (w *World) Search(search WorldSearch) ds.Iterable[Entity] {
	w.Activate()

	return ds.NewFilterIterable[Entity](&w.entity, search.IsMatch)
}

type worldValues[V any] struct {
	iterables []ds.Iterable[Value[*V]]
	iterable  ds.Iterable[Value[*V]]
	systems   []DataSystem[V]
	staging   ds.SparseList[V]
	datas     [DATA_MAX]worldValueData[V]
	dataIDs   DataIDs
}

func (wv *worldValues[V]) init(ctx Context) error {
	for _, sys := range wv.systems {
		err := sys.Init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (wv *worldValues[V]) update(ctx Context) {
	for _, sys := range wv.systems {
		sys.Update(wv.iterable, ctx)
	}
}

func (wv *worldValues[V]) destroy(ctx Context) {
	for _, sys := range wv.systems {
		sys.Destroy(ctx)
	}
	wv.systems = wv.systems[:0]
	wv.staging.Clear()
}

func (wv *worldValues[V]) stage(e *Entity, valueID DataID, ctx Context) *V {
	value, offset := wv.staging.Take()

	e.staging = append(e.staging, valueStaging{
		valueID:     valueID,
		valueOffset: DataOffset(offset),
	})
	for _, sys := range wv.systems {
		sys.OnStage(value, e, ctx)
	}
	return value
}

func (wv *worldValues[V]) unstage(e *Entity, stageOffset DataOffset, target DataID, targetOffset DataOffset, ctx Context) {
	index := uint32(stageOffset)
	stageValue := wv.staging.At(index)
	liveValue := wv.datas[target].get(targetOffset, true)
	*liveValue = *stageValue
	wv.staging.Free(index)

	for _, sys := range wv.systems {
		sys.OnLive(liveValue, e, ctx)
	}
}

func (wv *worldValues[V]) remove(e *Entity, data DataID, dataOffset DataOffset, live bool, ctx Context) {
	value := wv.datas[data].get(dataOffset, live)

	for _, sys := range wv.systems {
		sys.OnRemove(value, e, ctx)
	}

	wv.datas[data].free(dataOffset, live)
}

func (wv *worldValues[V]) move(sourceID DataID, sourceOffset DataOffset, targetID DataID, targetOffset DataOffset) {
	value := wv.datas[sourceID].get(sourceOffset, true)
	target := wv.datas[targetID].get(targetOffset, true)
	*target = *value
	wv.datas[sourceID].free(sourceOffset, true)
}

type valueData[D any, V any] struct {
	dataValue *dataValue[D, V]
	data      *worldDatas[D]
	value     *worldValues[D]
}

var _ worldValueData[int] = &valueData[int, int]{}

func (vd *valueData[D, V]) get(offset DataOffset, live bool) *V {
	var value *V
	if live {
		entityData := vd.data.data.At(uint32(offset))
		value = vd.dataValue.get(&entityData.Data)
	} else {
		data := vd.value.staging.At(uint32(offset))
		value = vd.dataValue.get(data)
	}
	return value
}

func (vd *valueData[D, V]) free(offset DataOffset, live bool) {
	if live {
		vd.data.data.Free(uint32(offset))
	} else {
		vd.value.staging.Free(uint32(offset))
	}
}

type worldDatas[D any] struct {
	data     ds.SparseList[Value[D]]
	initial  D
	valueIDs DataIDs
	dataSize uintptr
}

func (datas worldDatas[D]) getValueIDs() DataIDs {
	return datas.valueIDs
}

func (datas worldDatas[D]) getDataSize() uintptr {
	return datas.dataSize
}

func (datas *worldDatas[D]) add(e *Entity) DataOffset {
	data, offset := datas.data.Take()
	data.ID = e.id
	data.Data = datas.initial
	return DataOffset(offset)
}

func (datas *worldDatas[D]) remove(e *Entity, offset DataOffset) {
	dataOffset := uint32(offset)
	data := datas.data.At(dataOffset)
	data.ID = 0
	datas.data.Free(dataOffset)
}

func (datas *worldDatas[D]) destroy() {
	datas.data.Clear()
}
