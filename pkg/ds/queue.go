package ds

type Queue[V any] interface {
	Iterable[V]
	Linear[V]
	Clearable
	Sized
}

type CircularQueue[V any] struct {
	values []V
	head   int
	tail   int
}

var _ Queue[int] = &CircularQueue[int]{}
var _ Indexed[int] = &CircularQueue[int]{}

func NewCircularQueue[V any](initialCapacity int) CircularQueue[V] {
	return CircularQueue[V]{
		values: make([]V, initialCapacity),
		head:   0,
		tail:   0,
	}
}

func (q *CircularQueue[V]) At(index int) V {
	var value V
	if index < q.Len() && index >= 0 {
		value = q.values[q.wrap(q.head+index)]
	}
	return value
}
func (q *CircularQueue[V]) RemoveAt(index int) {
	if index < q.Len() && index >= 0 {
		actual := q.wrap(q.head + index)
		if actual < q.tail {
			copy(q.values[actual:q.tail-1], q.values[actual+1:q.tail])
		} else {
			first := q.values[0]
			n := len(q.values)
			copy(q.values[0:q.tail-1], q.values[1:q.tail])
			copy(q.values[actual:n-1], q.values[actual+1:n])
			q.values[n-1] = first
		}
	}
}
func (q *CircularQueue[V]) Peek() V {
	var value V
	if q.tail != q.head {
		value = q.values[q.head]
	}
	return value
}
func (q *CircularQueue[V]) Pop() V {
	var value V
	if q.tail != q.head {
		value = q.values[q.head]
		q.head = q.wrap(q.head + 1)
	}
	return value
}
func (q *CircularQueue[V]) Push(value V) bool {
	if q.Len() == q.Cap() {
		if q.tail < q.head {
			q.values = append(q.values, q.values[0])
			copy(q.values[1:q.tail], q.values[0:q.tail])
		} else {
			var empty V
			q.values = append(q.values, empty)
		}
	}
	q.values[q.tail] = value
	q.tail = q.wrap(q.tail + 1)
	return true
}
func (q CircularQueue[V]) Len() int {
	n := len(q.values)
	return ((q.tail + n) - q.head) % n
}
func (q CircularQueue[V]) Cap() int {
	return len(q.values)
}
func (q CircularQueue[V]) IsEmpty() bool {
	return q.head == q.tail
}
func (q CircularQueue[V]) wrap(i int) int {
	return i % len(q.values)
}
func (q *CircularQueue[V]) Clear() {
	q.head = 0
	q.tail = 0
}
func (q *CircularQueue[V]) Iterator() Iterator[V] {
	return &circularQueueIterable[V]{q, q.head - 1, false}
}

type circularQueueIterable[V any] struct {
	queue   *CircularQueue[V]
	index   int
	removed bool
}

var _ Iterator[int] = &circularQueueIterable[int]{}

func (i *circularQueueIterable[V]) Reset() {
	i.index = i.queue.head - 1
	i.removed = false
}
func (i *circularQueueIterable[V]) HasNext() bool {
	return i.nextIndex() != i.queue.tail
}
func (i *circularQueueIterable[V]) Next() *V {
	next := i.nextIndex()
	if next == i.queue.tail {
		return nil
	}
	i.index = next
	i.removed = false
	return &i.queue.values[next]
}
func (i *circularQueueIterable[V]) nextIndex() int {
	return i.queue.wrap(i.index + 1)
}
func (i *circularQueueIterable[V]) Remove() {
	if !i.removed && i.index >= i.queue.head && i.index != i.queue.tail {
		i.queue.RemoveAt(i.index)
		i.index--
		i.removed = true
	}
}

type LinkQueue[V any] struct {
	head *linkQueueNode[V]
	tail *linkQueueNode[V]
	size int
}

type linkQueueNode[V any] struct {
	value V
	next  *linkQueueNode[V]
}

var _ Queue[int] = &LinkQueue[int]{}

func (q *LinkQueue[V]) Peek() V {
	var value V
	if q.size != 0 {
		value = q.head.value
	}
	return value
}
func (q *LinkQueue[V]) Pop() V {
	var value V
	if q.size != 0 {
		value = q.head.value
		q.head = q.head.next
		q.size--
		if q.size == 0 {
			q.head = nil
			q.tail = nil
		}
	}
	return value
}
func (q *LinkQueue[V]) Push(value V) bool {
	node := &linkQueueNode[V]{value, nil}
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
	return true
}
func (q *LinkQueue[V]) Len() int {
	return q.size
}
func (q *LinkQueue[V]) Cap() int {
	return -1
}
func (q *LinkQueue[V]) IsEmpty() bool {
	return q.size == 0
}
func (q *LinkQueue[V]) Clear() {
	q.head = nil
	q.tail = nil
	q.size = 0
}
func (q *LinkQueue[V]) Iterator() Iterator[V] {
	return &linkQueueIterator[V]{q, q.head, nil, false}
}

type linkQueueIterator[V any] struct {
	queue    *LinkQueue[V]
	current  *linkQueueNode[V]
	previous *linkQueueNode[V]
	removed  bool
}

var _ Iterator[int] = &linkQueueIterator[int]{}

func (i *linkQueueIterator[V]) Reset() {
	i.current = i.queue.head
	i.previous = nil
	i.removed = false
}
func (i *linkQueueIterator[V]) HasNext() bool {
	return i.current != nil
}
func (i *linkQueueIterator[V]) Next() *V {
	if i.current == nil {
		return nil
	}
	value := &i.current.value
	i.previous = i.current
	i.current = i.current.next
	i.removed = false
	return value
}
func (i *linkQueueIterator[V]) Remove() {
	if !i.removed && i.current != nil {
		if i.previous == nil {
			i.queue.head = i.current.next
		} else {
			i.previous.next = i.current.next
			i.current = i.previous
		}
		i.queue.size--
		i.removed = true
	}
}
