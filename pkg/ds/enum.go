package ds

type EnumMap[K ~int, V any] []V

func NewEnumMap[K ~int, V any](m map[K]V) EnumMap[K, V] {
	em := make(EnumMap[K, V], len(m))
	for k, v := range m {
		em.Set(k, v)
	}
	return em
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

func (m EnumMap[K, V]) Remove(key K) {
	if m != nil && int(key) < len(m) {
		var zero V
		m[key] = zero
	}
}
