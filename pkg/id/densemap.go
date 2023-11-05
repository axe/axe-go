package id

import "golang.org/x/exp/constraints"

type DenseMap[V any, A constraints.Unsigned, L constraints.Unsigned] struct {
	area   *Area[uint32, A]
	local  Area[A, L]
	values []V
}

func NewDenseMap[V any, A constraints.Unsigned, L constraints.Unsigned](opts ...Option) DenseMap[V, A, L] {
	sm := DenseMap[V, A, L]{
		values: make([]V, 0),
	}
	processOptions(&sm, opts)
	return sm
}

func (m *DenseMap[V, A, L]) setCapacity(capacity int) {
	m.values = make([]V, 0, capacity)
	m.local.setCapacity(capacity)
}

func (m *DenseMap[V, A, L]) setArea(area *Area[uint32, A]) {
	m.area = area
}

func (m *DenseMap[V, A, L]) indexOf(key Identifier) int {
	if !key.Exists() {
		return -1
	}
	areaID := m.area.Peek(uint32(key))
	if areaID < 0 {
		return -1
	}
	return m.local.Peek(A(areaID))
}

func (m *DenseMap[V, A, L]) SetKeyValues(kv []KeyValue[V]) {
	setKeyValues[V](m, kv)
}

func (s *DenseMap[V, A, L]) SetMap(m map[Identifier]V) {
	setMap[V](s, m)
}

func (s *DenseMap[V, A, L]) SetStringMap(m map[string]V) {
	setStringMap[V](s, m)
}

func (m DenseMap[V, A, L]) Values() []V {
	return m.values
}

func (m *DenseMap[V, A, L]) Set(key Identifier, value V) {
	*m.Take(key) = value
}

func (m *DenseMap[V, A, L]) Get(key Identifier) V {
	i := m.indexOf(key)
	if i == -1 {
		var empty V
		return empty
	}
	return m.values[i]
}

func (m *DenseMap[V, A, L]) Ptr(key Identifier) *V {
	i := m.indexOf(key)
	if i == -1 {
		return nil
	}
	return &m.values[i]
}

func (m *DenseMap[V, A, L]) Has(key Identifier) bool {
	i := m.indexOf(key)
	return i >= 0 && i < len(m.values)
}

func (m *DenseMap[V, A, L]) Take(key Identifier) *V {
	if !key.Exists() {
		return nil
	}
	areaID := m.area.Translate(uint32(key))
	index := m.local.Translate(areaID)
	if int(index) == len(m.values) {
		var empty V
		m.values = append(m.values, empty)
	}
	return &m.values[index]
}

func (m DenseMap[V, A, L]) Empty() bool {
	return len(m.values) == 0
}

func (m DenseMap[V, A, L]) Len() int {
	return len(m.values)
}

func (m *DenseMap[V, A, L]) Clear() {
	m.values = make([]V, 0, cap(m.values))
	m.local.Clear()
}

func (m *DenseMap[V, A, L]) Remove(key Identifier, maintainOrder bool) bool {
	if !key.Exists() {
		return false
	}
	areaID := m.area.Peek(uint32(key))
	if areaID == -1 {
		return false
	}
	index := m.local.Remove(A(areaID), maintainOrder)
	if index == -1 {
		return false
	}
	if maintainOrder {
		m.values = removeAt(m.values, index)
	} else {
		m.values = moveEndTo(m.values, index)
	}
	return true
}
