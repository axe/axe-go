package ds

type SparseList[T any] struct {
	items []T
	free  []int
}

func NewSparseList[T any](capacity int, freeCapacity int) *SparseList[T] {
	list := new(SparseList[T])
	list.items = make([]T, 0, capacity)
	list.free = make([]int, 0, freeCapacity)

	return list
}

func (this *SparseList[T]) At(index int) *T {
	return &this.items[index]
}

// Returns a pointer and the index to a value in the list.
// The value could be in any state since.
func (this *SparseList[T]) Take() (*T, int) {
	var index int
	lastFree := len(this.free) - 1
	if lastFree >= 0 {
		index = this.free[lastFree]
		this.free = this.free[:lastFree]
	} else {
		var value T
		index = len(this.items)
		this.items = append(this.items, value)
	}
	return &this.items[index], index
}

/**
 *
 */
func (this *SparseList[T]) Add(value T) int {
	ref, index := this.Take()
	*ref = value

	return index
}

func (this *SparseList[T]) Free(index int) {
	this.free = append(this.free, index)
}

// Removes the value at the given index and replaces it with the value at the end of the list and returns
// the index of that last item.
func (this *SparseList[T]) Remove(index int) int {
	replacedWith := len(this.items) - 1
	this.items[index] = this.items[replacedWith]
	this.items = this.items[:replacedWith]
	return replacedWith
}

func (this *SparseList[T]) Compress(moved func(newIndex int, oldIndex int, item *T)) {
	if len(this.free) > 0 {
		freeMap := this.FreeMap()

		var newIndex int
		for oi, item := range this.items {
			oldIndex := oi
			if _, exists := freeMap[oldIndex]; !exists {
				if newIndex != oldIndex {
					moved(newIndex, oldIndex, &item)
					this.items[newIndex] = item
				}
				newIndex++
			}
		}
		this.items = this.items[:newIndex]
		this.free = this.free[:0]
	}
}

func (this *SparseList[T]) FreeMap() map[int]struct{} {
	freeMap := map[int]struct{}{}
	for _, index := range this.free {
		freeMap[index] = struct{}{}
	}
	return freeMap
}

func (this *SparseList[T]) Iterate(handle func(item *T, index int, liveIndex int)) {
	if len(this.free) == 0 {
		for i := range this.items {
			index := i
			handle(&this.items[index], index, index)
		}
	} else {
		freeMap := this.FreeMap()
		liveIndex := 0
		for i := range this.items {
			index := i
			if _, exists := freeMap[index]; !exists {
				handle(&this.items[index], index, liveIndex)
				liveIndex++
			}
		}
	}
}

func (this *SparseList[T]) Pointers() []*T {
	slice := make([]*T, 0, this.Size())
	this.Iterate(func(item *T, _ int, _ int) {
		slice = append(slice, item)
	})
	return slice
}

func (this *SparseList[T]) Values() []T {
	slice := make([]T, 0, this.Size())
	this.Iterate(func(item *T, _ int, _ int) {
		slice = append(slice, *item)
	})
	return slice
}

func (this *SparseList[T]) Size() int {
	return len(this.items) - len(this.free)
}

func (this *SparseList[T]) Remaining() int {
	return cap(this.items) - len(this.items) + len(this.free)
}
