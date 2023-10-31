package types

import (
	"strings"

	"github.com/axe/axe-go/pkg/util"
	"golang.org/x/exp/constraints"
)

// Auto-incrementing value generator.
type Incrementor[V constraints.Ordered] struct {
	current   V
	increment V
}

// Returns a new incrementor with the given initial value and how much to increment by on each Get() call.
func NewIncrementor[V constraints.Ordered](start V, increment V) Incrementor[V] {
	return Incrementor[V]{current: start, increment: increment}
}

// Returns an auto-incrementing value.
func (i Incrementor[V]) Get() V {
	c := i.current
	i.current = i.current + i.increment
	return c
}

// A map of named values. This structure is useful when the most common
// operation is iteration where adding & removing are less common.
type NameMap[V any] struct {
	Values []V

	getName     func(V) string
	indices     map[string]int
	insensitive bool
	ordered     bool
}

// Creates a new named map optionally having the names be case insensitive and
// optionally retaining insertion order when remove is performed.
func NewNameMap[V any](getName func(V) string, capacity int, insensitive bool, ordered bool) NameMap[V] {
	return NameMap[V]{
		Values: make([]V, 0, capacity),

		getName:     getName,
		indices:     make(map[string]int, capacity),
		insensitive: insensitive,
		ordered:     ordered,
	}
}

// Adds a named value to the map and returns true on success. False is returned
// if a value with the given name already exists.
func (m *NameMap[V]) Add(value V) bool {
	key := m.keyOf(value)
	_, exists := m.indices[key]
	if !exists {
		i := len(m.Values)
		m.Values = append(m.Values, value)
		m.indices[key] = i
	}
	return !exists
}

// Sets the named value to the map whether it exists or not. If it exists
// already the existing value is replaced.
func (m *NameMap[V]) Set(value V) {
	key := m.keyOf(value)
	if existingIndex, exists := m.indices[key]; exists {
		m.Values[existingIndex] = value
	} else {
		i := len(m.Values)
		m.Values = append(m.Values, value)
		m.indices[key] = i
	}
}

// Removes the value from the map.
func (m *NameMap[V]) Remove(value V) bool {
	return m.RemoveName(m.getName(value))
}

// Removes the value with the name from the map.
func (m *NameMap[V]) RemoveName(name string) bool {
	key := m.toKey(name)
	index, exists := m.indices[key]
	if exists {
		last := len(m.Values) - 1
		notLast := index < last
		if m.ordered && !notLast {
			m.Values = util.SliceRemoveAt(m.Values, index)
			delete(m.indices, key)
			for mapKey, mapIndex := range m.indices {
				if mapIndex > index {
					m.indices[mapKey] = mapIndex - 1
				}
			}
		} else {
			if !notLast {
				lastKey := m.keyOf(m.Values[last])
				m.indices[lastKey] = index
				m.Values[index] = m.Values[last]
			}
			m.Values = m.Values[:last]
			delete(m.indices, key)
		}
	}
	return exists
}

// Gets a value with the name from the map and whether it exists at all.
func (m NameMap[V]) Get(name string) (V, bool) {
	key := m.toKey(name)
	if index, exists := m.indices[key]; exists {
		return m.Values[index], true
	}
	var empty V
	return empty, false
}

// Clears the map of all values.
func (m *NameMap[V]) Clear() {
	cap := cap(m.Values)
	m.Values = make([]V, 0, cap)
	m.indices = make(map[string]int, cap)
}

// Returns the number of items in the map.
func (m NameMap[V]) Len() int {
	return len(m.Values)
}

// Handles the value being renamed given the old name. If the old name
// does not exist a rebuild is performed.
func (m *NameMap[V]) Rename(oldName string) {
	oldKey := m.toKey(oldName)
	if index, exists := m.indices[oldKey]; exists {
		value := m.Values[index]
		newKey := m.keyOf(value)
		delete(m.indices, oldKey)
		m.indices[newKey] = index
	} else {
		m.Rebuild()
	}
}

// Handles the value being renamed given the old name. If the old name
// does not exist a rebuild is performed.
func (m *NameMap[V]) Rebuild() {
	m.indices = make(map[string]int, len(m.indices))
	for i, v := range m.Values {
		key := m.keyOf(v)
		m.indices[key] = i
	}
}

// Returns whether the internal state of the map is valid. The internal state of the
// map can be invalid if the Values was manipulated directly or if any
// of the names of the values has changed and a Rebuild or Update was not called.
func (m NameMap[V]) IsValid() bool {
	if len(m.Values) != len(m.indices) {
		return false
	}
	for i, v := range m.Values {
		k := m.keyOf(v)
		if m.indices[k] != i {
			return false
		}
	}
	return true
}

// Returns whether this map has a value with the given name.
func (m NameMap[V]) Has(name string) bool {
	_, exists := m.indices[m.toKey(name)]
	return exists
}

// Converts the name to a key used in the index map.
func (m NameMap[V]) toKey(name string) string {
	if m.insensitive {
		return strings.ToLower(name)
	} else {
		return name
	}
}

// Converts the value to a key used in the index map.
func (m NameMap[V]) keyOf(value V) string {
	return m.toKey(m.getName(value))
}
