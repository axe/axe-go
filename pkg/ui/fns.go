package ui

import "math"

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
