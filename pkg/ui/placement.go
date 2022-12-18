package ui

import (
	"math"

	"github.com/axe/axe-go/pkg/geom"
)

type Placement struct {
	Left   Anchor
	Right  Anchor
	Top    Anchor
	Bottom Anchor
}

func Maximized() Placement {
	p := Placement{}
	p.Maximize()
	return p
}

func Centered(width float32, height float32) Placement {
	p := Placement{}
	p.Center(width, height)
	return p
}

func (p *Placement) Relative(leftAnchor float32, topAnchor float32, rightAnchor float32, bottomAnchor float32) {
	p.Left.Set(0, leftAnchor)
	p.Top.Set(0, topAnchor)
	p.Right.Set(0, rightAnchor)
	p.Bottom.Set(0, bottomAnchor)
}

func (p *Placement) Maximize() {
	p.Relative(0, 0, 1, 1)
}

func (p *Placement) Attach(dx float32, dy float32, width float32, height float32) {
	p.Left.Set(width*-dx, dx)
	p.Right.Set(width*(1-dx), dx)
	p.Bottom.Set(height*-dy, dy)
	p.Top.Set(height*(1-dy), dy)
}

func (p *Placement) Center(width float32, height float32) {
	p.Attach(0.5, 0.5, width, height)
}

func (p *Placement) TopFixedHeight(topOffset float32, height float32, leftOffset float32, rightOffset float32) {
	p.Left.Set(leftOffset, 0)
	p.Right.Set(rightOffset, 1)
	p.Top.Set(topOffset, 0)
	p.Bottom.Set(topOffset+height, 0)
}

func (p *Placement) BottomFixedHeight(bottomOffset float32, height float32, leftOffset float32, rightOffset float32) {
	p.Left.Set(leftOffset, 0)
	p.Right.Set(rightOffset, 1)
	p.Top.Set(bottomOffset+height, 0)
	p.Bottom.Set(bottomOffset, 0)
}

func (p *Placement) LeftFixedWidth(leftOffset float32, width float32, topOffset float32, bottomOffset float32) {
	p.Left.Set(leftOffset, 0)
	p.Right.Set(leftOffset+width, 0)
	p.Top.Set(topOffset, 0)
	p.Bottom.Set(bottomOffset, 1)
}

func (p *Placement) RightFixedWidth(rightOffset float32, width float32, topOffset float32, bottomOffset float32) {
	p.Left.Set(rightOffset+width, 1)
	p.Right.Set(rightOffset, 1)
	p.Top.Set(topOffset, 0)
	p.Bottom.Set(bottomOffset, 1)
}

func (p *Placement) GetBoundsf(parentWidth float32, parentHeight float32) geom.Bounds2f {
	return geom.Bounds2f{
		Min: geom.Vec2f{
			X: p.Left.Get(parentWidth),
			Y: p.Top.Get(parentHeight),
		},
		Max: geom.Vec2f{
			X: p.Right.Get(parentWidth),
			Y: p.Bottom.Get(parentHeight),
		},
	}
}

func (p *Placement) GetBoundsi(parentWidth float32, parentHeight float32) geom.Bounds2i {
	return geom.Bounds2i{
		Min: geom.Vec2i{
			X: int(p.Left.Get(parentWidth)),
			Y: int(p.Top.Get(parentHeight)),
		},
		Max: geom.Vec2i{
			X: int(p.Right.Get(parentWidth)),
			Y: int(p.Bottom.Get(parentHeight)),
		},
	}
}

func (p Placement) Contains(x float32, y float32, parentWidth float32, parentHeight float32) bool {
	return !(x < p.Left.Get(parentWidth) ||
		y > p.Top.Get(parentHeight) ||
		x > p.Right.Get(parentWidth) ||
		y < p.Bottom.Get(parentHeight))
}

func (p Placement) GetWidth(parentWidth float32) float32 {
	return p.GetMinWidth(parentWidth, 0)
}

func (p Placement) GetMinWidth(parentWidth float32, minWidth float32) float32 {
	return float32(math.Max(float64(minWidth), float64(p.Right.Get(parentWidth)-p.Left.Get(parentWidth))))
}

func (p Placement) GetHeight(parentHeight float32) float32 {
	return p.GetMinHeight(parentHeight, 0)
}

func (p Placement) GetMinHeight(parentHeight float32, minHeight float32) float32 {
	return float32(math.Max(float64(minHeight), float64(p.Top.Get(parentHeight)-p.Bottom.Get(parentHeight))))
}

func (p Placement) PreferredWidth() float32 {
	if p.Left.Delta == p.Right.Delta {
		return p.Right.Base - p.Left.Base
	}
	return 0
}

func (p Placement) MinParentWidth() float32 {
	w := p.PreferredWidth()
	if p.Left.Delta == 0 {
		w += p.Left.Base
	}
	if p.Right.Delta == 1 {
		w += p.Right.Base
	}
	return w
}

func (p Placement) PreferredHeight() float32 {
	if p.Top.Delta == p.Bottom.Delta {
		return p.Bottom.Base - p.Top.Base
	}
	return 0
}

func (p Placement) MinParentHeight() float32 {
	h := p.PreferredHeight()
	if p.Top.Delta == 0 {
		h += p.Top.Base
	}
	if p.Bottom.Delta == 1 {
		h += p.Bottom.Base
	}
	return h
}
