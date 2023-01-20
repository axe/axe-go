package ecs

func New() *Entity {
	return ActiveWorld().New()
}

type ID uint32

func (id ID) Entity() *Entity {
	index := uint32(id)
	world := ActiveWorld()
	return world.entity.At(index)
}

type Entity struct {
	id          ID
	arrangement ArrangementID
	offsets     []DataOffset
	staging     []valueStaging
}

func (e Entity) ID() ID {
	return e.id
}

func (e Entity) Live() bool {
	return len(e.offsets) > 0
}

func (e Entity) Staging() bool {
	return e.offsets != nil && len(e.offsets) == 0
}

func (e Entity) Deleted() bool {
	return e.offsets == nil
}

func (e Entity) Has(comp DataBase) bool {
	return e.Values().Get(comp.ID())
}

func (e Entity) HasLive(comp DataBase) bool {
	return e.LiveValues().Get(comp.ID())
}

func (e Entity) HasStaging(comp DataBase) bool {
	return e.StagingValues().Get(comp.ID())
}

func (e Entity) Values() DataIDs {
	return e.LiveValues() | e.StagingValues()
}

func (e Entity) StagingValues() DataIDs {
	values := DataIDs(0)
	if e.staging != nil {
		for _, stagingValue := range e.staging {
			values.Set(stagingValue.valueID, true)
		}
	}
	return values
}

func (e Entity) LiveValues() DataIDs {
	w := ActiveWorld()
	values, _ := e.getIDs(w)
	return values
}

func (e Entity) LiveDatas() DataIDs {
	w := ActiveWorld()
	_, datas := e.getIDs(w)
	return datas
}

func (e *Entity) Delete() {
	w := ActiveWorld()
	w.Delete(e)
}

func (e *Entity) locationFor(w *World, valueID DataID) dataLocation {
	if e.offsets == nil {
		return e.locationForStaging(w, valueID)
	}
	ids := e.getArrangement(w).pairs[valueID]
	if ids.live {
		return dataLocation{
			live:       true,
			dataID:     ids.dataID,
			valueID:    ids.valueID,
			dataOffset: e.offsets[ids.offsetIndex],
		}
	} else {
		return e.locationForStaging(w, valueID)
	}
}

func (e *Entity) locationForStaging(w *World, valueID DataID) dataLocation {
	if e.staging != nil {
		for _, value := range e.staging {
			if value.valueID == valueID {
				return dataLocation{
					live:       false,
					dataID:     value.valueID,
					valueID:    value.valueID,
					dataOffset: value.valueOffset,
				}
			}
		}
	}
	return dataLocation{valueID: valueID + 1}
}

func (e Entity) getArrangement(w *World) arrangement {
	return w.arrangements[e.arrangement]
}

func (e Entity) getIDs(w *World) (values DataIDs, datas DataIDs) {
	arr := e.getArrangement(w)
	values = arr.values
	datas = arr.datas
	return
}

type arrangement struct {
	id           ArrangementID
	pairs        []dataValuePair // pairs[valueID]
	datas        DataIDs
	values       DataIDs
	datasOrdered []DataID // datasOrdered[offsetIndex]
}

func (a arrangement) getDataPair(dataID DataID) *dataValuePair {
	for _, pair := range a.pairs {
		if pair.live && pair.dataID == dataID {
			return &pair
		}
	}
	return nil
}

func (a arrangement) getValuePair(valueID DataID) *dataValuePair {
	for _, pair := range a.pairs {
		if pair.live && pair.valueID == valueID {
			return &pair
		}
	}
	return nil
}
