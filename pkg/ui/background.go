package ui

type Background interface {
	Init(b *Base, init Init)
	Update(b *Base, update Update) Dirty
	Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex)
}

var _ Background = BackgroundColor{}
var _ Background = BackgroundLinearGradient{}
var _ Background = BackgroundRadialGradient{}
var _ Background = BackgroundImage{}

type BackgroundColor struct {
	Color Color
}

func (bc BackgroundColor) Init(b *Base, init Init)             {}
func (bc BackgroundColor) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bc BackgroundColor) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	out.AddColor(bc.Color.R, bc.Color.G, bc.Color.B, bc.Color.A)
}

type BackgroundLinearGradient struct {
	StartColor Color
	Start      Coord
	EndColor   Color
	End        Coord
}

func (bc BackgroundLinearGradient) Init(b *Base, init Init)             {}
func (bc BackgroundLinearGradient) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bg BackgroundLinearGradient) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	dx := bg.End.X - bg.Start.X
	dy := bg.End.Y - bg.Start.Y
	lenSq := dx*dx + dy*dy
	px := bounds.Dx(out.X) - bg.Start.X
	py := bounds.Dy(out.Y) - bg.Start.Y
	delta := Clamp(((dx*px)+(dy*py))/lenSq, 0, 1)
	out.AddColor(
		Lerp(bg.StartColor.R, bg.EndColor.R, delta),
		Lerp(bg.StartColor.G, bg.EndColor.G, delta),
		Lerp(bg.StartColor.B, bg.EndColor.B, delta),
		Lerp(bg.StartColor.A, bg.EndColor.A, delta),
	)
}

type BackgroundImage struct {
	Tile Tile
}

func (bi BackgroundImage) Init(b *Base, init Init)             {}
func (bc BackgroundImage) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bi BackgroundImage) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	out.SetCoord(
		bi.Tile.Texture,
		Lerp(bi.Tile.Coords.Left, bi.Tile.Coords.Right, bounds.Dx(out.X)),
		Lerp(bi.Tile.Coords.Top, bi.Tile.Coords.Bottom, bounds.Dy(out.Y)),
	)
}

type BackgroundRadialGradient struct {
	InnerColor Color
	OuterColor Color
	Radius     AmountPoint
	Offset     AmountPoint
}

func (bg BackgroundRadialGradient) Init(b *Base, init Init)             {}
func (bc BackgroundRadialGradient) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bg BackgroundRadialGradient) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	offX, offY := bg.Offset.Get(ctx.AmountContext)
	cx := bounds.Dx(0.5) + offX
	cy := bounds.Dy(0.5) + offY
	rx, ry := bg.Radius.Get(ctx.AmountContext)
	dx := out.X - cx
	dy := out.Y - cy
	len := Length(dx, dy)
	nx := dx / len
	ny := dy / len
	olen := Length(nx*rx, ny*ry)
	delta := Clamp(olen/len, 0, 1)

	out.AddColor(
		Lerp(bg.InnerColor.R, bg.OuterColor.R, delta),
		Lerp(bg.InnerColor.G, bg.OuterColor.G, delta),
		Lerp(bg.InnerColor.B, bg.OuterColor.B, delta),
		Lerp(bg.InnerColor.A, bg.OuterColor.A, delta),
	)
}
