package ui

import (
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/util"
)

type Placement struct {
	Left   Anchor
	Right  Anchor
	Top    Anchor
	Bottom Anchor
}

func Maximized() Placement {
	return Placement{}.Maximize()
}

func Centered(width float32, height float32) Placement {
	return Placement{}.Center(width, height)
}

func Absolute(left, top, width, height float32) Placement {
	p := Placement{}
	p.Left.Set(left, 0)
	p.Right.Set(left+width, 0)
	p.Top.Set(top, 0)
	p.Bottom.Set(top+height, 0)
	return p
}

func MaximizeOffset(left, top, right, bottom float32) Placement {
	p := Placement{}
	p.Left.Set(left, 0)
	p.Top.Set(top, 0)
	p.Right.Set(-right, 1)
	p.Bottom.Set(-bottom, 1)
	return p
}

func (p *Placement) Init(defaultPlacement Placement) {
	if !p.Defined() {
		*p = defaultPlacement
	}
}

func (p Placement) WithWidth(width float32) Placement {
	delta := p.Left.Delta
	if delta == p.Right.Delta {
		reference := util.Lerp(p.Left.Base, p.Right.Base, delta)
		p.Left.Base = reference - width*delta
		p.Right.Base = reference + width*(1-delta)
	}
	return p
}

func (p Placement) WithWidthRelative(width, parentWidth float32) Placement {
	if p.Left.Delta == p.Right.Delta {
		return p.WithWidth(width)
	}
	currentWidth := p.GetWidth(parentWidth)
	adjustX := (currentWidth - width) * 0.5
	p.Left.Base += adjustX
	p.Right.Base -= adjustX
	return p
}

func (p Placement) WithHeight(height float32) Placement {
	delta := p.Top.Delta
	if delta == p.Bottom.Delta {
		reference := util.Lerp(p.Top.Base, p.Bottom.Base, delta)
		p.Top.Base = reference - height*delta
		p.Bottom.Base = reference + height*(1-delta)
	}
	return p
}

func (p Placement) WithHeightRelative(height, parentHeight float32) Placement {
	if p.Top.Delta == p.Bottom.Delta {
		return p.WithHeight(height)
	}
	currentHeight := p.GetHeight(parentHeight)
	adjustY := (currentHeight - height) * 0.5
	p.Top.Base += adjustY
	p.Bottom.Base -= adjustY
	return p
}

func (p Placement) WithSize(width, height float32) Placement {
	return p.WithWidth(width).WithHeight(height)
}

func (p Placement) WithSizeRelative(width, height, parentWidth, parentHeight float32) Placement {
	return p.WithWidthRelative(width, parentWidth).WithHeightRelative(height, parentHeight)
}

func (p Placement) WithLeft(left, parentWidth float32) Placement {
	p.Left.Base = left - p.Left.Delta*parentWidth
	return p
}

func (p Placement) WithTop(top, parentHeight float32) Placement {
	p.Top.Base = top - p.Top.Delta*parentHeight
	return p
}

func (p Placement) WithRight(right, parentWidth float32) Placement {
	p.Right.Base = right - p.Right.Delta*parentWidth
	return p
}

func (p Placement) WithBottom(bottom, parentHeight float32) Placement {
	p.Bottom.Base = bottom - p.Bottom.Delta*parentHeight
	return p
}

func (p Placement) Defined() bool {
	return !p.Left.IsZero() || !p.Right.IsZero() || !p.Top.IsZero() || !p.Bottom.IsZero()
}

func (p Placement) Shrink(amount float32) Placement {
	p.Left.Base += amount
	p.Top.Base += amount
	p.Right.Base -= amount
	p.Bottom.Base -= amount
	return p
}

func (p Placement) Grow(amount float32) Placement {
	return p.Shrink(-amount)
}

func (p Placement) Shift(dx, dy float32) Placement {
	p.Left.Base += dx
	p.Top.Base += dy
	p.Right.Base += dx
	p.Bottom.Base += dy
	return p
}

func (p Placement) FitInside(width, height float32, keepSize bool) Placement {
	b := p.GetBounds(width, height)
	if b.Right > width {
		over := b.Right - width
		p.Right.Base -= over
		if keepSize {
			p.Left.Base -= over
		}
		b.Left -= over
		b.Right = width
	}
	if b.Left < 0 {
		p.Left.Base -= b.Left
		if keepSize {
			p.Right.Base -= b.Left
		}
	}
	if b.Bottom > height {
		over := b.Bottom - height
		p.Bottom.Base -= over
		if keepSize {
			p.Top.Base -= over
		}
		b.Bottom = height
		b.Top -= over
	}
	if b.Top < 0 {
		p.Top.Base -= b.Top
		if keepSize {
			p.Bottom.Base -= b.Top
		}
	}
	return p
}

func (p Placement) Padded(padding Bounds) Placement {
	p.Left.Base += padding.Left
	p.Right.Base -= padding.Right
	p.Top.Base += padding.Top
	p.Bottom.Base -= padding.Bottom
	return p
}

func (p Placement) Relative(leftAnchor float32, topAnchor float32, rightAnchor float32, bottomAnchor float32) Placement {
	p.Left.Set(0, leftAnchor)
	p.Top.Set(0, topAnchor)
	p.Right.Set(0, rightAnchor)
	p.Bottom.Set(0, bottomAnchor)
	return p
}

func (p Placement) Maximize() Placement {
	return p.Relative(0, 0, 1, 1)
}

func (p Placement) IsMaximized() bool {
	return p.Left.Is(0, 0) && p.Top.Is(0, 0) && p.Right.Is(0, 1) && p.Bottom.Is(0, 1)
}

func (p Placement) Attach(dx float32, dy float32, width float32, height float32) Placement {
	p.Left.Set(width*-dx, dx)
	p.Right.Set(width*(1-dx), dx)
	p.Top.Set(height*-dy, dy)
	p.Bottom.Set(height*(1-dy), dy)
	return p
}

func (p Placement) Center(width float32, height float32) Placement {
	return p.Attach(0.5, 0.5, width, height)
}

func (p Placement) TopFixedHeight(topOffset float32, height float32, leftOffset float32, rightOffset float32) Placement {
	p.Left.Set(leftOffset, 0)
	p.Right.Set(-rightOffset, 1)
	p.Top.Set(topOffset, 0)
	p.Bottom.Set(topOffset+height, 0)
	return p
}

func (p Placement) BottomFixedHeight(bottomOffset float32, height float32, leftOffset float32, rightOffset float32) Placement {
	p.Left.Set(leftOffset, 0)
	p.Right.Set(-rightOffset, 1)
	p.Top.Set(-bottomOffset-height, 1)
	p.Bottom.Set(-bottomOffset, 1)
	return p
}

func (p Placement) LeftFixedWidth(leftOffset float32, width float32, topOffset float32, bottomOffset float32) Placement {
	p.Left.Set(leftOffset, 0)
	p.Right.Set(leftOffset+width, 0)
	p.Top.Set(topOffset, 0)
	p.Bottom.Set(-bottomOffset, 1)
	return p
}

func (p Placement) RightFixedWidth(rightOffset float32, width float32, topOffset float32, bottomOffset float32) Placement {
	p.Left.Set(-rightOffset-width, 1)
	p.Right.Set(-rightOffset, 1)
	p.Top.Set(topOffset, 0)
	p.Bottom.Set(-bottomOffset, 1)
	return p
}

func (p Placement) GetBoundsIn(parent Bounds) Bounds {
	w := parent.Width()
	h := parent.Height()
	return Bounds{
		Left:   p.Left.Get(w) + parent.Left,
		Top:    p.Top.Get(h) + parent.Top,
		Right:  p.Right.Get(w) + parent.Left,
		Bottom: p.Bottom.Get(h) + parent.Top,
	}
}

func (p Placement) GetBounds(parentWidth float32, parentHeight float32) Bounds {
	return Bounds{
		Left:   p.Left.Get(parentWidth),
		Top:    p.Top.Get(parentHeight),
		Right:  p.Right.Get(parentWidth),
		Bottom: p.Bottom.Get(parentHeight),
	}
}

func (p *Placement) GetBoundsi(parentWidth float32, parentHeight float32) geom.Bounds2i {
	return geom.Bounds2i{
		Min: geom.Vec2i{
			X: int32(p.Left.Get(parentWidth)),
			Y: int32(p.Top.Get(parentHeight)),
		},
		Max: geom.Vec2i{
			X: int32(p.Right.Get(parentWidth)),
			Y: int32(p.Bottom.Get(parentHeight)),
		},
	}
}

func (p Placement) Contains(x float32, y float32, parentWidth float32, parentHeight float32) bool {
	return !(x < p.Left.Get(parentWidth) ||
		y < p.Top.Get(parentHeight) ||
		x > p.Right.Get(parentWidth) ||
		y > p.Bottom.Get(parentHeight))
}

func (p Placement) GetWidth(parentWidth float32) float32 {
	return p.GetMinWidth(parentWidth, 0)
}

func (p Placement) GetMinWidth(parentWidth float32, minWidth float32) float32 {
	return util.Max(minWidth, p.Right.Get(parentWidth)-p.Left.Get(parentWidth))
}

func (p Placement) GetHeight(parentHeight float32) float32 {
	return p.GetMinHeight(parentHeight, 0)
}

func (p Placement) GetMinHeight(parentHeight float32, minHeight float32) float32 {
	return util.Max(minHeight, p.Top.Get(parentHeight)-p.Bottom.Get(parentHeight))
}

func (p Placement) PreferredWidth() float32 {
	if p.Left.Delta == p.Right.Delta {
		return p.Right.Base - p.Left.Base
	}
	return 0
}

func (p Placement) MinParentSize() gfx.Coord {
	return gfx.Coord{
		X: p.MinParentWidth(),
		Y: p.MinParentHeight(),
	}
}

func (p Placement) MinParentWidth() float32 {
	return p.ParentWidth(p.PreferredWidth())
}

func (p Placement) PreferredHeight() float32 {
	if p.Top.Delta == p.Bottom.Delta {
		return p.Bottom.Base - p.Top.Base
	}
	return 0
}

func (p Placement) MinParentHeight() float32 {
	return p.ParentHeight(p.PreferredHeight())
}

func (p Placement) Padding() gfx.Coord {
	return gfx.Coord{
		X: p.ParentWidth(0),
		Y: p.ParentHeight(0),
	}
}

func (p Placement) ParentSize(minWidth, minHeight float32) gfx.Coord {
	return gfx.Coord{
		X: p.ParentWidth(minWidth),
		Y: p.ParentHeight(minHeight),
	}
}

func (p Placement) ParentWidth(minWidth float32) float32 {
	w := minWidth
	if p.Left.Delta == 0 {
		w += p.Left.Base
	}
	if p.Right.Delta == 1 {
		w -= p.Right.Base
	}
	return w
}

func (p Placement) ParentHeight(minHeight float32) float32 {
	h := minHeight
	if p.Top.Delta == 0 {
		h += p.Top.Base
	}
	if p.Bottom.Delta == 1 {
		h -= p.Bottom.Base
	}
	return h
}
