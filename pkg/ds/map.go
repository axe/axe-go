package ds

import "github.com/axe/axe-go/pkg/util"

type mapValueIterator[K comparable, V any] struct {
	m     map[K]V
	keys  []K
	index int
}

var _ Iterator[int] = &mapValueIterator[int, int]{}

func NewMapValueIterator[K comparable, V any](m map[K]V) Iterator[V] {
	return &mapValueIterator[K, V]{m, util.MapKeys(m), -1}
}

func NewMapValueIterable[K comparable, V any](m map[K]V) Iterable[V] {
	return NewIterable(func() Iterator[V] {
		return NewMapValueIterator(m)
	})
}

func (iter *mapValueIterator[K, V]) Reset() {
	iter.index = -1
	iter.keys = util.MapKeys(iter.m)
}

func (iter *mapValueIterator[K, V]) HasNext() bool {
	return iter.index+1 < len(iter.keys)
}

func (iter *mapValueIterator[K, V]) Next() *V {
	iter.index++
	if iter.index >= len(iter.keys) {
		return nil
	}
	key := iter.keys[iter.index]
	value := iter.m[key]
	return &value
}

func (iter *mapValueIterator[K, V]) Remove() {
	if iter.index >= 0 && iter.index < len(iter.keys) {
		delete(iter.m, iter.keys[iter.index])
	}
}

type mapKeyIterator[K comparable, V any] struct {
	m     map[K]V
	keys  []K
	index int
}

var _ Iterator[int] = &mapKeyIterator[int, int]{}

func NewMapKeyIterator[K comparable, V any](m map[K]V) Iterator[K] {
	return &mapKeyIterator[K, V]{m, util.MapKeys(m), -1}
}

func NewMapKeyIterable[K comparable, V any](m map[K]V) Iterable[K] {
	return NewIterable(func() Iterator[K] {
		return NewMapKeyIterator(m)
	})
}

func (iter *mapKeyIterator[K, V]) Reset() {
	iter.index = -1
	iter.keys = util.MapKeys(iter.m)
}

func (iter *mapKeyIterator[K, V]) HasNext() bool {
	return iter.index+1 < len(iter.keys)
}

func (iter *mapKeyIterator[K, V]) Next() *K {
	iter.index++
	if iter.index >= len(iter.keys) {
		return nil
	}
	key := iter.keys[iter.index]
	return &key
}

func (iter *mapKeyIterator[K, V]) Remove() {
	if iter.index >= 0 && iter.index < len(iter.keys) {
		delete(iter.m, iter.keys[iter.index])
	}
}

type MapKeyValue[K comparable, V any] struct {
	m     map[K]V
	Key   K
	Value V
}

func (kv *MapKeyValue[K, V]) Delete() {
	delete(kv.m, kv.Key)
}

func (kv *MapKeyValue[K, V]) Exists() bool {
	_, exists := kv.m[kv.Key]
	return exists
}

func (kv *MapKeyValue[K, V]) Rekey(key K) bool {
	if _, exists := kv.m[kv.Key]; exists {
		delete(kv.m, kv.Key)
		kv.m[key] = kv.Value
		kv.Key = key
		return true
	}
	return false
}

func (kv *MapKeyValue[K, V]) Set(value V) {
	kv.m[kv.Key] = value
	kv.Value = value
}

func (kv *MapKeyValue[K, V]) Refresh() {
	kv.Value = kv.m[kv.Key]
}

type mapKeyValueIterator[K comparable, V any] struct {
	m     map[K]V
	keys  []K
	index int
}

var _ Iterator[int] = &mapKeyIterator[int, int]{}

func NewMapKeyValueIterator[K comparable, V any](m map[K]V) Iterator[MapKeyValue[K, V]] {
	return &mapKeyValueIterator[K, V]{m, util.MapKeys(m), -1}
}

func NewMapKeyValueIterable[K comparable, V any](m map[K]V) Iterable[MapKeyValue[K, V]] {
	return NewIterable(func() Iterator[MapKeyValue[K, V]] {
		return NewMapKeyValueIterator(m)
	})
}

func (iter *mapKeyValueIterator[K, V]) Reset() {
	iter.index = -1
	iter.keys = util.MapKeys(iter.m)
}

func (iter *mapKeyValueIterator[K, V]) HasNext() bool {
	return iter.index+1 < len(iter.keys)
}

func (iter *mapKeyValueIterator[K, V]) Next() *MapKeyValue[K, V] {
	iter.index++
	if iter.index >= len(iter.keys) {
		return nil
	}
	key := iter.keys[iter.index]
	value := iter.m[key]
	return &MapKeyValue[K, V]{iter.m, key, value}
}

func (iter *mapKeyValueIterator[K, V]) Remove() {
	if iter.index >= 0 && iter.index < len(iter.keys) {
		delete(iter.m, iter.keys[iter.index])
	}
}
