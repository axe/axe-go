package script

type Buffer[V any] struct {
	data []V
	i    int
	n    int
}

func NewBuffer[V any](initialCapacity int) Buffer[V] {
	return Buffer[V]{data: make([]V, 0, initialCapacity)}
}

func (r *Buffer[V]) Clear(initialCapacity int) {
	r.Set(make([]V, 0, initialCapacity))
}

func (r *Buffer[V]) Set(data []V) {
	r.data = data
	r.i = 0
	r.n = len(data)
}

func (r Buffer[V]) Peek() *V {
	if r.i >= r.n {
		return nil
	}
	return &r.data[r.i]
}

func (r *Buffer[V]) Read() *V {
	v := r.Peek()
	r.i++
	return v
}

func (r Buffer[V]) Pos() int {
	return r.i
}

func (r Buffer[V]) Len() int {
	return r.n
}

func (r Buffer[V]) Ended() bool {
	return r.i >= r.n
}

func (r *Buffer[V]) Reset(i int) {
	r.i = i
}

func (r *Buffer[V]) Add(data V) {
	r.data = append(r.data, data)
	r.n++
}

func (r Buffer[V]) Last() *V {
	if len(r.data) == 0 {
		return nil
	}
	return &r.data[len(r.data)-1]
}

func (r *Buffer[V]) RemoveLast() *V {
	if len(r.data) == 0 {
		return nil
	}
	r.n = len(r.data) - 1
	last := &r.data[r.n]
	r.data = r.data[:r.n]
	return last
}

type Stack[V any] struct {
	data []V
}

func (s *Stack[V]) Pop() *V {
	i := len(s.data) - 1
	if i == -1 {
		return nil
	}
	last := &s.data[i]
	s.data = s.data[:i]
	return last
}

func (s Stack[V]) Peek() *V {
	i := len(s.data) - 1
	if i == -1 {
		return nil
	}
	return &s.data[i]
}

func (s *Stack[V]) Push(value V) {
	s.data = append(s.data, value)
}

type Set[V comparable] map[V]struct{}

func NewSet[V comparable](values ...V) Set[V] {
	s := make(Set[V], len(values))
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s Set[V]) Has(value V) bool {
	_, exists := s[value]
	return exists
}

func (s Set[V]) Add(value V) {
	s[value] = struct{}{}
}

func (s Set[V]) Remove(value V) {
	delete(s, value)
}

func (s Set[V]) Values() []V {
	values := make([]V, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}
