package axe

import (
	"math"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

type Attr[V any] interface {
	Distance(a V) float32
}

type Vec2[D Numeric] struct {
	X D
	Y D
}

var _ Attr[Vec2[float32]] = &Vec2[float32]{}

func (v Vec2[D]) Distance(a Vec2[D]) float32 {
	dx := v.X - a.X
	dy := v.Y - a.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

type Vec3[D Numeric] struct {
	X D
	Y D
	Z D
}

var _ Attr[Vec3[float32]] = &Vec3[float32]{}

func (v Vec3[D]) Distance(a Vec3[D]) float32 {
	dx := v.X - a.X
	dy := v.Y - a.Y
	dz := v.Z - a.Z
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

type Shape[D Numeric, V Attr[V]] interface {
	Finite() bool
	Radius() float32
	Sign(position V, point V) int
	Distance(position V, point V) float32
	DistanceSq(position V, point V) float32
	Normal(position V, point V, out *V) bool
	Raytrace(position V, point V, direction V) bool
	Bounds(position V, bounds Bounds[D, V]) bool
}

type Bounds[D Numeric, V Attr[V]] struct {
	Min       V
	Max       V
	Thickness D
}

var _ Shape[float32, Vec2[float32]] = &Bounds[float32, Vec2[float32]]{}

func (b Bounds[D, V]) Finite() bool {
	return true
}
func (b Bounds[D, V]) Radius() float32 {
	return b.Max.Distance(b.Min) + float32(b.Thickness)
}
func (b Bounds[D, V]) Sign(position V, point V) int {
	return 0
}
func (b Bounds[D, V]) Distance(position V, point V) float32 {
	return 0
}
func (b Bounds[D, V]) DistanceSq(position V, point V) float32 {
	return 0
}
func (b Bounds[D, V]) Normal(position V, point V, out *V) bool {
	return true
}
func (b Bounds[D, V]) Raytrace(position V, point V, direction V) bool {
	return true
}
func (b Bounds[D, V]) Bounds(position V, bounds Bounds[D, V]) bool {
	return true
}
