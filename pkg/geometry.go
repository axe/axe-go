package axe

import (
	"math"
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

type Vec2[D Numeric] struct {
	X D
	Y D
}

type Vec2f = Vec2[float32]
type Vec2i = Vec2[int]

var _ Attr[Vec2f] = &Vec2f{}

func (v Vec2[D]) Distance(value Vec2[D]) float32 {
	return float32(math.Sqrt(float64(v.DistanceSq(value))))
}
func (v Vec2[D]) DistanceSq(a Vec2[D]) float32 {
	dx := v.X - a.X
	dy := v.Y - a.Y
	return float32(dx*dx + dy*dy)
}
func (v Vec2[D]) Dot(a Vec2[D]) float32 {
	dx := v.X * a.X
	dy := v.Y * a.Y
	return float32(dx + dy)
}
func (v Vec2[D]) Components() int { return 2 }
func (v Vec2[D]) GetComponent(index int) float32 {
	switch index {
	case 0:
		return float32(v.X)
	case 1:
		return float32(v.Y)
	}
	return 0
}
func (v Vec2[D]) SetComponent(index int, value float32, out *Vec2[D]) {
	switch index {
	case 0:
		out.X = D(value)
	case 1:
		out.Y = D(value)
	}
}
func (v Vec2[D]) SetComponents(value float32, out *Vec2[D]) {
	d := D(value)
	out.X = d
	out.Y = d
}
func (v Vec2[D]) Scale(amount float32, out *Vec2[D]) {
	out.X = D(float32(v.X) * amount)
	out.Y = D(float32(v.Y) * amount)
}
func (v Vec2[D]) AddScaled(value Vec2[D], scale float32, out *Vec2[D]) {
	out.X = v.X + D(float32(value.X)*scale)
	out.Y = v.Y + D(float32(value.Y)*scale)
}
func (v Vec2[D]) Add(value Vec2[D], out *Vec2[D]) {
	out.X = v.X + value.X
	out.Y = v.Y + value.Y
}
func (v Vec2[D]) Sub(value Vec2[D], out *Vec2[D]) {
	out.X = v.X - value.X
	out.Y = v.Y - value.Y
}
func (v Vec2[D]) Mul(factor Vec2[D], out *Vec2[D]) {
	out.X = v.X * factor.X
	out.Y = v.Y * factor.Y
}
func (v Vec2[D]) Div(factor Vec2[D], out *Vec2[D]) {
	out.X = div(v.X, factor.X)
	out.Y = div(v.Y, factor.Y)
}
func (v Vec2[D]) Interpolate(start Vec2[D], end Vec2[D], delta float32, out *Vec2[D]) {
	out.X = D(float32(end.X-start.X)*delta) + start.X
	out.Y = D(float32(end.Y-start.Y)*delta) + start.Y
}

type Vec3[D Numeric] struct {
	X D
	Y D
	Z D
}

type Vec3f = Vec3[float32]
type Vec3i = Vec3[int]

var _ Attr[Vec3f] = &Vec3f{}

func (v Vec3[D]) Distance(value Vec3[D]) float32 {
	return float32(math.Sqrt(float64(v.DistanceSq(value))))
}
func (v Vec3[D]) DistanceSq(value Vec3[D]) float32 {
	dx := v.X - value.X
	dy := v.Y - value.Y
	dz := v.Z - value.Z
	return float32(dx*dx + dy*dy + dz*dz)
}
func (v Vec3[D]) Dot(a Vec3[D]) float32 {
	dx := v.X * a.X
	dy := v.Y * a.Y
	dz := v.Z * a.Z
	return float32(dx + dy + dz)
}
func (v Vec3[D]) Components() int { return 3 }
func (v Vec3[D]) GetComponent(index int) float32 {
	switch index {
	case 0:
		return float32(v.X)
	case 1:
		return float32(v.Y)
	case 2:
		return float32(v.Z)
	}
	return 0
}
func (v Vec3[D]) SetComponent(index int, value float32, out *Vec3[D]) {
	switch index {
	case 0:
		out.X = D(value)
	case 1:
		out.Y = D(value)
	case 2:
		out.Z = D(value)
	}
}
func (v Vec3[D]) SetComponents(value float32, out *Vec3[D]) {
	d := D(value)
	out.X = d
	out.Y = d
	out.Z = d
}
func (v Vec3[D]) Scale(amount float32, out *Vec3[D]) {
	out.X = D(float32(v.X) * amount)
	out.Y = D(float32(v.Y) * amount)
	out.Z = D(float32(v.Z) * amount)
}
func (v Vec3[D]) AddScaled(value Vec3[D], scale float32, out *Vec3[D]) {
	out.X = v.X + D(float32(value.X)*scale)
	out.Y = v.Y + D(float32(value.Y)*scale)
	out.Z = v.Z + D(float32(value.Z)*scale)
}
func (v Vec3[D]) Add(value Vec3[D], out *Vec3[D]) {
	out.X = v.X + value.X
	out.Y = v.Y + value.Y
	out.Z = v.Z + value.Z
}
func (v Vec3[D]) Sub(value Vec3[D], out *Vec3[D]) {
	out.X = v.X - value.X
	out.Y = v.Y - value.Y
	out.Z = v.Z - value.Z
}
func (v Vec3[D]) Mul(value Vec3[D], out *Vec3[D]) {
	out.X = v.X * value.X
	out.Y = v.Y * value.Y
	out.Z = v.Z * value.Z
}
func (v Vec3[D]) Div(value Vec3[D], out *Vec3[D]) {
	out.X = div(v.X, value.X)
	out.Y = div(v.Y, value.Y)
	out.Z = div(v.Z, value.Z)
}
func (v Vec3[D]) Interpolate(start Vec3[D], end Vec3[D], delta float32, out *Vec3[D]) {
	out.X = D(float32(end.X-start.X)*delta) + start.X
	out.Y = D(float32(end.Y-start.Y)*delta) + start.Y
	out.Z = D(float32(end.Z-start.Z)*delta) + start.Z
}

type Vec4[D Numeric] struct {
	X D
	Y D
	Z D
	W D
}

type Vec4f = Vec4[float32]
type Vec4i = Vec4[int]

var _ Attr[Vec4f] = &Vec4f{}

func (v Vec4[D]) Distance(value Vec4[D]) float32 {
	return float32(math.Sqrt(float64(v.DistanceSq(value))))
}
func (v Vec4[D]) DistanceSq(value Vec4[D]) float32 {
	dx := v.X - value.X
	dy := v.Y - value.Y
	dz := v.Z - value.Z
	dw := v.Z - value.Z
	return float32(dx*dx + dy*dy + dz*dz + dw*dw)
}
func (v Vec4[D]) Dot(a Vec4[D]) float32 {
	dx := v.X * a.X
	dy := v.Y * a.Y
	dz := v.Z * a.Z
	dw := v.W * a.W
	return float32(dx + dy + dz + dw)
}
func (v Vec4[D]) Components() int { return 4 }
func (v Vec4[D]) GetComponent(index int) float32 {
	switch index {
	case 0:
		return float32(v.X)
	case 1:
		return float32(v.Y)
	case 2:
		return float32(v.Z)
	case 3:
		return float32(v.W)
	}
	return 0
}
func (v Vec4[D]) SetComponent(index int, value float32, out *Vec4[D]) {
	switch index {
	case 0:
		out.X = D(value)
	case 1:
		out.Y = D(value)
	case 2:
		out.Z = D(value)
	case 3:
		out.W = D(value)
	}
}
func (v Vec4[D]) SetComponents(value float32, out *Vec4[D]) {
	d := D(value)
	out.X = d
	out.Y = d
	out.Z = d
	out.W = d
}
func (v Vec4[D]) Scale(amount float32, out *Vec4[D]) {
	out.X = D(float32(v.X) * amount)
	out.Y = D(float32(v.Y) * amount)
	out.Z = D(float32(v.Z) * amount)
	out.W = D(float32(v.W) * amount)
}
func (v Vec4[D]) AddScaled(value Vec4[D], scale float32, out *Vec4[D]) {
	out.X = v.X + D(float32(value.X)*scale)
	out.Y = v.Y + D(float32(value.Y)*scale)
	out.Z = v.Z + D(float32(value.Z)*scale)
	out.W = v.W + D(float32(value.W)*scale)
}
func (v Vec4[D]) Add(value Vec4[D], out *Vec4[D]) {
	out.X = v.X + value.X
	out.Y = v.Y + value.Y
	out.Z = v.Z + value.Z
	out.W = v.W + value.W
}
func (v Vec4[D]) Sub(value Vec4[D], out *Vec4[D]) {
	out.X = v.X - value.X
	out.Y = v.Y - value.Y
	out.Z = v.Z - value.Z
	out.W = v.W - value.W
}
func (v Vec4[D]) Mul(value Vec4[D], out *Vec4[D]) {
	out.X = v.X * value.X
	out.Y = v.Y * value.Y
	out.Z = v.Z * value.Z
	out.W = v.W * value.W
}
func (v Vec4[D]) Div(value Vec4[D], out *Vec4[D]) {
	out.X = div(v.X, value.X)
	out.Y = div(v.Y, value.Y)
	out.Z = div(v.Z, value.Z)
	out.W = div(v.W, value.W)
}
func (v Vec4[D]) Interpolate(start Vec4[D], end Vec4[D], delta float32, out *Vec4[D]) {
	out.X = D(float32(end.X-start.X)*delta) + start.X
	out.Y = D(float32(end.Y-start.Y)*delta) + start.Y
	out.Z = D(float32(end.Z-start.Z)*delta) + start.Z
	out.W = D(float32(end.W-start.W)*delta) + start.W
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

func AttrDistance[A Attr[A]](start A, end A) float32 {
	sq := float32(0)
	for i := 0; i < start.Components(); i++ {
		d := start.GetComponent(i) - end.GetComponent(i)
		sq += d * d
	}
	return float32(math.Sqrt(float64(sq)))
}

func AttrSet[A Attr[A]](attr *A, value float32) {
	for i := 0; i < (*attr).Components(); i++ {
		(*attr).SetComponent(i, value, attr)
	}
}

func AttrZero[A Attr[A]](attr *A, value float32) {
	AttrSet(attr, 0)
}

func AttrOne[A Attr[A]](attr *A, value float32) {
	AttrSet(attr, 1)
}
