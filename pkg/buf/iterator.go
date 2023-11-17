package buf

import "github.com/axe/axe-go/pkg/util"

type Iterable[D any, B Bufferable[D]] interface {
	Current() int
	Len() int
	At(index int) *B
}

type DataIterator[D any, B Bufferable[D]] struct {
	// The index of the starting buffer
	startBuffer int
	// The index of the starting data. Initially this is -1 since HasNext() looks at +1.
	startData int
	// The index of the current buffer
	currentBuffer int
	// The index of the current data
	currentData int
	// If set to a positive value, the exclusive buffer end to iterate.
	limitBuffer int
	// If set to a positive value, the exclusive data end to iterate on the last buffer.
	limitData int
	// The iterable object
	iterable Iterable[D, B]
}

func NewDataIterator[D any, B Bufferable[D]](iterable Iterable[D, B], beginning bool) DataIterator[D, B] {
	start := -1
	current := 0
	if !beginning {
		current = util.Max(0, iterable.Current())
		if current < iterable.Len() {
			start = (*iterable.At(current)).DataCount() - 1
		}
	}

	return DataIterator[D, B]{
		startBuffer:   current,
		startData:     start,
		currentBuffer: current,
		currentData:   start,
		limitBuffer:   -1,
		limitData:     -1,
		iterable:      iterable,
	}
}

func (i *DataIterator[D, B]) Flip() {
	i.Limit()
	i.startBuffer = 0
	i.startData = 0
	i.Reset()
}

func (i *DataIterator[D, B]) Reset() {
	i.currentBuffer = i.startBuffer
	i.currentData = i.startData
}

func (i *DataIterator[D, B]) ClearLimit() {
	i.limitBuffer = -1
	i.limitData = -1
}

func (i *DataIterator[D, B]) Limit() {
	i.limitBuffer = i.iterable.Len()
	if i.limitBuffer > 0 {
		i.limitData = (*i.iterable.At(i.limitBuffer - 1)).DataCount()
	} else {
		i.limitData = 0
	}
}

func (i *DataIterator[D, B]) dataLimit() int {
	if i.currentBuffer >= i.bufferLimit() {
		return 0
	}
	if i.limitData != -1 && i.currentBuffer+1 >= i.limitBuffer {
		return i.limitData
	}
	return i.current().DataCount()
}

func (i *DataIterator[D, B]) bufferLimit() int {
	if i.limitBuffer != -1 {
		return i.limitBuffer
	}
	return i.iterable.Len()
}

func (i *DataIterator[D, B]) next() (buffer, data int) {
	if i.currentData+1 < i.dataLimit() {
		return i.currentBuffer, i.currentData + 1
	} else if i.currentBuffer+1 < i.bufferLimit() {
		return i.currentBuffer + 1, 0
	} else {
		return -1, -1
	}
}

func (i *DataIterator[D, B]) current() B {
	return (*i.iterable.At(i.currentBuffer))
}

func (i *DataIterator[D, B]) HasNext() bool {
	n, _ := i.next()
	return n != -1
}

func (i *DataIterator[D, B]) Next() *D {
	i.currentBuffer, i.currentData = i.next()
	return i.Current()
}

func (i *DataIterator[D, B]) Current() *D {
	return i.current().DataAt(i.currentData)
}

type IndexIterator[D any, B Bufferable[D]] struct {
	startBuffer   int
	startIndex    int
	currentBuffer int
	currentIndex  int
	limitBuffer   int
	limitIndex    int
	iterable      Iterable[D, B]
}

func NewIndexIterator[D any, B Bufferable[D]](iterable Iterable[D, B], beginning bool) IndexIterator[D, B] {
	start := -1
	current := 0
	if !beginning {
		current = util.Max(0, iterable.Current())
		if current < iterable.Len() {
			start = (*iterable.At(current)).DataCount() - 1
		}
	}

	return IndexIterator[D, B]{
		startBuffer:   current,
		startIndex:    start,
		currentBuffer: current,
		currentIndex:  start,
		limitBuffer:   -1,
		limitIndex:    -1,
		iterable:      iterable,
	}
}

func (i *IndexIterator[D, B]) Flip() {
	i.Limit()
	i.startBuffer = 0
	i.startIndex = 0
	i.Reset()
}

func (i *IndexIterator[D, B]) Reset() {
	i.currentBuffer = i.startBuffer
	i.currentIndex = i.startIndex
}

func (i *IndexIterator[D, B]) ClearLimit() {
	i.limitBuffer = -1
	i.limitIndex = -1
}

func (i *IndexIterator[D, B]) Limit() {
	i.limitBuffer = i.iterable.Len()
	if i.limitBuffer > 0 {
		i.limitIndex = (*i.iterable.At(i.limitBuffer - 1)).IndexCount()
	} else {
		i.limitIndex = 0
	}
}

func (i *IndexIterator[D, B]) indexLimit() int {
	if i.currentBuffer >= i.bufferLimit() {
		return 0
	}
	if i.limitIndex != -1 && i.currentBuffer+1 >= i.limitBuffer {
		return i.limitIndex
	}
	return i.current().IndexCount()
}

func (i *IndexIterator[D, B]) bufferLimit() int {
	if i.limitBuffer != -1 {
		return i.limitBuffer
	}
	return i.iterable.Len()
}

func (i *IndexIterator[D, B]) next() (buffer, index int) {
	if i.currentIndex+1 < i.indexLimit() {
		return i.currentBuffer, i.currentIndex + 1
	} else if i.currentBuffer+1 < i.bufferLimit() {
		return i.currentBuffer + 1, 0
	} else {
		return -1, -1
	}
}

func (i *IndexIterator[D, B]) current() B {
	return (*i.iterable.At(i.currentBuffer))
}

func (i *IndexIterator[D, B]) HasNext() bool {
	n, _ := i.next()
	return n != -1
}

func (i *IndexIterator[D, B]) Next() int {
	i.currentBuffer, i.currentIndex = i.next()
	return i.Current()
}

func (i *IndexIterator[D, B]) NextData() *D {
	i.currentBuffer, i.currentIndex = i.next()
	return i.CurrentData()
}

func (i *IndexIterator[D, B]) Current() int {
	return i.current().IndexAt(i.currentIndex)
}

func (i *IndexIterator[D, B]) CurrentData() *D {
	current := i.current()
	return current.DataAt(current.IndexAt(i.currentIndex))
}
