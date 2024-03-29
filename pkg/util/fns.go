package util

import (
	"fmt"
	"reflect"
	"sort"
)

func SliceCopy[E any](slice []E) []E {
	return append(make([]E, 0, len(slice)), slice...)
}

func SliceRemoveAt[E any](slice []E, index int) []E {
	return append(slice[:index], slice[index+1:]...)
}

func SliceRemoveAtReplace[E any](slice []E, index int) []E {
	last := len(slice) - 1
	slice[index] = slice[last]
	return slice[:last]
}

func SliceRemove[E comparable](slice []E, value E) []E {
	for i := range slice {
		if slice[i] == value {
			return SliceRemoveAt(slice, i)
		}
	}
	return slice
}

func SliceMap[S any, T any](source []S, transform func(value S) T) []T {
	transformed := make([]T, len(source))
	for i := range source {
		transformed[i] = transform(source[i])
	}
	return transformed
}

func SliceMove[V any](slice []V, from, to int) {
	if from == to || to < 0 || from < 0 || from >= len(slice) || to >= len(slice) {
		return
	}
	if from < to {
		value := slice[from]
		copy(slice[from:to], slice[from+1:to+1])
		slice[to] = value
	} else {
		value := slice[to]
		copy(slice[to+1:from+1], slice[to:from])
		slice[from] = value
	}
}

func SliceAppendAt[D any](target []D, at int, values []D) []D {
	valueCount := len(values)
	if valueCount == 0 {
		return target
	}
	space := len(target) - at
	if space >= valueCount {
		copy(target[at:], values)
		return target
	} else {
		return append(target[:at], values...)
	}
}

func SliceIndexOf[V comparable](slice []V, value V) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}
	return -1
}

func SliceResize[V any](slice []V, size int) []V {
	existingSize := len(slice)
	if size == existingSize {
		return slice
	} else if size < existingSize {
		return slice[:size]
	} else {
		return append(slice, make([]V, size-existingSize)...)
	}
}

func SliceEnsureSize[V any](slice []V, size int) []V {
	existingSize := len(slice)
	if size > existingSize {
		return append(slice, make([]V, size-existingSize)...)
	} else {
		return slice
	}
}

func SliceReverse[E any](s []E) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SliceSorted[E any](slice []E, isLess func(E, E) bool) []E {
	copy := SliceCopy(slice)
	sort.Slice(copy, func(i, j int) bool {
		return isLess(copy[i], copy[j])
	})
	return copy
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for k := range m {
		values = append(values, m[k])
	}
	return values
}

func CoalesceJoin[V any](a, b V, swap bool, isNil func(V) bool, join func(V, V) V) V {
	if isNil(a) {
		return b
	} else if !isNil(b) {
		var first, second V
		if swap {
			first = b
			second = a
		} else {
			first = a
			second = b
		}

		return join(first, second)
	}
	return a
}

func Coalesce[V any](nilable *V, nonNil V) V {
	if nilable != nil {
		return *nilable
	}
	return nonNil
}

func Coalesce2[V any](nilable1 *V, nilable2 *V, nonNil V) V {
	if nilable1 != nil {
		return *nilable1
	}
	if nilable2 != nil {
		return *nilable2
	}
	return nonNil
}

func Clone[V any](nilable *V) *V {
	if nilable == nil {
		return nil
	}
	copy := *nilable
	return &copy
}

func If[V any](condition bool, trueValue, falseValue V) V {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func Copy(dst any, src any) {
	d := reflect.ValueOf(dst)
	s := reflect.ValueOf(src)
	d.Elem().Set(s.Elem())
}

func Assert(condition bool, messageFormat string, args ...any) {
	if !condition {
		panic(fmt.Sprintf(messageFormat, args...))
	}
}
