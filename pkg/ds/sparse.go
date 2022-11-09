package ds

type SparseList[T any] struct {
	items []T
	free  Bits
}

func NewSparseList[T any](capacity uint32) SparseList[T] {
	return SparseList[T]{
		items: make([]T, 0, capacity),
		free:  NewBits(capacity + 1),
	}
}

func (this *SparseList[T]) At(index uint32) *T {
	return &this.items[index]
}

func (this *SparseList[T]) IsFree(index uint32) bool {
	return this.free.Get(index)
}

// Returns a pointer and the index to a value in the list.
// The value could be in any state since.
func (this *SparseList[T]) Take() (*T, uint32) {
	var index uint32
	if this.free.IsEmpty() {
		var value T
		index = uint32(len(this.items))
		this.items = append(this.items, value)
		this.free.EnsureMax(index)
	} else {
		index = uint32(this.free.TakeFirst())
	}
	return &this.items[index], index
}

/**
 *
 */
func (this *SparseList[T]) Add(value T) uint32 {
	ref, index := this.Take()
	*ref = value

	return index
}

func (this *SparseList[T]) Free(index uint32) {
	this.free.Set(index, true)
}

// Removes the value at the given index and replaces it with the value at the end of the list and returns
// the index of that last item.
func (this *SparseList[T]) Remove(index uint32) uint32 {
	replacedWith := uint32(len(this.items) - 1)
	this.items[index] = this.items[replacedWith]
	this.items = this.items[:replacedWith]
	return replacedWith
}

func (this *SparseList[T]) Compress(keepOrder bool, moved func(newIndex uint32, oldIndex uint32, item *T)) {
	free := &this.free
	if free.IsEmpty() {
		return
	}
	items := this.items
	if this.Size() == 0 {
		if !free.IsEmpty() {
			this.items = items[:0]
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
		this.items = items[:newIndex]
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
		this.items = items[:oldIndex+1]
		free.Clear()
	}
}

func (this *SparseList[T]) Iterate(handle func(item *T, index uint32, liveIndex uint32) bool) {
	free := this.free
	if free.IsEmpty() {
		for i := range this.items {
			index := uint32(i)
			if !handle(&this.items[index], index, index) {
				break
			}
		}
	} else {
		liveIndex := uint32(0)
		for i := range this.items {
			if !free.Get(uint32(i)) {
				if !handle(&this.items[i], uint32(i), liveIndex) {
					break
				}
				liveIndex++
			}
		}
	}
}

func (this *SparseList[T]) Pointers() []*T {
	slice := make([]*T, 0, this.Size())
	this.Iterate(func(item *T, _ uint32, _ uint32) bool {
		slice = append(slice, item)
		return true
	})
	return slice
}

func (this *SparseList[T]) Values() []T {
	slice := make([]T, 0, this.Size())
	this.Iterate(func(item *T, _ uint32, _ uint32) bool {
		slice = append(slice, *item)
		return true
	})
	return slice
}

func (this *SparseList[T]) Size() uint32 {
	return uint32(len(this.items)) - this.free.Ons()
}

func (this *SparseList[T]) Remaining() uint32 {
	return uint32(cap(this.items)-len(this.items)) + this.free.Ons()
}
