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

type HasBuffers[B any] interface {
	GetBuffers() []B
}

type BufferInit[B any] func(buffer *B, capacity int)

type BufferReset[B any] func(buffer *B, pos Position)

type Buffers[D any, B Bufferable[D]] struct {
	// The underlying buffers
	buffers []B
	// The function to use to initial a buffer before we start using it.
	init BufferInit[B]
	// The function to use to reset a buffer to a specific position.
	reset BufferReset[B]
	// The capacity to initialize new buffers with.
	capacity int
	// The number of buffers. Len may be one less if the last buffer is empty.
	count int
}

var _ Iterable[int, Buffer[int]] = &Buffers[int, Buffer[int]]{}
var _ HasBuffers[Buffer[int]] = &Buffers[int, Buffer[int]]{}

func NewBuffers[D any, B Bufferable[D]](bufferCapacity int, buffers int, init BufferInit[B], reset BufferReset[B]) *Buffers[D, B] {
	b := &Buffers[D, B]{
		buffers:  make([]B, buffers),
		init:     init,
		reset:    reset,
		capacity: bufferCapacity,
		count:    0,
	}
	for i := range b.buffers {
		init(&b.buffers[i], bufferCapacity)
	}
	return b
}

func (b *Buffers[D, B]) Capacity() int {
	return b.capacity
}

func (b *Buffers[D, B]) SetCapacity(capacity int) {
	b.capacity = capacity
}

func (b *Buffers[D, B]) Clear() {
	for b.count > 0 {
		b.count--
		b.init(&b.buffers[b.count], b.capacity)
	}
}

func (b *Buffers[D, B]) Reserve(buffers int) {
	required := b.Len() + buffers
	existing := len(b.buffers)
	b.buffers = util.SliceEnsureSize(b.buffers, required)
	for i := existing; i < required; i++ {
		b.init(&b.buffers[i], b.capacity)
	}
}

func (b *Buffers[D, B]) ReservedNext() *B {
	if !b.buffers[b.count].Empty() {
		b.count++
	}
	next := &b.buffers[b.count]
	b.count++
	return next
}

func (b *Buffers[D, B]) Add() *B {
	if b.count >= len(b.buffers) {
		var empty B
		b.init(&empty, b.capacity)
		b.buffers = append(b.buffers, empty)
	}
	added := &b.buffers[b.count]
	b.count++
	return added
}

func (b *Buffers[D, B]) Buffer() *B {
	if b.count == 0 {
		return b.Add()
	} else {
		return b.At(b.count - 1)
	}
}

func (b *Buffers[D, B]) BufferFor(count int) *B {
	curr := b.Buffer()
	if (*curr).Remaining() >= count {
		return curr
	}
	return b.Add()
}

func (b *Buffers[D, B]) Len() int {
	n := b.count
	if n > 0 && b.buffers[n-1].Empty() {
		n--
	}
	return n
}

func (b *Buffers[D, B]) Current() int {
	return b.count - 1
}

func (b *Buffers[D, B]) At(i int) *B {
	if i < 0 || i >= b.count {
		return nil
	}
	return &b.buffers[i]
}

func (b *Buffers[D, B]) Empty() bool {
	return b.count == 0 || (b.count == 1 && b.buffers[0].Empty())
}

func (b *Buffers[D, B]) GetBuffers() []B {
	if b == nil {
		return nil
	}
	return b.buffers[:b.Len()]
}
