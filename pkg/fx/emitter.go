package fx

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

type EmitterType struct {
	Delay     Range[float32]
	Frequency Range[float32]
	Size      Range[int]
	Count     Range[int]
}

func (et EmitterType) Create() Emitter {
	return Emitter{
		Delay:     et.Delay.Get(),
		Frequency: et.Frequency,
		Size:      et.Size,
		Count:     et.Count.Get(),
	}
}

type Emitter struct {
	Delay     float32
	Frequency Range[float32]
	Size      Range[int]
	Count     int

	Time   float32
	Next   float32
	Bursts int
}

func (e *Emitter) Update(dt float32) (size int, over float32) {
	if e.Done() {
		return
	}
	if !e.Started() {
		e.Next = e.Delay
	}
	e.Time += dt
	if e.Time >= e.Next {
		over = e.Next - e.Time
		size = e.Size.Get()
		e.Time -= e.Next
		e.Bursts++
		e.Next = e.Frequency.Get()
	}
	return
}
func (e Emitter) Started() bool {
	return e.Bursts > 0 || e.Next > 0
}
func (e Emitter) Done() bool {
	return e.Bursts >= e.Count
}

type Number interface {
	constraints.Integer | constraints.Float
}

type Range[V Number] struct {
	Min, Max V
}

func (r Range[V]) Get() V {
	return V(float32(r.Max-r.Min)*rand.Float32() + float32(r.Min))
}

func NewSingle[V Number](value V) Range[V] {
	return Range[V]{Min: value, Max: value}
}

func NewRange[V Number](min, maxExclusive V) Range[V] {
	return Range[V]{Min: min, Max: maxExclusive}
}
