package ds;


type SparseList[T any] struct {
	items		[]T;
	free		[]uint32;
}

func NewSparseList[T any](capacity uint32, freeCapacity uint32) *SparseList[T] {
	list := new(SparseList[T]);
	list.items = make([]T, 0, capacity);
	list.free = make([]uint32, 0, freeCapacity);

	return list;
}

func (this *SparseList[T]) At(index uint32) *T {
	return &this.items[index];
}

// Returns a pointer and the index to a value in the list. 
// The value could be in any state since.
func (this *SparseList[T]) Take() (*T, uint32) {
	var index uint32
	lastFree := len(this.free) - 1;
	if (lastFree >= 0) {
		index = this.free[lastFree];
		this.free = this.free[:lastFree]
	} else {
		var value T
		index = uint32(len(this.items));
		this.items = append(this.items, value);
	}
	return &this.items[index], index;
}

/**
 *
 */
func (this *SparseList[T]) Add(value T) uint32 {
	ref, index := this.Take();
	*ref = value;

	return index;
}

func (this *SparseList[T]) Free(index uint32) {
	this.free = append(this.free, index);
}

// Removes the value at the given index and replaces it with the value at the end of the list and returns
// the index of that last item.
func (this *SparseList[T]) Remove(index int) int {
	replacedWith := len(this.items) - 1;
	this.items[index] = this.items[replacedWith];
	this.items = this.items[:replacedWith];
	return replacedWith;
}

func (this *SparseList[T]) Compress(moved func (newIndex uint32, oldIndex uint32, item *T)) {
	if (len(this.free) > 0) {
		freeMap := this.FreeMap();

		var newIndex uint32;
		for oi, item := range this.items {
			oldIndex := uint32(oi);
			if _, exists := freeMap[oldIndex]; !exists {
				if newIndex != oldIndex {
					moved(newIndex, oldIndex, &item);
					this.items[newIndex] = item;
				}
				newIndex++;
			}
		}
		this.items = this.items[:newIndex];
		this.free = this.free[:0];
	}
}

func (this *SparseList[T]) FreeMap() map[uint32]struct{} {
	freeMap := map[uint32]struct{};
	for _, index := range this.free {
		freeMap[index] = struct{}{};
	}
	return freeMap;
}

func (this *SparseList[T]) Iterate(handle func (item *T, index uint32, liveIndex uint32)) {
	if (len(this.free) == 0) {
		for i := range this.items {
			index := uint32(i);
			handle(&this.items[index], index, index);
		}
	} else {
		freeMap := this.FreeMap();
		liveIndex := uint32(0);
		for i := range this.items {
			index := uint32(i);
			if _, exists := freeMap[index]; !exists {
				handle(&this.items[index], index, liveIndex);
				liveIndex++;
			}
		}
	}
}

func (this *SparseList[T]) Pointers() []*T {
	slice := make([]*T, 0, this.Size());
	this.Iterate(func (item *T, _ uint32, _ uint32) {
		slice = append(slice, item);
	});
	return slice;
}

func (this *SparseList[T]) Values() []T {
	slice := make([]T, 0, this.Size());
	this.Iterate(func (item *T, _ uint32, _ uint32) {
		slice = append(slice, *item);
	});
	return slice;
}

func (this *SparseList[T]) Size() int {
	return len(this.items) - len(this.free);
}

func (this *SparseList[T]) Remaining() int {
	return cap(this.items) - len(this.items) + len(this.free);
}