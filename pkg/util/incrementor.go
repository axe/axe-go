package util

import "golang.org/x/exp/constraints"

type Incrementor[I constraints.Ordered] struct {
	current   I
	increment I
}

func NewIncrementor[I constraints.Ordered](initial, increment I) Incrementor[I] {
	return Incrementor[I]{current: initial, increment: increment}
}

func (i Incrementor[I]) Peek() I {
	return i.current
}

func (i *Incrementor[I]) Get() I {
	next := i.current
	i.current += i.increment
	return next
}
