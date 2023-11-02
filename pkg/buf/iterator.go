package buf

type DataIterator[D any, B Bufferable[D]] struct {
	startBuffer   int
	startData     int
	currentBuffer int
	currentData   int
	buffers       *Buffers[D, B]
}

func NewDataIterator[D any, B Bufferable[D]](buffers *Buffers[D, B]) DataIterator[D, B] {
	start := -1
	if current := buffers.Buffer(); current != nil {
		start = (*current).DataCount() - 1
	}

	return DataIterator[D, B]{
		startBuffer:   buffers.current,
		startData:     start,
		currentBuffer: buffers.current,
		currentData:   start,
		buffers:       buffers,
	}
}

func (i *DataIterator[D, B]) Reset() {
	i.currentBuffer = i.startBuffer
	i.currentData = i.startData
}

func (i *DataIterator[D, B]) next() (buffer, data int) {
	if i.currentData+1 < i.current().DataCount() {
		return i.currentBuffer, i.currentData + 1
	} else if i.currentBuffer+1 < i.buffers.Len() {
		return i.currentBuffer + 1, 0
	} else {
		return -1, -1
	}
}

func (i *DataIterator[D, B]) current() B {
	return (*i.buffers.At(i.currentBuffer))
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
	buffers       *Buffers[D, B]
}

func NewIndexIterator[D any, B Bufferable[D]](buffers *Buffers[D, B]) IndexIterator[D, B] {
	start := -1
	if current := buffers.Buffer(); current != nil {
		start = (*current).IndexCount() - 1
	}

	return IndexIterator[D, B]{
		startBuffer:   buffers.current,
		startIndex:    start,
		currentBuffer: buffers.current,
		currentIndex:  start,
		buffers:       buffers,
	}
}

func (i *IndexIterator[D, B]) Reset() {
	i.currentBuffer = i.startBuffer
	i.currentIndex = i.startIndex
}

func (i *IndexIterator[D, B]) next() (buffer, index int) {
	current := i.current()
	if i.currentIndex+1 < current.IndexCount() {
		return i.currentBuffer, i.currentIndex + 1
	} else if i.currentBuffer+1 < i.buffers.Len() {
		return i.currentBuffer + 1, 0
	} else {
		return -1, -1
	}
}

func (i *IndexIterator[D, B]) current() B {
	return (*i.buffers.At(i.currentBuffer))
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
