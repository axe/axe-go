package ds

type SparseList[T any] struct {
	items []T
	free  Bits[uint32]
}

var _ Iterable[int] = &SparseList[int]{}
var _ Sized = &SparseList[int]{}

func NewSparseList[T any](capacity uint32) SparseList[T] {
	return SparseList[T]{
		items: make([]T, 0, capacity),
		free:  NewBits[uint32](capacity + 1),
	}
}

func (list *SparseList[T]) At(index uint32) *T {
	return &list.items[index]
}

func (list *SparseList[T]) IsFree(index uint32) bool {
	return list.free.Get(index)
}

// Returns a pointer and the index to a value in the list.
// The value could be in any state since.
func (list *SparseList[T]) Take() (*T, uint32) {
	var index uint32
	if list.free.IsEmpty() {
		var value T
		index = uint32(len(list.items))
		list.items = append(list.items, value)
		list.free.EnsureMax(index)
	} else {
		index = uint32(list.free.TakeFirst())
	}
	return &list.items[index], index
}

/**
 *
 */
func (list *SparseList[T]) Add(value T) uint32 {
	ref, index := list.Take()
	*ref = value

	return index
}

func (list *SparseList[T]) Free(index uint32) {
	list.free.Set(index, true)
}

func (list *SparseList[T]) Clear() {
	list.items = list.items[:0]
	list.free.Clear()
}

// Removes the value at the given index and replaces it with the value at the end of the list and returns
// the index of that last item.
func (list *SparseList[T]) Remove(index uint32) uint32 {
	replacedWith := uint32(len(list.items) - 1)
	list.items[index] = list.items[replacedWith]
	list.items = list.items[:replacedWith]
	return replacedWith
}

func (list *SparseList[T]) Compress(keepOrder bool, moved func(newIndex uint32, oldIndex uint32, item *T)) {
	free := &list.free
	if free.IsEmpty() {
		return
	}
	items := list.items
	if list.Size() == 0 {
		if !free.IsEmpty() {
			list.items = items[:0]
			free.Clear()
		}
	} else if keepOrder {
		var newIndex uint32
		for itemIndex, item := range items {
			oldIndex := uint32(itemIndex)
			if !free.Get(oldIndex) {
				if newIndex != oldIndex {
					if moved != nil {
						moved(newIndex, oldIndex, &item)
					}
					items[newIndex] = item
				}
				newIndex++
			}
		}
		list.items = items[:newIndex]
		free.Clear()
	} else {
		newIndex := free.TakeFirst()
		oldIndex := uint32(len(items) - 1)
		for newIndex != -1 {
			for free.Get(oldIndex) {
				oldIndex--
			}
			if int32(oldIndex) == newIndex {
				break
			}
			item := &items[oldIndex]
			if moved != nil {
				moved(uint32(newIndex), oldIndex, item)
			}
			items[newIndex] = *item

			newIndex = free.TakeFirst()
			oldIndex--
		}
		list.items = items[:oldIndex+1]
		free.Clear()
	}
}

func (list *SparseList[T]) Iterate(handle func(item *T, index uint32, liveIndex uint32) bool) {
	free := list.free
	if free.IsEmpty() {
		for i := range list.items {
			index := uint32(i)
			if !handle(&list.items[index], index, index) {
				break
			}
		}
	} else {
		liveIndex := uint32(0)
		for i := range list.items {
			if !free.Get(uint32(i)) {
				if !handle(&list.items[i], uint32(i), liveIndex) {
					break
				}
				liveIndex++
			}
		}
	}
}

func (list *SparseList[T]) Pointers() []*T {
	slice := make([]*T, 0, list.Size())
	list.Iterate(func(item *T, _ uint32, _ uint32) bool {
		slice = append(slice, item)
		return true
	})
	return slice
}

func (list *SparseList[T]) Values() []T {
	slice := make([]T, 0, list.Size())
	list.Iterate(func(item *T, _ uint32, _ uint32) bool {
		slice = append(slice, *item)
		return true
	})
	return slice
}

func (list *SparseList[T]) Size() uint32 {
	return uint32(len(list.items)) - list.free.Ons()
}

func (list *SparseList[T]) Len() int {
	return int(list.Size())
}

func (list *SparseList[T]) Cap() int {
	return len(list.items)
}

func (list *SparseList[T]) IsEmpty() bool {
	return list.Size() == 0
}

func (list *SparseList[T]) Remaining() uint32 {
	return uint32(cap(list.items)-len(list.items)) + list.free.Ons()
}

func (list *SparseList[T]) Iterator() Iterator[T] {
	return &sparseIterator[T]{list, -1}
}

type sparseIterator[T any] struct {
	list  *SparseList[T]
	index int32
}

func (i *sparseIterator[T]) Reset() {
	i.index = -1
}
func (i *sparseIterator[T]) HasNext() bool {
	n := i.nextOn()
	return n != -1 && n < i.size()
}
func (i *sparseIterator[T]) Next() *T {
	i.index = i.nextOn()
	if i.index != -1 && i.index < i.size() {
		return &i.list.items[i.index]
	}
	return nil
}
func (i *sparseIterator[T]) nextOn() int32 {
	return i.list.free.OffAfter(i.index)
}
func (i *sparseIterator[T]) size() int32 {
	return int32(len(i.list.items))
}
