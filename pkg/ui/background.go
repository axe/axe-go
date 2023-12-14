package ui

import (
	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/util"
)

type Background interface {
	Init(b *Base)
	Update(b *Base, update Update) Dirty
	Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex)
}

var _ Background = BackgroundColor{}
var _ Background = BackgroundLinearGradient{}
var _ Background = BackgroundRadialGradient{}
var _ Background = BackgroundImage{}

type BackgroundColor struct {
	Color color.Able
}

func (bc BackgroundColor) Init(b *Base)                        {}
func (bc BackgroundColor) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bc BackgroundColor) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	out.AddColor(bc.Color.GetColor(b))
}

type BackgroundLinearGradient struct {
	StartColor color.Able
	Start      gfx.Coord
	EndColor   color.Able
	End        gfx.Coord
}

func (bc BackgroundLinearGradient) Init(b *Base)                        {}
func (bc BackgroundLinearGradient) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bg BackgroundLinearGradient) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	dx := bg.End.X - bg.Start.X
	dy := bg.End.Y - bg.Start.Y
	lenSq := dx*dx + dy*dy
	px := bounds.Dx(out.X) - bg.Start.X
	py := bounds.Dy(out.Y) - bg.Start.Y
	delta := util.Clamp(((dx*px)+(dy*py))/lenSq, 0, 1)
	startColor := bg.StartColor.GetColor(b)
	endColor := bg.EndColor.GetColor(b)

	out.AddColor(startColor.Lerp(endColor, delta))
}

type BackgroundImage struct {
	Tile gfx.Tile
	// TODO instead of stretching, support:
	// TileWidth   Amount
	// TileHeight  Amount
	// AspectRatio float32
}

func (bi BackgroundImage) Init(b *Base)                        {}
func (bc BackgroundImage) Update(b *Base, update Update) Dirty { return DirtyNone }
func (bi BackgroundImage) Backgroundify(b *Base, bounds Bounds, ctx *RenderContext, out *Vertex) {
	out.SetCoord(bi.Tile.Coord(bounds.Dx(out.X), bounds.Dy(out.Y)))
}

type BackgroundRadialGradient struct {
	InnerColor color.Able
	OuterColor color.Able
	Radius     AmountPoint
	Offset     AmountPoint
}

func (bg BackgroundRadialGradient) Init(b *Base)                        {}
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
	delta := util.Clamp(olen/len, 0, 1)
	innerColor := bg.InnerColor.GetColor(b)
	outerColor := bg.OuterColor.GetColor(b)

	out.AddColor(innerColor.Lerp(outerColor, delta))
}
