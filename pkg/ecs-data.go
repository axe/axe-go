package axe

import (
	"reflect"

	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

type EntityDataSettings struct {
	Capacity             uint32
	StageCapacity        uint32
	ExcludeDefaultSystem bool
}

type EntityDataBase interface {
	ID() EntityDataID
	Name() string
	Enable(settings EntityDataSettings)
}

type EntityData[D any] struct {
	id            EntityDataID
	name          string
	initial       D
	values        [ENTITY_DATA_MAX]entityDataValueData[D]
	defaultSystem EntityDataSystem[D]
}

var _ EntityDataBase = &EntityData[int]{}

var nextDataID EntityDataID

func newEntityData[D any](name string, initial D) *EntityData[D] {
	id := nextDataID
	nextDataID++

	data := &EntityData[D]{
		id:      id,
		name:    name,
		initial: initial,
	}

	data.values[id] = &entityDataValue[D, D]{
		dataId:  id,
		valueID: id,
		get: func(data *D) *D {
			return data
		},
	}

	return data
}

func (d EntityData[V]) ID() EntityDataID {
	return d.id
}

func (d EntityData[V]) Name() string {
	return d.name
}

func (d *EntityData[V]) get(e *Entity, create bool) *V {
	w := ActiveWorld()
	location := e.locationFor(w, d.id)
	if location.valueID != d.id {
		if !create {
			return nil
		}
		if len(e.staging) == 0 && e.Live() {
			w.stagingComponents.Push(e.id)
		}
		values := w.values[d.id].(*worldValues[V])
		value := values.stage(e, d.id, w.ctx)
		return value
	}
	values := w.values[d.id].(*worldValues[V])
	data := values.datas[location.dataID]
	value := data.get(location.dataOffset, location.live)
	return value
}

func (d *EntityData[V]) SetSystem(sys EntityDataSystem[V]) *EntityData[V] {
	d.defaultSystem = sys
	return d
}

func (d *EntityData[V]) Get(e *Entity) *V {
	return d.get(e, false)
}

func (d *EntityData[V]) Add(e *Entity) *V {
	return d.get(e, true)
}

func (d *EntityData[V]) Set(e *Entity, value V) {
	ptr := d.get(e, true)
	*ptr = value
}

func (d *EntityData[V]) Iterable() ds.Iterable[EntityValue[*V]] {
	w := ActiveWorld()
	if w.values[d.id] == nil {
		return ds.NewEmptyIterable[EntityValue[*V]]()
	}
	values := w.values[d.id].(*worldValues[V])
	return values.iterable
}

func (d *EntityData[V]) AddSystem(sys EntityDataSystem[V]) {
	w := ActiveWorld()
	util.Assert(w.values[d.id] != nil, "a component must be enabled before adding a system")
	values := w.values[d.id].(*worldValues[V])
	values.systems = append(values.systems, sys)
}

func (d *EntityData[D]) Enable(settings EntityDataSettings) {
	w := ActiveWorld()
	dataType := reflect.TypeOf(d.initial)

	data := &worldDatas[D]{
		data:     ds.NewSparseList[EntityValue[D]](settings.Capacity),
		dataSize: dataType.Size(),
		initial:  d.initial,
		valueIDs: 0,
	}

	value := &worldValues[D]{
		iterables: make([]ds.Iterable[EntityValue[*D]], 0),
		iterable:  ds.NewEmptyIterable[EntityValue[*D]](),
		systems:   make([]EntityDataSystem[D], 0),
		staging:   ds.NewSparseList[D](settings.StageCapacity),
		dataIDs:   0,
	}

	w.typeMap[dataType] = d.id
	w.datas[d.id] = data
	w.values[d.id] = value
	w.datasBase[d.id] = d
	w.datasSorted = append(w.datasSorted, d.id)
	w.valuesSorted = append(w.valuesSorted, value)

	for _, value := range d.values {
		if value != nil {
			value.addTo(w, data)
		}
	}

	w.sortDatas()

	if !settings.ExcludeDefaultSystem && d.defaultSystem != nil {
		value.systems = append(value.systems, d.defaultSystem)
	}
}

var _ entityDataValueData[int] = &entityDataValue[int, int]{}

func (dv *entityDataValue[D, V]) addTo(w *World, data *worldDatas[D]) {
	values := w.values[dv.valueID].(*worldValues[V])
	dataValues := w.values[dv.dataId].(*worldValues[D])

	values.datas[dv.dataId] = &entityValueData[D, V]{dv, data, dataValues}
	values.dataIDs.Set(dv.dataId, true)

	data.valueIDs.Set(dv.valueID, true)

	values.iterables = append(values.iterables, ds.NewTranslateIterable[EntityValue[*V], EntityValue[D]](&data.data, func(source *EntityValue[D]) *EntityValue[*V] {
		value := dv.get(&source.Data)
		return &EntityValue[*V]{
			ID:   source.ID,
			Data: value,
		}
	}))
	values.iterable = ds.NewMultiIterable(values.iterables)
}
