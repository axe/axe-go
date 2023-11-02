package buf

type Buffer[D any] struct {
	data       []D
	dataCount  int
	index      []int
	indexCount int
}

var _ Bufferable[int] = Buffer[int]{}

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

func (b Buffer[D]) Empty() bool {
	return b.dataCount == 0
}

func (b Buffer[D]) DataCount() int {
	return b.dataCount
}

func (b *Buffer[D]) Data() []D {
	return b.data[:b.dataCount]
}

func (b Buffer[D]) DataAt(index int) *D {
	return &b.data[index]
}

func (b *Buffer[D]) DataSpan() DataSpan[D] {
	return b.DataSpanAt(b.dataCount)
}

func (b *Buffer[D]) DataSpanAt(start int) DataSpan[D] {
	return NewDataSpan(b, start)
}

func (b Buffer[D]) IndexCount() int {
	return b.indexCount
}

func (b *Buffer[D]) IndexSpan() IndexSpan[D] {
	return b.IndexSpanAt(b.indexCount)
}

func (b *Buffer[D]) IndexSpanAt(start int) IndexSpan[D] {
	return NewIndexSpan(b, start)
}

func (b *Buffer[D]) Index() []int {
	return b.index[:b.indexCount]
}

func (b Buffer[D]) IndexAt(index int) int {
	return b.index[index]
}

func (b Buffer[D]) Remaining() int {
	return len(b.data) - b.dataCount
}

func (b *Buffer[D]) Add(data ...D) int {
	index := b.dataCount
	b.data = addToSliceAt(b.data, b.dataCount, data)
	b.dataCount += len(data)
	return index
}

func (b *Buffer[D]) AddIndexed(data ...D) {
	index := b.Add(data...)
	startAt := b.indexCount
	b.AddIndex(make([]int, len(data))...)
	for i := startAt; i < b.indexCount; i++ {
		b.index[i] += index
		index++
	}
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
