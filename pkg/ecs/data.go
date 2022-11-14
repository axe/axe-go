package ecs

import (
	"reflect"

	"github.com/axe/axe-go/pkg/ds"
)

type DataSettings struct {
	Capacity      uint32
	StageCapacity uint32
}

type DataBase interface {
	ID() DataID
	Name() string
	Enable(settings DataSettings)
}

type Data[D any] struct {
	id     DataID
	name   string
	inital D
	values [DATA_MAX]dataValueData[D]
}

var _ DataBase = &Data[int]{}

var nextDataID DataID

func newData[D any](name string, initial D) *Data[D] {
	id := nextDataID
	nextDataID++

	data := &Data[D]{
		id:     id,
		name:   name,
		inital: initial,
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
			w.stagingComponents.Push(e)
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

func (d *Data[V]) Iterable() ds.Iterable[EntityData[*V]] {
	w := ActiveWorld()
	values := w.values[d.id].(*worldValues[V])
	return values.iterable
}

func (d *Data[V]) AddSystem(sys System[V]) {
	w := ActiveWorld()
	values := w.values[d.id].(*worldValues[V])
	values.systems = append(values.systems, sys)
}

func (d *Data[D]) Enable(settings DataSettings) {
	w := ActiveWorld()
	dataType := reflect.TypeOf(d.inital)

	data := &worldDatas[D]{
		data:     ds.NewSparseList[EntityData[D]](settings.Capacity),
		dataSize: dataType.Size(),
		initial:  d.inital,
		valueIDs: 0,
	}

	value := &worldValues[D]{
		iterables: make([]ds.Iterable[EntityData[*D]], 0),
		iterable:  ds.NewEmptyIterable[EntityData[*D]](),
		systems:   make([]System[D], 0),
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
}

var _ dataValueData[int] = &dataValue[int, int]{}

func (dv *dataValue[D, V]) addTo(w *World, data *worldDatas[D]) {
	values := w.values[dv.valueID].(*worldValues[V])
	dataValues := w.values[dv.dataId].(*worldValues[D])

	values.datas[dv.dataId] = &valueData[D, V]{dv, data, dataValues}
	values.dataIDs.Set(dv.dataId, true)

	data.valueIDs.Set(dv.valueID, true)

	values.iterables = append(values.iterables, ds.NewTranslateIterable[EntityData[*V], EntityData[D]](&data.data, func(source *EntityData[D]) *EntityData[*V] {
		value := dv.get(&source.Data)
		return &EntityData[*V]{
			Entity: source.Entity,
			Data:   value,
		}
	}))
	values.iterable = ds.NewMultiIterable(values.iterables)
}
