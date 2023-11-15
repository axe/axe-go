package buf

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	bufs := NewBuffers[int](128, 4, func(buffer *Buffer[int], capacity int) {
		buffer.Init(capacity)
	}, func(buffer *Buffer[int], pos Position) {
		buffer.Reset(pos)
	})

	iter := NewDataIterator(bufs)
	bufs.Buffer().Add(1, 2, 3)

	for iter.HasNext() {
		fmt.Printf("%d\n", *iter.Next())
	}
}
