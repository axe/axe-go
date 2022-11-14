package axe

func NewEntity() *Entity {
	return ActiveWorld().New()
}

type Entity struct {
	id          uint32
	arrangement EntityArrangementID
	offsets     []EntityDataOffset
	staging     []entityValueStaging
}

func (e Entity) ID() uint32 {
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

func (e Entity) Has(comp EntityDataBase) bool {
	return e.Values().Get(comp.ID())
}

func (e Entity) HasLive(comp EntityDataBase) bool {
	return e.LiveValues().Get(comp.ID())
}

func (e Entity) HasStaging(comp EntityDataBase) bool {
	return e.StagingValues().Get(comp.ID())
}

func (e Entity) Values() EntityDataIDs {
	return e.LiveValues() | e.StagingValues()
}

func (e Entity) StagingValues() EntityDataIDs {
	values := EntityDataIDs(0)
	if e.staging != nil {
		for _, stagingValue := range e.staging {
			values.Set(stagingValue.valueID, true)
		}
	}
	return values
}

func (e Entity) LiveValues() EntityDataIDs {
	w := ActiveWorld()
	values, _ := e.getIDs(w)
	return values
}

func (e Entity) LiveDatas() EntityDataIDs {
	w := ActiveWorld()
	_, datas := e.getIDs(w)
	return datas
}

func (e *Entity) Delete() {
	w := ActiveWorld()
	w.Delete(e)
}

func (e *Entity) locationFor(w *World, valueID EntityDataID) entityDataLocation {
	if e.offsets == nil {
		return e.locationForStaging(w, valueID)
	}
	ids := e.getArrangement(w).pairs[valueID]
	if ids.live {
		return entityDataLocation{
			live:       true,
			dataID:     ids.dataID,
			valueID:    ids.valueID,
			dataOffset: e.offsets[ids.offsetIndex],
		}
	} else {
		return e.locationForStaging(w, valueID)
	}
}

func (e *Entity) locationForStaging(w *World, valueID EntityDataID) entityDataLocation {
	if e.staging != nil {
		for _, value := range e.staging {
			if value.valueID == valueID {
				return entityDataLocation{
					live:       false,
					dataID:     value.valueID,
					valueID:    value.valueID,
					dataOffset: value.valueOffset,
				}
			}
		}
	}
	return entityDataLocation{valueID: valueID + 1}
}

func (e Entity) getArrangement(w *World) arrangement {
	return w.arrangements[e.arrangement]
}

func (e Entity) getIDs(w *World) (values EntityDataIDs, datas EntityDataIDs) {
	arr := e.getArrangement(w)
	values = arr.values
	datas = arr.datas
	return
}

type arrangement struct {
	id           EntityArrangementID
	pairs        []entityDataValuePair // pairs[valueID]
	datas        EntityDataIDs
	values       EntityDataIDs
	datasOrdered []EntityDataID // datasOrdered[offsetIndex]
}

func (a arrangement) getDataPair(dataID EntityDataID) *entityDataValuePair {
	for _, pair := range a.pairs {
		if pair.live && pair.dataID == dataID {
			return &pair
		}
	}
	return nil
}

func (a arrangement) getValuePair(valueID EntityDataID) *entityDataValuePair {
	for _, pair := range a.pairs {
		if pair.live && pair.valueID == valueID {
			return &pair
		}
	}
	return nil
}
