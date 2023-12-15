package fx

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

type EmitterType interface {
	Create() Emitter
}

type Emitter interface {
	Update(dt float32, rnd *rand.Rand) (size int, over float32)
	Done() bool
}

type BurstEmitterType struct {
	Delay     Range[float32]
	Frequency Range[float32]
	Size      Range[int]
	Count     Range[int]
}

func (et BurstEmitterType) Create(rnd *rand.Rand) BurstEmitter {
	return BurstEmitter{
		Delay:     et.Delay.Get(rnd),
		Frequency: et.Frequency,
		Size:      et.Size,
		Count:     et.Count.Get(rnd),
	}
}

type BurstEmitter struct {
	Delay     float32
	Frequency Range[float32]
	Size      Range[int]
	Count     int

	Time   float32
	Next   float32
	Bursts int
}

func (e *BurstEmitter) Update(dt float32, rnd *rand.Rand) (size int, over float32) {
	if e.Done() {
		return
	}
	if !e.Started() {
		e.Next = e.Delay
	}
	e.Time += dt
	if e.Time >= e.Next {
		over = e.Next - e.Time
		size = e.Size.Get(rnd)
		e.Time -= e.Next
		e.Bursts++
		e.Next = e.Frequency.Get(rnd)
	}
	return
}
func (e BurstEmitter) Started() bool {
	return e.Bursts > 0 || e.Next > 0
}
func (e BurstEmitter) Done() bool {
	return e.Bursts >= e.Count
}

type Number interface {
	constraints.Integer | constraints.Float
}

type Range[V Number] struct {
	Min, Max V
}

func (r Range[V]) Get(rnd *rand.Rand) V {
	return V(float32(r.Max-r.Min)*rnd.Float32() + float32(r.Min))
}

func NewSingle[V Number](value V) Range[V] {
	return Range[V]{Min: value, Max: value}
}

func NewRange[V Number](min, maxExclusive V) Range[V] {
	return Range[V]{Min: min, Max: maxExclusive}
}
