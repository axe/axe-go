package util

import "reflect"

func SliceRemoveAt[E any](slice []E, index int) []E {
	return append(slice[:index], slice[index+1:]...)
}

func Copy(dst any, src any) {
	d := reflect.ValueOf(dst)
	s := reflect.ValueOf(src)
	d.Elem().Set(s.Elem())
}
