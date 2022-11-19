package axe

import (
	"math/rand"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

type Attr[V any] interface {
	// distance between this and value
	Distance(value V) float32
	// squared distance between this and value
	DistanceSq(value V) float32
	// distance between this and zero
	Length() float32
	// squared distance between this and zero
	LengthSq() float32
	// dot product between this and value
	Dot(value V) float32
	// the number of float components that make up this attribute
	Components() int
	// gets the float component at the given index
	GetComponent(index int) float32
	// out[index] = value
	SetComponent(index int, value float32, out *V)
	// out[all] = value
	SetComponents(value float32, out *V)
	// out = this
	Set(out *V)
	// out = this + value
	Add(addend V, out *V)
	// out = this + value * scale
	AddScaled(value V, scale float32, out *V)
	// out = this - value
	Sub(subtrahend V, out *V)
	// out = this * value
	Mul(factor V, out *V)
	// out = this / value
	Div(factor V, out *V)
	// out = this * scale
	Scale(scale float32, out *V)
	// out = (end - start) * delta + start
	Interpolate(start V, end V, delta float32, out *V)
}

func div[D Numeric](a D, b D) D {
	if b == 0 {
		return 0
	}
	return a / b
}

type Range[A Attr[A]] struct {
	Min A
	Max A
}

func (r Range[A]) At(delta float32) A {
	var at A
	at.Interpolate(r.Min, r.Max, delta, &at)
	return at
}

func (r Range[A]) Random(rnd rand.Rand) A {
	return r.At(rnd.Float32())
}

type NumericRange[D Numeric] struct {
	Min D
	Max D
}

func (r NumericRange[D]) At(delta float32) D {
	return D(float32(r.Max-r.Min) * delta)
}

func (r NumericRange[D]) Random(rnd rand.Rand) D {
	return r.At(rnd.Float32())
}

type Shape[A Attr[A]] interface {
	Finite() bool
	Radius() float32
	Sign(position A, point A) int
	Distance(position A, point A) float32
	DistanceSq(position A, point A) float32
	Normal(position A, point A, out *A) bool
	Raytrace(position A, point A, direction A) bool
	Bounds(position A, bounds Bounds[A]) bool
}

type Bounds[A Attr[A]] struct {
	Min       A
	Max       A
	Thickness float32
}

type Bounds2f = Bounds[Vec2f]

var _ Shape[Vec2f] = &Bounds2f{}

func (b Bounds[A]) Finite() bool {
	return true
}
func (b Bounds[A]) Radius() float32 {
	return b.Max.Distance(b.Min)*0.5 + float32(b.Thickness)
}
func (b Bounds[A]) Sign(position A, point A) int {
	return 0
}
func (b Bounds[A]) Distance(position A, point A) float32 {
	return 0
}
func (b Bounds[A]) DistanceSq(position A, point A) float32 {
	return 0
}
func (b Bounds[A]) Normal(position A, point A, out *A) bool {
	return true
}
func (b Bounds[A]) Raytrace(position A, point A, direction A) bool {
	return true
}
func (b Bounds[A]) Bounds(position A, bounds Bounds[A]) bool {
	return true
}
