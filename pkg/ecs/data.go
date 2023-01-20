package ecs

import (
	"reflect"

	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

type DataSettings struct {
	Capacity             uint32
	StageCapacity        uint32
	ExcludeDefaultSystem bool
}

type DataBase interface {
	ID() DataID
	Name() string
	Enable(settings DataSettings)
}

type Data[D any] struct {
	id            DataID
	name          string
	initial       D
	values        [DATA_MAX]dataValueData[D]
	defaultSystem DataSystem[D]
}

var _ DataBase = &Data[int]{}

var nextDataID DataID

func newEntityData[D any](name string, initial D) *Data[D] {
	id := nextDataID
	nextDataID++

	data := &Data[D]{
		id:      id,
		name:    name,
		initial: initial,
	}

	data.values[id] = &dataValue[D, D]{
		dataId:  id,
		valueID: id,
		get: func(data *D) *D {
			return data
		},
	}

	return data
}

func (d Data[V]) ID() DataID {
	return d.id
}

func (d Data[V]) Name() string {
	return d.name
}

func (d *Data[V]) get(e *Entity, create bool) *V {
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

func (d *Data[V]) SetSystem(sys DataSystem[V]) *Data[V] {
	d.defaultSystem = sys
	return d
}

func (d *Data[V]) Get(e *Entity) *V {
	return d.get(e, false)
}

func (d *Data[V]) Add(e *Entity) *V {
	return d.get(e, true)
}

func (d *Data[V]) Set(e *Entity, value V) {
	ptr := d.get(e, true)
	*ptr = value
}

func (d *Data[V]) Iterable() ds.Iterable[Value[*V]] {
	w := ActiveWorld()
	if w.values[d.id] == nil {
		return ds.NewEmptyIterable[Value[*V]]()
	}
	values := w.values[d.id].(*worldValues[V])
	return values.iterable
}

func (d *Data[V]) AddSystem(sys DataSystem[V]) {
	w := ActiveWorld()
	util.Assert(w.values[d.id] != nil, "a component must be enabled before adding a system")
	values := w.values[d.id].(*worldValues[V])
	values.systems = append(values.systems, sys)
}

func (d *Data[D]) Enable(settings DataSettings) {
	w := ActiveWorld()
	dataType := reflect.TypeOf(d.initial)

	data := &worldDatas[D]{
		data:     ds.NewSparseList[Value[D]](settings.Capacity),
		dataSize: dataType.Size(),
		initial:  d.initial,
		valueIDs: 0,
	}

	value := &worldValues[D]{
		iterables: make([]ds.Iterable[Value[*D]], 0),
		iterable:  ds.NewEmptyIterable[Value[*D]](),
		systems:   make([]DataSystem[D], 0),
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

var _ dataValueData[int] = &dataValue[int, int]{}

func (dv *dataValue[D, V]) addTo(w *World, data *worldDatas[D]) {
	values := w.values[dv.valueID].(*worldValues[V])
	dataValues := w.values[dv.dataId].(*worldValues[D])

	values.datas[dv.dataId] = &valueData[D, V]{dv, data, dataValues}
	values.dataIDs.Set(dv.dataId, true)

	data.valueIDs.Set(dv.valueID, true)

	values.iterables = append(values.iterables, ds.NewTranslateIterable[Value[*V], Value[D]](&data.data, func(source *Value[D]) *Value[*V] {
		value := dv.get(&source.Data)
		return &Value[*V]{
			ID:   source.ID,
			Data: value,
		}
	}))
	values.iterable = ds.NewMultiIterable(values.iterables)
}
