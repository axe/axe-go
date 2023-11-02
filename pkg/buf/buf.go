package buf

type Bufferable interface {
	Init(capacity int)
	Empty() bool
	Remaining() int
}

type Buffers[B Bufferable] struct {
	buffers  []B
	capacity int
	current  int
}

func NewBuffers[B Bufferable](bufferCapacity int, buffers int) *Buffers[B] {
	b := &Buffers[B]{
		buffers:  make([]B, buffers),
		capacity: bufferCapacity,
		current:  0,
	}
	for i := range b.buffers {
		b.buffers[i].Init(bufferCapacity)
	}
	return b
}

func (b *Buffers[B]) Capacity() int {
	return b.capacity
}

func (b *Buffers[B]) Clear() {
	b.current = -1
	b.AddBuffer()
}

func (b *Buffers[B]) AddBuffer() B {
	b.current++
	if b.current >= len(b.buffers) {
		var empty B
		b.buffers = append(b.buffers, empty)
	}
	added := b.buffers[b.current]
	added.Init(b.capacity)
	return added
}

func (b *Buffers[B]) Buffer() B {
	return b.At(b.current)
}

func (b *Buffers[B]) BufferFor(count int) B {
	curr := b.Buffer()
	if curr.Remaining() >= count {
		return curr
	}
	return b.AddBuffer()
}

func (b *Buffers[B]) Len() int {
	last := b.Buffer()
	if last.Empty() {
		return b.current
	}
	return b.current + 1
}

func (b *Buffers[B]) Current() int {
	return b.current
}

func (b *Buffers[B]) At(i int) B {
	if i < 0 || i > b.current {
		var empty B
		return empty
	}
	return b.buffers[i]
}

type Buffer[D any] struct {
	data       []D
	dataCount  int
	index      []int
	indexCount int
}

var _ Bufferable = &Buffer[int]{}

func (b *Buffer[D]) Init(capacity int) {
	if b.data == nil {
		b.data = make([]D, capacity)
		b.index = make([]int, capacity*3/2)
	} else {
		b.Clear()
	}
}

func (b *Buffer[D]) Clear() {
	b.dataCount = 0
	b.indexCount = 0
}

func (b *Buffer[D]) Empty() bool {
	return b.dataCount == 0
}

func (b *Buffer[D]) DataCount() int {
	return b.dataCount
}

func (b *Buffer[D]) Data() []D {
	return b.data
}

func (b *Buffer[D]) DataSpan() DataSpan[D] {
	return b.DataSpanAt(b.dataCount)
}

func (b *Buffer[D]) DataSpanAt(start int) DataSpan[D] {
	return NewDataSpan(b, start)
}

func (b *Buffer[D]) IndexCount() int {
	return b.indexCount
}

func (b *Buffer[D]) IndexSpan() IndexSpan[D] {
	return b.IndexSpanAt(b.indexCount)
}

func (b *Buffer[D]) IndexSpanAt(start int) IndexSpan[D] {
	return NewIndexSpan(b, start)
}

func (b *Buffer[D]) Index() []int {
	return b.index
}

func (b *Buffer[D]) Remaining() int {
	return len(b.data) - b.dataCount
}

func (b *Buffer[D]) Add(data ...D) int {
	index := b.dataCount
	b.data = addToSliceAt(b.data, b.dataCount, data)
	b.dataCount += len(data)
	return index
}

func (b *Buffer[D]) AddIndex(index ...int) {
	b.index = addToSliceAt(b.index, b.indexCount, index)
	b.indexCount += len(index)
}

func (b *Buffer[D]) AddRelative(data []D, relative []int) {
	index := b.Add(data...)
	startAt := b.indexCount
	b.AddIndex(relative...)
	for i := startAt; i < b.indexCount; i++ {
		b.index[i] += index
	}
}

// A view into a buffer from when the span was created to the last
// data added to the buffer.
type DataSpan[D any] struct {
	start  int
	buffer *Buffer[D]
}

func NewDataSpan[D any](buffer *Buffer[D], start int) DataSpan[D] {
	return DataSpan[D]{
		start:  start,
		buffer: buffer,
	}
}

func (s DataSpan[D]) Start() int  { return s.start }
func (s DataSpan[D]) End() int    { return s.buffer.DataCount() }
func (s DataSpan[D]) Len() int    { return s.End() - s.start }
func (s DataSpan[D]) At(i int) *D { return &s.buffer.data[s.start+i] }

// A view into a buffer from when the span was created to the last
// index added to the buffer.
type IndexSpan[D any] struct {
	start  int
	buffer *Buffer[D]
}

func NewIndexSpan[D any](buffer *Buffer[D], start int) IndexSpan[D] {
	return IndexSpan[D]{
		start:  start,
		buffer: buffer,
	}
}

func (s IndexSpan[D]) Start() int        { return s.start }
func (s IndexSpan[D]) End() int          { return s.buffer.IndexCount() }
func (s IndexSpan[D]) Len() int          { return s.End() - s.start }
func (s IndexSpan[D]) At(i int) *D       { return &s.buffer.data[s.AtIndex(i)] }
func (s IndexSpan[D]) AtIndex(i int) int { return s.buffer.index[s.start+i] }

func addToSliceAt[D any](target []D, at int, values []D) []D {
	valueCount := len(values)
	if valueCount == 0 {
		return target
	}
	space := len(target) - at
	if space >= valueCount {
		copy(target[at:], values)
		return target
	} else {
		return append(target[:at], values...)
	}
}
