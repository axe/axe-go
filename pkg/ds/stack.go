package ds

type Stack[T any] struct {
	Items []T
	Count int
}

var _ Linear[int] = &Stack[int]{}

func NewStack[T any](capacity uint32) Stack[T] {
	return Stack[T]{
		Items: make([]T, capacity),
		Count: 0,
	}
}

func (s Stack[T]) IsEmpty() bool {
	return s.Count == 0
}

func (s *Stack[T]) Push(value T) bool {
	if s.Count == len(s.Items) {
		s.Items = append(s.Items, value)
	} else {
		s.Items[s.Count] = value
	}
	s.Count++
	return true
}

func (s *Stack[T]) Pop() T {
	var empty T
	if s.Count == 0 {
		return empty
	}
	s.Count--
	value := s.Items[s.Count]
	s.Items[s.Count] = empty
	return value
}

func (s *Stack[T]) Peek() T {
	var value T
	if s.Count != 0 {
		value = s.Items[s.Count-1]
	}
	return value
}

func (s *Stack[T]) Clear() {
	s.Count = 0
}
