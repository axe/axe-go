package id

import "golang.org/x/exp/constraints"

type Area[From constraints.Unsigned, To constraints.Unsigned] struct {
	tos          []To
	next         To
	resizeBuffer int
}

var _ HasCapacity = &Area[uint, uint]{}
var _ HasResizeBuffer = &Area[uint, uint]{}

func NewArea[From constraints.Unsigned, To constraints.Unsigned](opts ...Option) *Area[From, To] {
	a := &Area[From, To]{
		tos: make([]To, 0),
	}
	processOptions(a, opts)
	return a
}

func (m *Area[From, To]) setResizeBuffer(resizeBuffer int) {
	m.resizeBuffer = resizeBuffer
}

func (m *Area[From, To]) setCapacity(capacity int) {
	m.tos = make([]To, 0, capacity)
}

func (a *Area[From, To]) Translate(from From) To {
	if a == nil {
		return To(from)
	}
	if int(from) >= len(a.tos) {
		nextSize := int(from) + a.resizeBuffer + 1
		a.tos = resize(a.tos, nextSize)
	}
	if a.tos[from] == 0 {
		a.next++
		a.tos[from] = a.next
	}
	return a.tos[from] - 1
}

func (a Area[From, To]) Has(from From) bool {
	return int(from) < len(a.tos) && a.tos[from] != 0
}

func (a *Area[From, To]) Peek(from From) int {
	if a == nil {
		return int(from)
	}
	if a != nil && int(from) < len(a.tos) {
		return int(a.tos[from]) - 1
	}
	return -1
}

func (a *Area[From, To]) Remove(from From, maintainOrder bool) int {
	if int(from) >= len(a.tos) || a.tos[from] == 0 {
		return -1
	}
	a.next--
	removedTo := a.tos[from]
	a.tos[from] = 0

	if removedTo == a.next {
		return int(removedTo) - 1
	}

	if maintainOrder {
		for i, to := range a.tos {
			if to > removedTo {
				a.tos[i] = to - 1
			}
		}
	} else {
		for i := len(a.tos) - 1; i >= 0; i-- {
			if a.tos[i] == a.next {
				a.tos[i] = removedTo
				break
			}
		}
	}
	return int(removedTo) - 1
}

func (a *Area[From, To]) Clear() {
	a.tos = make([]To, 0, cap(a.tos))
	a.next = 0
}

func (a Area[From, To]) Len() int {
	return len(a.tos)
}

func (a Area[From, To]) Empty() bool {
	return len(a.tos) == 0
}
