package ecs

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
