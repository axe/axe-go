package reflect

type Iterator[K comparable, V any] interface {
	Next() bool
	Get() V
	Set(value V)
	Key() K
	Remove()
	Reset()
}

type sliceIterator[V any] struct {
	index int
	slice []V
}

var _ Iterator[int, string] = &sliceIterator[string]{}

func (i *sliceIterator[V]) Next() bool {
	i.index++
	return i.index < len(i.slice)
}
func (i sliceIterator[V]) Get() V {
	return i.slice[i.index]
}
func (i sliceIterator[V]) Set(value V) {
	i.slice[i.index] = value
}
func (i sliceIterator[V]) Key() int {
	return i.index
}
func (i *sliceIterator[V]) Remove() {
	i.slice = append(i.slice[0:i.index], i.slice[i.index+1:]...)
}
func (i *sliceIterator[V]) Reset() {
	i.index = -1
}

func IterateSlice[V any](slice []V) Iterator[int, V] {
	return &sliceIterator[V]{slice: slice, index: -1}
}

type mapIterator[K comparable, V any] struct {
	index int
	keys  []K
	m     map[K]V
}

var _ Iterator[int, string] = &mapIterator[int, string]{}

func (i *mapIterator[K, V]) Next() bool {
	i.index++
	return i.index < len(i.keys)
}
func (i mapIterator[K, V]) Get() V {
	return i.m[i.keys[i.index]]
}
func (i mapIterator[K, V]) Set(value V) {
	i.m[i.keys[i.index]] = value
}
func (i mapIterator[K, V]) Key() K {
	return i.keys[i.index]
}
func (i *mapIterator[K, V]) Remove() {
	delete(i.m, i.keys[i.index])
}
func (i *mapIterator[K, V]) Reset() {
	i.index = -1
	i.keys = make([]K, 0, len(i.m))
	for k := range i.m {
		i.keys = append(i.keys, k)
	}
}

func IterateMap[K comparable, V any](m map[K]V) Iterator[K, V] {
	i := &mapIterator[K, V]{m: m}
	i.Reset()
	return i
}

// func Transform[K comparable, V any, T any](i Iterator[K, V], transform func(V) T) Iterator[K, T] {

// }
// func Filter
// func Values
// func Keys
// func Reduce
