package ds

type Set[V comparable] map[V]struct{}

var _ Sized = Set[int]{}
var _ Clearable = &Set[int]{}
var _ Linear[int] = Set[int]{}

func (s Set[V]) Add(v V) {
	s[v] = struct{}{}
}

func (s Set[V]) Has(v V) bool {
	_, exists := s[v]
	return exists
}

func (s Set[V]) Remove(v V) {
	delete(s, v)
}

func (s Set[V]) Len() int {
	return len(s)
}

func (s Set[V]) Cap() int {
	return len(s)
}

func (s Set[V]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[V]) Peek() V {
	for v := range s {
		return v
	}
	var empty V
	return empty
}

func (s Set[V]) Pop() V {
	for v := range s {
		delete(s, v)
		return v
	}
	var empty V
	return empty
}

func (s Set[V]) Push(v V) bool {
	s.Add(v)
	return true
}

func (s *Set[V]) Clear() {
	*s = make(Set[V])
}
