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
	Color Colorable
}

func (bc BackgroundColor) Init(b *Base, init Init)             {}
func (bc BackgroundColor) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bc BackgroundColor) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	out.AddColor(bc.Color.GetColor(b))
}

type BackgroundLinearGradient struct {
	StartColor Colorable
	Start      Coord
	EndColor   Colorable
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
	startColor := bg.StartColor.GetColor(b)
	endColor := bg.EndColor.GetColor(b)

	out.AddColor(startColor.Lerp(endColor, delta))
}

type BackgroundImage struct {
	Tile Tile
	// TODO instead of stretching, support:
	// TileWidth   Amount
	// TileHeight  Amount
	// AspectRatio float32
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
	InnerColor Colorable
	OuterColor Colorable
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
	innerColor := bg.InnerColor.GetColor(b)
	outerColor := bg.OuterColor.GetColor(b)

	out.AddColor(innerColor.Lerp(outerColor, delta))
}
