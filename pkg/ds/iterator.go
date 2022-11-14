package ds

import "github.com/axe/axe-go/pkg/util"

type Iterator[T any] interface {
	Reset()
	HasNext() bool
	Next() *T
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type factoryIterable[T any] func() Iterator[T]

var _ Iterable[int] = factoryIterable[int](nil)

func NewIterable[T any](create func() Iterator[T]) Iterable[T] {
	return factoryIterable[T](create)
}

func (iter factoryIterable[T]) Iterator() Iterator[T] {
	return iter()
}

type sliceIterator[T any] struct {
	data  []T
	index int
}

var _ Iterator[int] = &sliceIterator[int]{}

func NewSliceIterator[T any](data []T) Iterator[T] {
	return &sliceIterator[T]{data, -1}
}

func NewSliceIterable[T any](data []T) Iterable[T] {
	return NewIterable(func() Iterator[T] {
		return NewSliceIterator(data)
	})
}

func (iter *sliceIterator[T]) Reset() {
	iter.index = -1
}

func (iter *sliceIterator[T]) HasNext() bool {
	return iter.index+1 < len(iter.data)
}

func (iter *sliceIterator[T]) Next() *T {
	iter.index++
	if iter.index >= len(iter.data) {
		return nil
	}
	return &iter.data[iter.index]
}

type multiIterator[T any] struct {
	iterators []Iterator[T]
	index     int
}

var _ Iterator[int] = &multiIterator[int]{}

func NewMultiIterator[T any](data []Iterator[T]) Iterator[T] {
	return &multiIterator[T]{data, 0}
}

func NewMultiIterable[T any](iterables []Iterable[T]) Iterable[T] {
	return NewIterable(func() Iterator[T] {
		return NewMultiIterator(util.SliceMap(iterables, func(s Iterable[T]) Iterator[T] {
			return s.Iterator()
		}))
	})
}

func (iter *multiIterator[T]) Reset() {
	for _, inner := range iter.iterators {
		inner.Reset()
	}
	iter.index = 0
}

func (iter *multiIterator[T]) HasNext() bool {
	for i := iter.index; i < len(iter.iterators); i++ {
		if iter.iterators[i].HasNext() {
			return true
		}
	}
	return false
}

func (iter *multiIterator[T]) Next() *T {
	for iter.index < len(iter.iterators) {
		if iter.iterators[iter.index].HasNext() {
			return iter.iterators[iter.index].Next()
		} else {
			iter.index++
		}
	}
	return nil
}

type filterIterator[T any] struct {
	inner  Iterator[T]
	filter func(value *T) bool
	next   *T
}

var _ Iterator[int] = &filterIterator[int]{}

func NewFilterIterator[T any](iter Iterator[T], filter func(value *T) bool) Iterator[T] {
	return &filterIterator[T]{iter, filter, nil}
}

func NewFilterIterable[T any](iter Iterable[T], filter func(value *T) bool) Iterable[T] {
	return NewIterable(func() Iterator[T] {
		return NewFilterIterator(iter.Iterator(), filter)
	})
}

func (iter *filterIterator[T]) Reset() {
	iter.inner.Reset()
	iter.next = nil
}

func (iter *filterIterator[T]) HasNext() bool {
	iter.updateNext()
	return iter.next != nil
}

func (iter *filterIterator[T]) Next() *T {
	iter.updateNext()
	next := iter.next
	iter.next = nil
	return next
}

func (iter *filterIterator[T]) updateNext() {
	if iter.next != nil {
		return
	}
	for iter.inner.HasNext() {
		iter.next = iter.inner.Next()
		if iter.filter(iter.next) {
			return
		}
	}
	iter.next = nil
}

type translateIterator[T any, S any] struct {
	inner     Iterator[S]
	translate func(value *S) *T
}

var _ Iterator[int] = &translateIterator[int, int]{}

func NewTranslateIterator[T any, S any](iter Iterator[S], translate func(value *S) *T) Iterator[T] {
	return &translateIterator[T, S]{iter, translate}
}

func NewTranslateIterable[T any, S any](iter Iterable[S], translate func(value *S) *T) Iterable[T] {
	return NewIterable(func() Iterator[T] {
		return NewTranslateIterator(iter.Iterator(), translate)
	})
}

func (iter *translateIterator[T, S]) Reset() {
	iter.inner.Reset()
}

func (iter *translateIterator[T, S]) HasNext() bool {
	return iter.inner.HasNext()
}

func (iter *translateIterator[T, S]) Next() *T {
	source := iter.inner.Next()
	if source != nil {
		return iter.translate(source)
	}
	return nil
}

type emptyIterator[T any] struct{}

var _ Iterator[int] = &emptyIterator[int]{}

func NewEmptyIterator[T any]() Iterator[T] {
	return &emptyIterator[T]{}
}

func NewEmptyIterable[T any]() Iterable[T] {
	return NewIterable(func() Iterator[T] {
		return NewEmptyIterator[T]()
	})
}

func (iter *emptyIterator[T]) Reset()        {}
func (iter *emptyIterator[T]) HasNext() bool { return false }
func (iter *emptyIterator[T]) Next() *T      { return nil }
