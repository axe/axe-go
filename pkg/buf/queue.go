package buf

import "github.com/axe/axe-go/pkg/util"

type Queue[D any, B Bufferable[D]] struct {
	buffers []B
	count   int
}

var _ Iterable[int, Buffer[int]] = &Queue[int, Buffer[int]]{}
var _ HasBuffers[Buffer[int]] = &Queue[int, Buffer[int]]{}

func NewQueue[D any, B Bufferable[D]](capacity int) *Queue[D, B] {
	return &Queue[D, B]{
		buffers: make([]B, capacity),
	}
}

func (q *Queue[D, B]) Position() Position {
	return Position{current: q.count}
}

func (q *Queue[D, B]) Reset(pos Position) {
	q.count = pos.current
}

func (q *Queue[D, B]) Clear() {
	q.count = 0
}

func (q *Queue[D, B]) Add(has HasBuffers[B]) {
	if has == nil {
		return
	}
	buffers := has.GetBuffers()
	if n := len(buffers); n > 0 {
		q.buffers = util.SliceAppendAt(q.buffers, q.count, buffers)
		q.count += n
	}
}

func (q *Queue[D, B]) At(i int) *B {
	return &q.buffers[i]
}

func (q *Queue[D, B]) Len() int {
	return q.count
}

func (q *Queue[D, B]) Current() int {
	return q.count
}

func (q *Queue[D, B]) GetBuffers() []B {
	if q == nil {
		return nil
	}
	return q.buffers[:q.Len()]
}
