package ui

import (
	"math"
	"reflect"
)

func lerp(s, e, d float32) float32 {
	return (e-s)*d + s
}

func normal(a, b Coord) (nx, ny float32) {
	dx := b.X - a.X
	dy := b.Y - a.Y
	invLength := 1.0 / length(dx, dy)
	nx = dx * invLength
	ny = dy * invLength
	return
}

func length(dx, dy float32) float32 {
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

func collinear(a, b, c Coord) bool {
	return equal((b.Y-a.Y)/(b.X-a.X), (c.Y-b.Y)/(c.X-b.X)) ||
		equal((b.Y-a.Y)*(c.X-b.X), (c.Y-b.Y)*(b.X-a.X))
}

func clamp(v, min, max float32) float32 {
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

func inPolygon(polygon []Coord, pt Coord) bool {
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

func sliceIndexOf[V comparable](slice []V, value V) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}
	return -1
}

func sliceMove[V any](slice []V, from, to int) {
	if from == to || to < 0 || from < 0 || from >= len(slice) || to >= len(slice) {
		return
	}
	if from < to {
		value := slice[from]
		copy(slice[from:to], slice[from+1:to+1])
		slice[to] = value
	} else {
		value := slice[to]
		copy(slice[to+1:from+1], slice[to:from])
		slice[from] = value
	}
}
