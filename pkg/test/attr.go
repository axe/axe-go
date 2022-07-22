package test

import (
	"fmt"
)

type Attr[A any] interface {
	Add(other A)
	Adds(other A, scale float32)
	Scale(amount float32)
	Get() A
}

type BaseAttr[A any] struct{}

func (a *BaseAttr[A]) Adds(other A, scale float32) {}
func (a *BaseAttr[A]) Add(other A)                 { a.Adds(other, 1) }
func (a *BaseAttr[A]) Scale(amount float32)        { a.Adds(a.Get(), amount-1) }
func (a BaseAttr[A]) Get() A                       { var b A; return b }

type Vec2 struct {
	BaseAttr[Vec2]
	X float32
	Y float32
}

var _ Attr[Vec2] = &Vec2{}

func (v *Vec2) Adds(other Vec2, scale float32) {
	v.X += other.X * scale
	v.Y += other.Y * scale
}
func (v Vec2) Get() Vec2 {
	return v
}

func (a Vec2) String() string {
	return fmt.Sprintf("{%v, %v}", a.X, a.Y)
}

/*

type History[V any] struct {
	Values []TimedValue[V]
	Min    int
	Max    int
}

type TimedValue[V any] struct {
	Value V
	Time  int64
}

func NewHistory[V any](capacity int) *History[V] {
	return &History{
		Values: make([]TimedValue[V], capacity),
		Min:    0,
		Max:    0,
	}
}

func (h History[V]) wrapIndex(i int) int {
	n := len(h.Values)
	return (i + n) % n
}

func (h History[V]) maxIndex() int {
	return h.wrapIndex(h.Max - 1)
}

func (h History[V]) beforeIndex(time int64) int {
	if h.Empty() || time < h.Values[h.Min].Time {
		return -1
	}
	max := h.maxIndex()
	if time >= h.Values[max].Time {
		return -1
	}
	before := h.wrapIndex(h.Min - 1)
	for i := max; i != before; i = h.wrapIndex(i - 1) {
		if h.Values[i].Time <= time {
			return i
		}
	}
	return -1
}

func (h History[V]) Empty() bool {
	return h.Max == h.Min
}

func (h History[V]) Duration() int64 {
	if h.Empty() {
		return 0
	}
	return h.Values[h.maxIndex()].Time - h.Values[h.Min].Time
}

func (h History[V]) Before(time int64) *TimedValue {
	i := h.beforeIndex(time)
	if i == -1 {
		return nil
	}
	return &h.Values[i]
}

func (h History[V]) Around(time int64) (*TimedValue, *TimedValue) {
	i := h.beforeIndex(time)
	n := h.wrapIndex(i + 1)
	if i == -1 || n == h.Max {
		return nil, nil
	}
	return &h.Values[i], &h.Values[n]
}

func (h History[V]) Interpolate(time int64, inter func(a V, b V, delta float32) V) *V {
	before, after := h.Around(time)
	if before == nil || after == nil {
		return nil
	}
	delta := float32(time-before.Time) / float32(after.Time-before.Time)
	interpolated := inter(before.Value, after.Value, delta)
	return &interpolated
}

func (h History[V]) Closest(time int64) *TimedValue {
	var value V
	if h.Min == h.Max {
		return nil, -1
	}
	if time <= h.Times[h.Min] {
		return h.Values[h.Min]
	}
	max := h.maxIndex()
	if time >= h.Times[max] {
		return h.Values[max]
	}
	before, after := h.Around(time)
	if before == nil || after == nil {
		return value, -1
	}
	if time-before.Time < after.Time-time {
		return before
	} else {
		return after
	}
}

func (h *History[V]) Clear() {
	h.Max = 0
	h.Min = 0
}

func (h *History[V]) Add(value V, time int64) bool {
	if h.Empty() || time > h.Values[h.maxIndex()].Time {
		i := h.Max
		h.Max = h.wrapIndex(i + 1)
		tv := &h.Values[i]
		tv.Time = time
		tv.Value = value
		if h.Min == i {
			h.Min = h.Max
		}
	}
	b := h.beforeIndex(time)
}
*/
