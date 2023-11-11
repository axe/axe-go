package ui

import "math"

type Shape interface {
	Init(init Init)
	Shapify(b Bounds, ctx *RenderContext) []Coord
}

var _ Shape = ShapeRectangle{}
var _ Shape = ShapeRounded{}
var _ Shape = ShapeSharpen{}
var _ Shape = ShapePolygon{}

type ShapeRectangle struct{}

func (o ShapeRectangle) Init(init Init) {}
func (o ShapeRectangle) Shapify(b Bounds, ctx *RenderContext) []Coord {
	return []Coord{
		{X: b.Left, Y: b.Top},
		{X: b.Right, Y: b.Top},
		{X: b.Right, Y: b.Bottom},
		{X: b.Left, Y: b.Bottom},
	}
}

type ShapeRounded struct {
	Radius       AmountCorners
	UnitToPoints float32
}

var ShapeRoundedAngles = [][]float32{{math.Pi, math.Pi * 0.5}, {math.Pi * 0.5, 0}, {math.Pi * 2, math.Pi * 1.5}, {math.Pi * 1.5, math.Pi}}
var ShapeRoundedPlacements = [][]float32{{0, 0}, {1, 0}, {1, 1}, {0, 1}} // 0=1, 1=-1     *2 (0,2) -1 (-1,1)

func (o ShapeRounded) Init(init Init) {}
func (o ShapeRounded) Shapify(b Bounds, ctx *RenderContext) []Coord {
	amounts := []Amount{o.Radius.TopLeft, o.Radius.TopRight, o.Radius.BottomRight, o.Radius.BottomLeft}
	coords := make([]Coord, 0, 16)
	for i := 0; i < 4; i++ {
		amount := amounts[i]
		angles := ShapeRoundedAngles[i]
		placements := ShapeRoundedPlacements[i]
		radiusW := amount.Get(ctx.AmountContext, true)
		radiusH := amount.Get(ctx.AmountContext, false)
		points := int((radiusW+radiusH)*0.5*o.UnitToPoints) + 1
		originX := Lerp(b.Left, b.Right, placements[0]) - (radiusW * ((placements[0] * 2) - 1))
		originY := Lerp(b.Top, b.Bottom, placements[1]) - (radiusH * ((placements[1] * 2) - 1))

		for i := 0; i <= points; i++ {
			delta := float32(i) / float32(points)
			angle := Lerp(angles[0], angles[1], delta)
			coords = append(coords, Coord{
				X: float32(math.Cos(float64(angle)))*radiusW + originX,
				Y: float32(-math.Sin(float64(angle)))*radiusH + originY,
			})
		}
	}
	return coords
}

type ShapeSharpen struct {
	Shape Shape
	Times int
}

func (o ShapeSharpen) Init(init Init) {
	o.Shape.Init(init)
}
func (o ShapeSharpen) Shapify(b Bounds, ctx *RenderContext) []Coord {
	points := o.Shape.Shapify(b, ctx)
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
			sharp.X = Lerp(prevPoint.X, nextPoint.X, delta)
			sharp.Y = Lerp(prevPoint.Y, nextPoint.Y, delta)
		}
	}
	return sharpened
}

type ShapePolygon struct {
	Points   []Coord
	Absolute bool
	Copy     bool
}

func (o ShapePolygon) Init(init Init) {

}
func (o ShapePolygon) Shapify(b Bounds, ctx *RenderContext) []Coord {
	if o.Absolute {
		if o.Copy {
			return append(make([]Coord, 0, len(o.Points)), o.Points...)
		} else {
			return o.Points
		}
	} else {
		n := len(o.Points)
		points := make([]Coord, n)
		for i := 0; i < n; i++ {
			points[i].X = b.Lerpx(o.Points[i].X)
			points[i].Y = b.Lerpy(o.Points[i].Y)
		}
		return points
	}
}
