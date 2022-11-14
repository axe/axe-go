package ecs

import "github.com/axe/axe-go/pkg/ds"

type worldValues[V any] struct {
	iterables []ds.Iterable[EntityData[*V]]
	iterable  ds.Iterable[EntityData[*V]]
	systems   []System[V]
	staging   ds.SparseList[V]
	datas     [DATA_MAX]worldValueData[V]
	dataIDs   DataIDs
}

func (wv *worldValues[V]) init(ctx SystemContext) error {
	for _, sys := range wv.systems {
		err := sys.Init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (wv *worldValues[V]) update(ctx SystemContext) {
	for _, sys := range wv.systems {
		sys.Update(wv.iterable, ctx)
	}
}

func (wv *worldValues[V]) destroy(ctx SystemContext) {
	for _, sys := range wv.systems {
		sys.Destroy(ctx)
	}
	wv.systems = wv.systems[:0]
	wv.staging.Clear()
}

func (wv *worldValues[V]) stage(e *Entity, valueID DataID, ctx SystemContext) *V {
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

func (wv *worldValues[V]) unstage(e *Entity, stageOffset DataOffset, target DataID, targetOffset DataOffset, ctx SystemContext) {
	index := uint32(stageOffset)
	stageValue := wv.staging.At(index)
	liveValue := wv.datas[target].get(targetOffset, true)
	*liveValue = *stageValue
	wv.staging.Free(index)

	for _, sys := range wv.systems {
		sys.OnLive(liveValue, e, ctx)
	}
}

func (wv *worldValues[V]) remove(e *Entity, data DataID, dataOffset DataOffset, live bool, ctx SystemContext) {
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
	data     ds.SparseList[EntityData[D]]
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
	data.Entity = e
	data.Data = datas.initial
	return DataOffset(offset)
}

func (datas *worldDatas[D]) remove(e *Entity, offset DataOffset) {
	dataOffset := uint32(offset)
	data := datas.data.At(dataOffset)
	data.Entity = nil
	datas.data.Free(dataOffset)
}

func (datas *worldDatas[D]) destroy() {
	datas.data.Clear()
}
