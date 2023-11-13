package buf

import "github.com/axe/axe-go/pkg/util"

type Bufferable[D any] interface {
	Empty() bool
	Remaining() int
	DataCount() int
	DataAt(index int) *D
	IndexCount() int
	IndexAt(index int) int
}

type BufferInit[D any, B Bufferable[D]] func(buffer *B, capacity int)

type Buffers[D any, B Bufferable[D]] struct {
	buffers  []B
	init     BufferInit[D, B]
	capacity int
	current  int
}

func NewBuffers[D any, B Bufferable[D]](bufferCapacity int, buffers int, init BufferInit[D, B]) *Buffers[D, B] {
	b := &Buffers[D, B]{
		buffers:  make([]B, buffers),
		init:     init,
		capacity: bufferCapacity,
		current:  0,
	}
	for i := range b.buffers {
		init(&b.buffers[i], bufferCapacity)
	}
	return b
}

func (b *Buffers[D, B]) Capacity() int {
	return b.capacity
}

func (b *Buffers[D, B]) Clear() {
	b.current = -1
	b.AddBuffer()
}

func (b *Buffers[D, B]) Reserve(buffers int) {
	total := buffers + b.current + 1
	size := len(b.buffers)
	b.buffers = util.SliceEnsureSize(b.buffers, total)
	for i := size; i < total; i++ {
		b.init(&b.buffers[i], b.capacity)
	}
}

func (b *Buffers[D, B]) ReservedNext() *B {
	next := &b.buffers[b.current]
	b.current++
	return next
}

func (b *Buffers[D, B]) AddBuffer() *B {
	b.current++
	if b.current >= len(b.buffers) {
		var empty B
		b.buffers = append(b.buffers, empty)
	}
	added := &b.buffers[b.current]
	b.init(added, b.capacity)
	return added
}

func (b *Buffers[D, B]) Buffer() *B {
	return b.At(b.current)
}

func (b *Buffers[D, B]) BufferFor(count int) *B {
	curr := b.Buffer()
	if (*curr).Remaining() >= count {
		return curr
	}
	return b.AddBuffer()
}

func (b *Buffers[D, B]) Len() int {
	last := b.Buffer()
	if (*last).Empty() {
		return b.current
	}
	return b.current + 1
}

func (b *Buffers[D, B]) Current() int {
	return b.current
}

func (b *Buffers[D, B]) At(i int) *B {
	if i < 0 || i > b.current {
		return nil
	}
	return &b.buffers[i]
}

func (b *Buffers[D, B]) Empty() bool {
	return b.current == 0 && (len(b.buffers) == 0 || b.buffers[0].Empty())
}

func (b *Buffers[D, B]) ResetTo(data DataIterator[D, B]) {
	b.current = data.startBuffer
}
