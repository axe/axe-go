package ds

type List[T any] struct {
	Items []T
	Size  int
}

var _ Iterable[int] = &List[int]{}

func NewList[T any](capacity int) *List[T] {
	return &List[T]{
		Items: make([]T, 0, capacity),
		Size:  0,
	}
}

func (l *List[T]) Empty() bool {
	return l.Size == 0
}
func (l *List[T]) Clear() {
	l.Size = 0
}
func (l *List[T]) Pad(space int) {
	l.Items = append(l.Items, make([]T, space)...)
}
func (l *List[T]) Add(item T) {
	l.Items[l.Size] = item
	l.Size++
}
func (l *List[T]) AddList(list List[T]) {
	n := list.Len()
	for i := 0; i < n; i++ {
		l.Add(list.Get(i))
	}
}
func (l *List[T]) Len() int {
	return l.Size
}
func (l *List[T]) Get(index int) T {
	return l.Items[index]
}
func (l *List[T]) Swap(i, j int) {
	t := l.Items[i]
	l.Items[i] = l.Items[j]
	l.Items[j] = t
}
func (l *List[T]) Pop() T {
	if l.Size == 0 {
		var empty T
		return empty
	}
	l.Size--
	return l.Items[l.Size]
}
func (l *List[T]) Last() T {
	if l.Size == 0 {
		var empty T
		return empty
	}
	return l.Items[l.Size-1]
}

func (l *List[T]) IndexOf(item T) int {
	for i := 0; i < l.Size; i++ {
		if &l.Items[i] == &item {
			return i
		}
	}
	return -1
}
func (l *List[T]) RemoveAt(index int) {
	if index >= 0 && index < l.Size {
		l.Size--
		l.Items[index] = l.Items[l.Size]
	}
}

func (l *List[T]) Iterator() Iterator[T] {
	return &listIterator[T]{l, -1, false}
}

type listIterator[T any] struct {
	list    *List[T]
	index   int
	removed bool
}

func (i *listIterator[T]) Reset() {
	i.index = -1
	i.removed = false
}
func (i *listIterator[T]) HasNext() bool {
	return i.index+1 < i.list.Size
}
func (i *listIterator[T]) Next() *T {
	i.index++
	i.removed = false
	if i.index < i.list.Size {
		return &i.list.Items[i.index]
	}
	return nil
}
func (i *listIterator[T]) Remove() {
	if !i.removed && i.index >= 0 && i.index < i.list.Size {
		i.removed = true
		i.list.RemoveAt(i.index)
		i.index--
	}
}
