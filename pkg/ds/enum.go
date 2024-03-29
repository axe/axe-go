package ds

import "golang.org/x/exp/constraints"

type EnumMap[K constraints.Unsigned, V any] []V

func NewEnumMap[K constraints.Unsigned, V any](m map[K]V) EnumMap[K, V] {
	em := make(EnumMap[K, V], len(m))
	for k, v := range m {
		em.Set(k, v)
	}
	return em
}

func (m EnumMap[K, V]) Len() int {
	return len(m)
}

func (m EnumMap[K, V]) Empty() bool {
	return len(m) == 0
}

func (m *EnumMap[K, V]) Set(key K, value V) {
	outside := int(key) - len(*m) + 1
	if outside > 0 {
		*m = append(*m, make([]V, outside)...)
	}
	(*m)[key] = value
}

func (m EnumMap[K, V]) Get(key K) (value V) {
	if m != nil && int(key) < len(m) {
		value = m[key]
	}
	return
}

func (m EnumMap[K, V]) Has(key K) bool {
	if m != nil && int(key) < len(m) {
		return true
	}
	return false
}

func (m EnumMap[K, V]) Remove(key K) {
	if m != nil && int(key) < len(m) {
		var zero V
		m[key] = zero
	}
}

func (m *EnumMap[K, V]) Merge(other EnumMap[K, V], replace bool, isMissing func(V) bool) {
	for index, value := range other {
		if !isMissing(value) {
			key := K(index)
			if replace || !m.Has(key) || isMissing((*m)[key]) {
				m.Set(key, value)
			}
		}
	}
}
