package id

import "golang.org/x/exp/constraints"

type SparseMap[V any, SID constraints.Unsigned] struct {
	area         *Area[uint32, SID]
	values       []V
	resizeBuffer int
}

var _ HasCapacity = &SparseMap[int, uint]{}
var _ HasResizeBuffer = &SparseMap[int, uint]{}

func NewSparseMap[V any, SID constraints.Unsigned](opts ...Option) SparseMap[V, SID] {
	sm := SparseMap[V, SID]{
		values: make([]V, 0),
	}
	processOptions(&sm, opts)
	return sm
}

func (m *SparseMap[V, SID]) setResizeBuffer(resizeBuffer int) {
	m.resizeBuffer = resizeBuffer
}

func (m *SparseMap[V, SID]) setCapacity(capacity int) {
	m.values = make([]V, 0, capacity)
}

func (m *SparseMap[V, SID]) setArea(area *Area[uint32, SID]) {
	m.area = area
}

func (m *SparseMap[V, SID]) SetKeyValues(kv []KeyValue[V]) {
	setKeyValues[V](m, kv)
}

func (s *SparseMap[V, SID]) SetMap(m map[Identifier]V) {
	setMap[V](s, m)
}

func (s *SparseMap[V, SID]) SetStringMap(m map[string]V) {
	setStringMap[V](s, m)
}

func (m SparseMap[V, SID]) Values() []V {
	return m.values
}

func (m *SparseMap[V, SID]) Set(key Identifier, value V) {
	*m.Take(key) = value
}

func (m *SparseMap[V, SID]) Get(key Identifier) V {
	p := m.Ptr(key)
	if p == nil {
		var empty V
		return empty
	}
	return *p
}

func (m *SparseMap[V, SID]) Ptr(key Identifier) *V {
	if key.Exists() {
		mapID := m.area.Peek(uint32(key))
		if mapID >= 0 && mapID < len(m.values) {
			return &m.values[mapID]
		}
	}
	return nil
}

func (m *SparseMap[V, SID]) Has(key Identifier) bool {
	if key.Exists() {
		mapID := m.area.Peek(uint32(key))
		if mapID >= 0 && mapID < len(m.values) {
			return true
		}
	}
	return false
}

func (m *SparseMap[V, SID]) Take(key Identifier) *V {
	if !key.Exists() {
		return nil
	}
	mapID := int(m.area.Translate(uint32(key)))
	if mapID >= len(m.values) {
		nextSize := mapID + m.resizeBuffer + 1
		m.values = resize(m.values, nextSize)
	}
	return &m.values[mapID]
}

func (m SparseMap[V, SID]) Empty() bool {
	return len(m.values) == 0
}

func (m SparseMap[V, SID]) Len() int {
	return len(m.values)
}
