package ds

type Stack[T any] struct {
	Items []T
	Count int
}

func NewStack[T any](capacity uint32) Stack[T] {
	return Stack[T]{
		Items: make([]T, capacity),
		Count: 0,
	}
}

func (s Stack[T]) IsEmpty() bool {
	return s.Count == 0
}

func (s *Stack[T]) Push(value T) {
	if s.Count == len(s.Items) {
		s.Items = append(s.Items, value)
	} else {
		s.Items[s.Count] = value
	}
	s.Count++
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
