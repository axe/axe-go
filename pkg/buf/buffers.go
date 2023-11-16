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
	buffers  []B
	init     BufferInit[B]
	reset    BufferReset[B]
	capacity int
	current  int
}

var _ Iterable[int, Buffer[int]] = &Buffers[int, Buffer[int]]{}
var _ HasBuffers[Buffer[int]] = &Buffers[int, Buffer[int]]{}

func NewBuffers[D any, B Bufferable[D]](bufferCapacity int, buffers int, init BufferInit[B], reset BufferReset[B]) *Buffers[D, B] {
	buffers = util.Max(1, buffers)
	b := &Buffers[D, B]{
		buffers:  make([]B, buffers),
		init:     init,
		reset:    reset,
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

func (b *Buffers[D, B]) SetCapacity(capacity int) {
	b.capacity = capacity
}

func (b *Buffers[D, B]) Clear() {
	b.current = 0
	for i := range b.buffers {
		b.init(&b.buffers[i], b.capacity)
	}
}

func (b *Buffers[D, B]) Reserve(buffers int) {
	total := b.current + buffers
	b.buffers = util.SliceEnsureSize(b.buffers, total)
	for i := b.current; i < total; i++ {
		b.init(&b.buffers[i], b.capacity)
	}
}

func (b *Buffers[D, B]) ReservedNext() *B {
	next := &b.buffers[b.current]
	b.current++
	return next
}

func (b *Buffers[D, B]) Add() *B {
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
	return b.Add()
}

func (b *Buffers[D, B]) Len() int {
	last := b.Buffer()
	if last == nil || (*last).Empty() {
		return b.current
	}
	return util.Min(b.current+1, len(b.buffers))
}

func (b *Buffers[D, B]) Current() int {
	return b.current
}

func (b *Buffers[D, B]) At(i int) *B {
	if i < 0 || i > b.current || i >= len(b.buffers) {
		return nil
	}
	return &b.buffers[i]
}

func (b *Buffers[D, B]) Empty() bool {
	return b.current == 0 && (len(b.buffers) == 0 || b.buffers[0].Empty())
}

func (b *Buffers[D, B]) GetBuffers() []B {
	if b == nil {
		return nil
	}
	return b.buffers[:b.Len()]
}
