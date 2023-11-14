package ui

import "github.com/axe/axe-go/pkg/util"

type Visual interface {
	Init(b *Base)
	Update(b *Base, update Update) Dirty
	Visualize(b *Base, bounds Bounds, ctx *RenderContext, out *VertexBuffers)
	PreferredSize(b *Base, ctx *RenderContext, maxWidth float32) Coord
}

var _ Visual = VisualFilled{}
var _ Visual = VisualBordered{}
var _ Visual = VisualShadow{}
var _ Visual = VisualFrame{}
var _ Visual = &VisualText{}

type VisualFilled struct {
	Shape Shape
}

func (s VisualFilled) Init(b *Base) {
	if s.Shape != nil {
		s.Shape.Init()
	}
}

func (s VisualFilled) Update(b *Base, update Update) Dirty {
	return DirtyNone
}

func (s VisualFilled) Visualize(b *Base, bounds Bounds, ctx *RenderContext, out *VertexBuffers) {
	points := coalesceShape(s.Shape, b.Shape).Shapify(bounds, ctx)
	center := Coord{}
	for _, p := range points {
		center.X += p.X
		center.Y += p.Y
	}
	scale := 1.0 / float32(len(points))
	center.X *= scale
	center.Y *= scale
	// TODO improvement to check for one of the points existing on a line between two others and producing a single triangle
	last := len(points) - 1
	prev := last
	buffer := out.Buffer()
	for next := 0; next <= last; next++ {
		prevPoint := points[prev]
		nextPoint := points[next]
		buffer.AddIndexed(
			Vertex{X: prevPoint.X, Y: prevPoint.Y},
			Vertex{X: nextPoint.X, Y: nextPoint.Y},
			Vertex{X: center.X, Y: center.Y},
		)
		prev = next
	}
}

func (s VisualFilled) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32) Coord {
	return Coord{}
}

type VisualBorderScale struct {
	NormalX, NormalY, Weight float32
	Spread                   bool
}

type VisualBordered struct {
	Width      float32
	OuterColor Colorable
	InnerColor Colorable
	Scales     []VisualBorderScale
	Shape      Shape
}

func (s VisualBordered) Init(b *Base) {
	if s.Shape != nil {
		s.Shape.Init()
	}
}

func (s VisualBordered) Update(b *Base, update Update) Dirty {
	return DirtyNone
}

func (s VisualBordered) Visualize(b *Base, bounds Bounds, ctx *RenderContext, out *VertexBuffers) {
	inner := coalesceShape(s.Shape, b.Shape).Shapify(bounds, ctx)
	outer := make([]Coord, len(inner))
	last := len(inner) - 1
	i0 := last - 1
	i1 := last
	w := s.Width * 0.5
	scaleCount := len(s.Scales)
	scaleIndex := 0
	for i2 := 0; i2 <= last; i2++ {
		p0 := inner[i0]
		p1 := inner[i1]
		p2 := inner[i2]
		n1dx, n1dy, _ := NormalBetween(p0, p1)
		n2dx, n2dy, _ := NormalBetween(p2, p1)
		nx := (n1dy + -n2dy)
		ny := (-n1dx + n2dx)
		for scaleIndex = 0; scaleIndex < scaleCount; scaleIndex++ {
			scale := s.Scales[scaleIndex]
			dot := scale.NormalX*nx + scale.NormalY*ny
			if dot > 0 {
				if scale.Spread {
					dotScaled := dot * scale.Weight
					nx *= dotScaled
					ny *= dotScaled
				} else {
					dotScaled := dot*scale.Weight - 1
					nx += dotScaled * scale.NormalX
					ny += dotScaled * scale.NormalY
				}
			}
		}
		outer[i1].X = nx*w + p1.X
		outer[i1].Y = ny*w + p1.Y
		i0 = i1
		i1 = i2
	}

	outerColor, hasOuterColor := GetColor(s.OuterColor, b)
	innerColor, hasInnerColor := GetColor(s.InnerColor, b)

	prev := last
	buffer := out.Buffer()
	for next := 0; next <= last; next++ {
		prevOuter := outer[prev]
		nextOuter := outer[next]
		prevInner := inner[prev]
		nextInner := inner[next]
		prev = next

		if Collinear(prevOuter, nextOuter, nextInner) && Collinear(nextInner, prevInner, prevOuter) {
			continue
		}

		buffer.AddQuad(
			Vertex{X: prevOuter.X, Y: prevOuter.Y, Color: outerColor, HasColor: hasOuterColor},
			Vertex{X: nextOuter.X, Y: nextOuter.Y, Color: outerColor, HasColor: hasOuterColor},
			Vertex{X: nextInner.X, Y: nextInner.Y, Color: innerColor, HasColor: hasInnerColor},
			Vertex{X: prevInner.X, Y: prevInner.Y, Color: innerColor, HasColor: hasInnerColor},
		)
	}
}

func (s VisualBordered) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32) Coord {
	return Coord{}
}

type VisualShadow struct {
	Shape      Shape
	Blur       AmountBounds
	Offsets    AmountBounds
	AlwaysFill bool
}

func (s VisualShadow) Init(b *Base) {
	if s.Shape != nil {
		s.Shape.Init()
	}
}

func (s VisualShadow) Update(b *Base, update Update) Dirty {
	return DirtyNone
}

func (s VisualShadow) Visualize(b *Base, bounds Bounds, ctx *RenderContext, out *VertexBuffers) {
	blur := s.Blur.GetBounds(ctx.AmountContext)
	offsets := s.Offsets.GetBounds(ctx.AmountContext)
	shape := coalesceShape(s.Shape, b.Shape)
	offsetBounds := bounds
	if !offsets.IsZero() {
		offsetBounds = offsetBounds.Add(offsets)
	}
	points := shape.Shapify(offsetBounds, ctx)

	innerShape := ShapePolygon{Points: points, Absolute: true}

	bordered := VisualBordered{
		Shape:      innerShape,
		Width:      1,
		InnerColor: ColorWhite,
		OuterColor: ColorTransparent,
	}
	if blur.IsUniform() {
		bordered.Width = blur.Left
	} else {
		bordered.Scales = []VisualBorderScale{
			{NormalX: -1, NormalY: 0, Weight: blur.Left},
			{NormalX: 0, NormalY: -1, Weight: blur.Top},
			{NormalX: 1, NormalY: 0, Weight: blur.Right},
			{NormalX: 0, NormalY: 1, Weight: blur.Bottom},
		}
	}
	bordered.Visualize(b, offsetBounds, ctx, out)

	if s.AlwaysFill || !bounds.Contains(offsetBounds) {
		filled := VisualFilled{
			Shape: innerShape,
		}
		start := NewVertexIterator(out)
		filled.Visualize(b, offsetBounds, ctx, out)
		for start.HasNext() {
			start.Next().AddColor(ColorWhite)
		}
	}
}

func (s VisualShadow) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32) Coord {
	return Coord{}
}

type VisualFrame struct {
	Sizes   AmountBounds
	Tile    []Tile
	Columns int
}

func (r VisualFrame) Init(b *Base) {

}

func (r VisualFrame) Update(b *Base, update Update) Dirty {
	return DirtyNone
}

func (r VisualFrame) Visualize(b *Base, bounds Bounds, ctx *RenderContext, out *VertexBuffers) {
	sizes := r.Sizes.GetBounds(ctx.AmountContext)
	axisX := []float32{bounds.Left, bounds.Left + sizes.Left, bounds.Right - sizes.Right, bounds.Right}
	axisY := []float32{bounds.Top, bounds.Top + sizes.Top, bounds.Bottom - sizes.Bottom, bounds.Bottom}

	columns := r.Columns
	if columns == 0 {
		switch len(r.Tile) {
		case 4, 8:
			columns = 2
		case 6, 9:
			columns = 3
		default:
			columns = len(r.Tile)
		}
	}

	buffer := out.Buffer()
	for i, tile := range r.Tile {
		indexX := i % columns
		indexY := i / columns
		buffer.AddQuad(
			Vertex{X: axisX[indexX+0], Y: axisY[indexY+0], Coord: tile.Coord(0, 0), HasCoord: true},
			Vertex{X: axisX[indexX+1], Y: axisY[indexY+0], Coord: tile.Coord(1, 0), HasCoord: true},
			Vertex{X: axisX[indexX+1], Y: axisY[indexY+1], Coord: tile.Coord(1, 1), HasCoord: true},
			Vertex{X: axisX[indexX+0], Y: axisY[indexY+1], Coord: tile.Coord(0, 1), HasCoord: true},
		)
	}
}

func (s VisualFrame) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32) Coord {
	return Coord{}
}

type VisualText struct {
	Paragraphs       Paragraphs
	VisibleThreshold *GlyphVisibility
	Clip             bool
	Animation        TextAnimation
	AnimationFactory TextAnimationFactory

	dirty          Dirty
	rendered       RenderedText
	renderedBounds Bounds
	animationTime  float32
}

func (s *VisualText) Init(b *Base) {
}

func (s *VisualText) Update(b *Base, update Update) Dirty {
	dirty := s.dirty
	s.dirty.Clear()
	if s.Animation != nil {
		if s.animationTime == 0 {
			s.Init(b)
		}
		dirty.Add(s.Animation.Update(b, s.animationTime, update))
		if s.Animation.IsDone(b, s.animationTime) {
			s.Animation = nil
		} else {
			s.animationTime += float32(update.DeltaTime.Seconds())
		}
	}
	return dirty
}

func (s *VisualText) Animate(factory TextAnimationFactory) *VisualText {
	s.animationTime = 0
	s.AnimationFactory = factory
	return s
}

func (s *VisualText) Clipped() *VisualText {
	partial := GlyphVisibilityPartial
	s.Clip = true
	s.VisibleThreshold = &partial
	return s.Dirty()
}

func (s *VisualText) SetText(text string) *VisualText {
	s.Paragraphs = MustTextToParagraphs(text)
	s.Animation = nil
	return s.Dirty()
}

func (s *VisualText) SetParagraphs(para Paragraphs) *VisualText {
	s.Paragraphs = para
	s.Animation = nil
	return s.Dirty()
}

func (s *VisualText) Dirty() *VisualText {
	s.renderedBounds = Bounds{}
	s.dirty = DirtyVisual
	return s
}

func (s *VisualText) Visibility(threshold GlyphVisibility) *VisualText {
	s.VisibleThreshold = &threshold
	return s.Dirty()
}

func (s *VisualText) WillClip() bool {
	return s.Clip && (s.VisibleThreshold == nil || *s.VisibleThreshold != GlyphVisibilityVisible)
}

func (s *VisualText) Visualize(b *Base, bounds Bounds, ctx *RenderContext, out *VertexBuffers) {
	if s.renderedBounds != bounds {
		s.Paragraphs.MaxWidth, s.Paragraphs.MaxHeight = bounds.Dimensions()
		s.rendered = s.Paragraphs.Render(ctx, b)
		s.rendered.Translate(bounds.Left, bounds.Top)
		s.renderedBounds = bounds

		if s.VisibleThreshold != nil {
			s.rendered.UpdateVisibility(bounds)
		}

		if s.animationTime == 0 && s.Animation == nil && s.AnimationFactory != nil {
			s.Animation = s.AnimationFactory.GetAnimation(&s.rendered)
		}
	}
	clipBounds := Bounds{}
	if s.WillClip() {
		clipBounds = bounds
	}
	out.ClipMaybe(clipBounds, func(vb *VertexBuffers) {
		buffer := vb.Buffer()
		if s.Animation != nil {
			s.Animation.Render(b, s.animationTime, bounds, ctx, buffer)
		} else {
			for _, g := range s.rendered.Glyphs {
				if g.Empty || (s.VisibleThreshold != nil && g.Visibility > *s.VisibleThreshold) {
					continue
				}
				buffer.AddQuad(g.Quad()...)
			}
		}
	})
}

var TextMeasureErrorAmount float32 = 5

func (s VisualText) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32) Coord {
	if maxWidth <= 0 {
		maxWidth = s.Paragraphs.MinWidth(ctx)
	}
	existingMaxWidth := s.Paragraphs.MaxWidth
	newMaxWidth := maxWidth
	if existingMaxWidth > 0 {
		newMaxWidth = util.Min(newMaxWidth, existingMaxWidth)
	}
	s.Paragraphs.MaxWidth = newMaxWidth
	size := s.Paragraphs.Measure(ctx)
	size.X += TextMeasureErrorAmount
	s.Paragraphs.MaxWidth = existingMaxWidth
	return size
}

func (s VisualText) Rendered() RenderedText {
	return s.rendered
}
func (s VisualText) RenderedBounds() Bounds {
	return s.renderedBounds
}

func coalesceShape(a, b Shape) Shape {
	if a == nil {
		a = b
	}
	return a
}
