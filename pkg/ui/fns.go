package ui

import (
	"math"
	"reflect"
)

func Lerp(s, e, d float32) float32 {
	return (e-s)*d + s
}

func Delta(s, e, v float32) float32 {
	return (v - s) / (e - s)
}

func Normal(a, b Coord) (nx, ny float32) {
	dx := b.X - a.X
	dy := b.Y - a.Y
	invLength := 1.0 / Length(dx, dy)
	nx = dx * invLength
	ny = dy * invLength
	return
}

func Length(dx, dy float32) float32 {
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func abs(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func equal(a, b float32) bool {
	return abs(a-b) < 0.0001
}

func Collinear(a, b, c Coord) bool {
	return equal((b.Y-a.Y)/(b.X-a.X), (c.Y-b.Y)/(c.X-b.X)) ||
		equal((b.Y-a.Y)*(c.X-b.X), (c.Y-b.Y)*(b.X-a.X))
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Clamp(v, min, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func toPtr(x any) uintptr {
	return reflect.ValueOf(x).Pointer()
}

func InPolygon(polygon []Coord, pt Coord) bool {
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

func Cos(rad float32) float32 {
	return float32(math.Cos(float64(rad)))
}

func Sin(rad float32) float32 {
	return float32(math.Sin(float64(rad)))
}

func CosSin(rad float32) (cos float32, sin float32) {
	rad64 := float64(rad)
	cos = float32(math.Cos(rad64))
	sin = float32(math.Sin(rad64))
	return
}

func Atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func Ease(delta float32, easing func(float32) float32) float32 {
	if easing == nil {
		return delta
	} else {
		return easing(delta)
	}
}
