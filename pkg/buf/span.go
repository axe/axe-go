package buf

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
