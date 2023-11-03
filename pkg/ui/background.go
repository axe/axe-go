package ui

type Background interface {
	Init(init Init)
	Update(update Update) Dirty
	Backgroundify(b Bounds, ctx *RenderContext, out *Vertex)
}

var _ Background = BackgroundColor{}
var _ Background = BackgroundLinearGradient{}
var _ Background = BackgroundRadialGradient{}
var _ Background = BackgroundImage{}

type BackgroundColor struct {
	Color Color
}

func (bc BackgroundColor) Init(init Init)             {}
func (bc BackgroundColor) Update(update Update) Dirty { return DirtyNone }
func (bc BackgroundColor) Backgroundify(b Bounds, ctx *RenderContext, out *Vertex) {
	out.AddColor(bc.Color.R, bc.Color.G, bc.Color.B, bc.Color.A)
}

type BackgroundLinearGradient struct {
	StartColor Color
	Start      Coord
	EndColor   Color
	End        Coord
}

func (bc BackgroundLinearGradient) Init(init Init)             {}
func (bc BackgroundLinearGradient) Update(update Update) Dirty { return DirtyNone }
func (bg BackgroundLinearGradient) Backgroundify(b Bounds, ctx *RenderContext, out *Vertex) {
	dx := bg.End.X - bg.Start.X
	dy := bg.End.Y - bg.Start.Y
	lenSq := dx*dx + dy*dy
	px := b.Dx(out.X) - bg.Start.X
	py := b.Dy(out.Y) - bg.Start.Y
	delta := clamp(((dx*px)+(dy*py))/lenSq, 0, 1)
	out.AddColor(
		lerp(bg.StartColor.R, bg.EndColor.R, delta),
		lerp(bg.StartColor.G, bg.EndColor.G, delta),
		lerp(bg.StartColor.B, bg.EndColor.B, delta),
		lerp(bg.StartColor.A, bg.EndColor.A, delta),
	)
}

type BackgroundImage struct {
	Tile Tile
}

func (bi BackgroundImage) Init(init Init)             {}
func (bc BackgroundImage) Update(update Update) Dirty { return DirtyNone }
func (bi BackgroundImage) Backgroundify(b Bounds, ctx *RenderContext, out *Vertex) {
	out.SetCoord(
		bi.Tile.Texture,
		lerp(bi.Tile.Coords.Left, bi.Tile.Coords.Right, b.Dx(out.X)),
		lerp(bi.Tile.Coords.Top, bi.Tile.Coords.Bottom, b.Dy(out.Y)),
	)
}

type BackgroundRadialGradient struct {
	InnerColor Color
	OuterColor Color
	Radius     AmountPoint
	Offset     AmountPoint
}

func (bg BackgroundRadialGradient) Init(init Init)             {}
func (bc BackgroundRadialGradient) Update(update Update) Dirty { return DirtyNone }
func (bg BackgroundRadialGradient) Backgroundify(b Bounds, ctx *RenderContext, out *Vertex) {
	offX, offY := bg.Offset.Get(ctx.AmountContext)
	cx := b.Dx(0.5) + offX
	cy := b.Dy(0.5) + offY
	rx, ry := bg.Radius.Get(ctx.AmountContext)
	dx := out.X - cx
	dy := out.Y - cy
	len := length(dx, dy)
	nx := dx / len
	ny := dy / len
	olen := length(nx*rx, ny*ry)
	delta := clamp(olen/len, 0, 1)

	out.AddColor(
		lerp(bg.InnerColor.R, bg.OuterColor.R, delta),
		lerp(bg.InnerColor.G, bg.OuterColor.G, delta),
		lerp(bg.InnerColor.B, bg.OuterColor.B, delta),
		lerp(bg.InnerColor.A, bg.OuterColor.A, delta),
	)
}
