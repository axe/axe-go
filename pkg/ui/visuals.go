package ui

type Visual interface {
	Init(init Init)
	Update(update Update)
	Visualize(b Bounds, ctx AmountContext, out *UIVertexBuffer)
}

var _ Visual = VisualFilled{}
var _ Visual = VisualBordered{}
var _ Visual = VisualFrame{}
var _ Visual = &VisualText{}

type VisualFilled struct {
	Outline Outline
}

func (s VisualFilled) Init(init Init) {
	s.Outline.Init(init)
}

func (s VisualFilled) Update(update Update) {

}

func (s VisualFilled) Visualize(b Bounds, ctx AmountContext, out *UIVertexBuffer) {
	points := s.Outline.Outlinify(b, ctx)
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
	for next := 0; next <= last; next++ {
		prevPoint := points[prev]
		nextPoint := points[next]
		i := out.Add(
			UIVertex{X: prevPoint.X, Y: prevPoint.Y},
			UIVertex{X: nextPoint.X, Y: nextPoint.Y},
			UIVertex{X: center.X, Y: center.Y},
		)
		out.AddIndices(i, i+1, i+2)
		prev = next
	}
}

type VisualBordered struct {
	Width         float32
	OuterColor    Color
	HasOuterColor bool
	InnerColor    Color
	HasInnerColor bool
	Outline       Outline
}

func (s VisualBordered) Init(init Init) {
	s.Outline.Init(init)
}

func (s VisualBordered) Update(update Update) {

}

func (s VisualBordered) Visualize(b Bounds, ctx AmountContext, out *UIVertexBuffer) {
	inner := s.Outline.Outlinify(b, ctx)
	outer := make([]Coord, len(inner))
	last := len(inner) - 1
	i0 := last - 1
	i1 := last
	hw := s.Width * 0.5
	for i2 := 0; i2 <= last; i2++ {
		p0 := inner[i0]
		p1 := inner[i1]
		p2 := inner[i2]
		n1dx, n1dy := normal(p0, p1)
		n2dx, n2dy := normal(p2, p1)
		nx := n1dy + -n2dy
		ny := -n1dx + n2dx
		outer[i1].X = nx*hw + p1.X
		outer[i1].Y = ny*hw + p1.Y
		i0 = i1
		i1 = i2
	}

	prev := last
	for next := 0; next <= last; next++ {
		prevOuter := outer[prev]
		nextOuter := outer[next]
		prevInner := inner[prev]
		nextInner := inner[next]
		i := out.Add(
			UIVertex{X: prevOuter.X, Y: prevOuter.Y, Color: s.OuterColor, HasColor: s.HasOuterColor},
			UIVertex{X: nextOuter.X, Y: nextOuter.Y, Color: s.OuterColor, HasColor: s.HasOuterColor},
			UIVertex{X: nextInner.X, Y: nextInner.Y, Color: s.InnerColor, HasColor: s.HasInnerColor},
			UIVertex{X: prevInner.X, Y: prevInner.Y, Color: s.InnerColor, HasColor: s.HasInnerColor},
		)
		out.AddQuad(i)
		prev = next
	}
}

type VisualFrame struct {
	Sizes AmountBounds
	Tile  []Tile
}

func (r VisualFrame) Init(init Init) {

}

func (r VisualFrame) Update(update Update) {

}

func (r VisualFrame) Visualize(b Bounds, ctx AmountContext, out *UIVertexBuffer) {
	sizes := r.Sizes.GetBounds(ctx)
	axisX := []float32{b.Left, b.Left + sizes.Left, b.Right - sizes.Right, b.Right}
	axisY := []float32{b.Top, b.Top + sizes.Top, b.Bottom - sizes.Bottom, b.Bottom}
	for i, tile := range r.Tile {
		indexX := i % 3
		indexY := i / 3
		k := out.Add(
			UIVertex{X: axisX[indexX+0], Y: axisY[indexY+0], Coord: tile.Coord(0, 0), HasCoord: true},
			UIVertex{X: axisX[indexX+1], Y: axisY[indexY+0], Coord: tile.Coord(1, 0), HasCoord: true},
			UIVertex{X: axisX[indexX+1], Y: axisY[indexY+1], Coord: tile.Coord(1, 1), HasCoord: true},
			UIVertex{X: axisX[indexX+0], Y: axisY[indexY+1], Coord: tile.Coord(0, 1), HasCoord: true},
		)
		out.AddQuad(k)
	}
}

type VisualText struct {
	Glyphs GlyphBlocks

	theme          *Theme
	rendered       RenderedGlyphBlock
	renderedBounds Bounds
}

func (s *VisualText) Init(init Init) {
	s.theme = init.Theme
}

func (s *VisualText) Update(update Update) {

}

func (s *VisualText) Visualize(b Bounds, ctx AmountContext, out *UIVertexBuffer) {
	if s.renderedBounds != b {
		s.Glyphs.MaxWidth, s.Glyphs.MaxHeight = b.Dimensions()
		s.rendered = s.Glyphs.Render(s.theme, ctx)
		s.rendered.Translate(b.Left, b.Top)
		s.renderedBounds = b
	}
	for _, g := range s.rendered.Glyphs {
		i := out.Add(
			UIVertex{X: g.Bounds.Left, Y: g.Bounds.Top, Coord: g.Coord(0, 0), HasCoord: true, Color: g.Color, HasColor: true},
			UIVertex{X: g.Bounds.Right, Y: g.Bounds.Top, Coord: g.Coord(1, 0), HasCoord: true, Color: g.Color, HasColor: true},
			UIVertex{X: g.Bounds.Right, Y: g.Bounds.Bottom, Coord: g.Coord(1, 1), HasCoord: true, Color: g.Color, HasColor: true},
			UIVertex{X: g.Bounds.Left, Y: g.Bounds.Bottom, Coord: g.Coord(0, 1), HasCoord: true, Color: g.Color, HasColor: true},
		)
		out.AddQuad(i)
	}
}
