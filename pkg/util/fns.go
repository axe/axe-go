package util

import (
	"fmt"
	"reflect"
)

func SliceRemoveAt[E any](slice []E, index int) []E {
	return append(slice[:index], slice[index+1:]...)
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
