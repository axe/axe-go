package ui

import (
	"reflect"

	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/util"
)

// Computes the normal between an origin and a point and returns the length as well.
func NormalBetween(origin, point gfx.Coord) (nx, ny, length float32) {
	return Normal(point.X-origin.X, point.Y-origin.Y)
}

// Computes the normal of a vector and returns the length as well.
func Normal(vx, vy float32) (nx, ny, length float32) {
	length = Length(vx, vy)
	invLength := util.Div(1.0, length)
	nx = vx * invLength
	ny = vy * invLength
	return
}

// Returns if the three points lie on the same line.
func Collinear(a, b, c gfx.Coord) bool {
	return util.Equal[float32]((b.Y-a.Y)/(b.X-a.X), (c.Y-b.Y)/(c.X-b.X)) ||
		util.Equal[float32]((b.Y-a.Y)*(c.X-b.X), (c.Y-b.Y)*(b.X-a.X))
}

// Computes the length of the given vector/difference.
func Length(dx, dy float32) float32 {
	return util.Sqrt(dx*dx + dy*dy)
}

// Computes the length of the given vector/difference.
func LengthSq(dx, dy float32) float32 {
	return dx*dx + dy*dy
}

func InPolygon(polygon []gfx.Coord, pt gfx.Coord) bool {
	in := false
	n := len(polygon)
	i := 0
	j := n - 1
	for i < n {
		jp := polygon[j]
		ip := polygon[i]
		if ((ip.Y > pt.Y) != (jp.Y > pt.Y)) && (pt.X < (jp.X-ip.X)*(pt.Y-ip.Y)/(jp.Y-ip.Y)+ip.X) {
			in = !in
		}
		j = i
		i++
	}
	return in
}

func Ease(delta float32, easing func(float32) float32) float32 {
	if easing == nil {
		return delta
	} else {
		return easing(delta)
	}
}

func toPtr(x any) uintptr {
	return reflect.ValueOf(x).Pointer()
}
