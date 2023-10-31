package ui

import "math"

type Outline interface {
	Init(init Init)
	Outlinify(b Bounds, ctx AmountContext) []Coord
}

var _ Outline = OutlineRectangle{}
var _ Outline = OutlineRounded{}
var _ Outline = OutlineSharpen{}

type OutlineRectangle struct{}

func (o OutlineRectangle) Init(init Init) {}
func (o OutlineRectangle) Outlinify(b Bounds, ctx AmountContext) []Coord {
	return []Coord{
		{X: b.Left, Y: b.Top},
		{X: b.Right, Y: b.Top},
		{X: b.Right, Y: b.Bottom},
		{X: b.Left, Y: b.Bottom},
	}
}

type OutlineRounded struct {
	Radius       AmountCorners
	UnitToPoints float32
}

var OutlineRoundedAngles = [][]float32{{math.Pi, math.Pi * 0.5}, {math.Pi * 0.5, 0}, {math.Pi * 2, math.Pi * 1.5}, {math.Pi * 1.5, math.Pi}}
var OutlineRoundedPlacements = [][]float32{{0, 0}, {1, 0}, {1, 1}, {0, 1}} // 0=1, 1=-1     *2 (0,2) -1 (-1,1)

func (o OutlineRounded) Init(init Init) {}
func (o OutlineRounded) Outlinify(b Bounds, ctx AmountContext) []Coord {
	amounts := []Amount{o.Radius.TopLeft, o.Radius.TopRight, o.Radius.BottomRight, o.Radius.BottomLeft}
	coords := make([]Coord, 0, 16)
	for i := 0; i < 4; i++ {
		amount := amounts[i]
		angles := OutlineRoundedAngles[i]
		placements := OutlineRoundedPlacements[i]
		radiusW := amount.Get(ctx.ForWidth())
		radiusH := amount.Get(ctx.ForHeight())
		points := int((radiusW+radiusH)*0.5*o.UnitToPoints) + 1
		originX := lerp(b.Left, b.Right, placements[0]) - (radiusW * ((placements[0] * 2) - 1))
		originY := lerp(b.Top, b.Bottom, placements[1]) - (radiusH * ((placements[1] * 2) - 1))

		for i := 0; i <= points; i++ {
			delta := float32(i) / float32(points)
			angle := lerp(angles[0], angles[1], delta)
			coords = append(coords, Coord{
				X: float32(math.Cos(float64(angle)))*radiusW + originX,
				Y: float32(-math.Sin(float64(angle)))*radiusH + originY,
			})
		}
	}
	return coords
}

type OutlineSharpen struct {
	Outline Outline
	Times   int
}

func (o OutlineSharpen) Init(init Init) {
	o.Outline.Init(init)
}
func (o OutlineSharpen) Outlinify(b Bounds, ctx AmountContext) []Coord {
	points := o.Outline.Outlinify(b, ctx)
	times := o.Times + 1
	sharpened := make([]Coord, len(points)*times)
	last := len(points) - 1
	prev := last
	for next := 0; next <= last; next++ {
		prevPoint := points[prev]
		nextPoint := points[next]
		for i := 0; i < times; i++ {
			delta := float32(i) / float32(times)
			sharp := &sharpened[next*times+i]
			sharp.X = lerp(prevPoint.X, nextPoint.X, delta)
			sharp.Y = lerp(prevPoint.Y, nextPoint.Y, delta)
		}
	}
	return sharpened
}
