package ui

import (
	"fmt"

	"github.com/axe/axe-go/pkg/util"
)

type Bounds struct {
	Left, Top, Right, Bottom float32
}

func NewBounds(left, top, right, bottom float32) Bounds {
	return Bounds{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
	}
}

func (b Bounds) IsZero() bool {
	return !(b.Left != 0 || b.Right != 0 || b.Top != 0 || b.Bottom != 0)
}
func (b Bounds) IsEmpty() bool {
	return b.Left >= b.Right && b.Top >= b.Bottom
}
func (b Bounds) IsUniform() bool {
	return b.Left == b.Right && b.Left == b.Top && b.Left == b.Bottom
}
func (b Bounds) IsPositive() bool {
	return b.Left > 0 || b.Right > 0 || b.Top > 0 || b.Bottom > 0
}
func (b Bounds) IsNegative() bool {
	return b.Left < 0 || b.Right < 0 || b.Top < 0 || b.Bottom < 0
}
func (b Bounds) Width() float32 {
	return b.Right - b.Left
}
func (b Bounds) Height() float32 {
	return b.Bottom - b.Top
}
func (b Bounds) Size() Coord {
	return Coord{X: b.Width(), Y: b.Height()}
}
func (b Bounds) Dimensions() (float32, float32) {
	return b.Width(), b.Height()
}
func (b Bounds) Center() (float32, float32) {
	return (b.Left + b.Right) * 0.5, (b.Top + b.Bottom) * 0.5
}
func (b *Bounds) Translate(x, y float32) {
	b.Left += x
	b.Right += x
	b.Top += y
	b.Bottom += y
}
func (b Bounds) Dx(x float32) float32 {
	return (x - b.Left) / b.Width()
}
func (b Bounds) Dy(y float32) float32 {
	return (y - b.Top) / b.Height()
}
func (b Bounds) Delta(x, y float32) (float32, float32) {
	return b.Dx(x), b.Dy(y)
}
func (b Bounds) Lerpx(dx float32) float32 {
	return b.Width()*dx + b.Left
}
func (b Bounds) Lerpy(dy float32) float32 {
	return b.Height()*dy + b.Top
}
func (b Bounds) Lerp(x, y float32) (float32, float32) {
	return b.Lerpx(x), b.Lerpy(y)
}
func (b Bounds) Inside(x, y float32) bool {
	return !(x < b.Left || x > b.Right || y < b.Top || y > b.Bottom)
}
func (b Bounds) InsideCoord(c Coord) bool {
	return !(c.X < b.Left || c.X > b.Right || c.Y < b.Top || c.Y > b.Bottom)
}
func (b Bounds) Expand(a Bounds) Bounds {
	return Bounds{
		Left:   b.Left - a.Left,
		Top:    b.Top - a.Top,
		Right:  b.Right + a.Right,
		Bottom: b.Bottom + a.Bottom,
	}
}
func (b Bounds) Add(a Bounds) Bounds {
	return Bounds{
		Left:   b.Left + a.Left,
		Top:    b.Top + a.Top,
		Right:  b.Right + a.Right,
		Bottom: b.Bottom + a.Bottom,
	}
}
func (b Bounds) Sub(a Bounds) Bounds {
	return Bounds{
		Left:   b.Left - a.Left,
		Top:    b.Top - a.Top,
		Right:  b.Right - a.Right,
		Bottom: b.Bottom - a.Bottom,
	}
}
func (b Bounds) Union(a Bounds) Bounds {
	if b.IsZero() {
		return a
	}
	return Bounds{
		Left:   util.Min(a.Left, b.Left),
		Top:    util.Min(a.Top, b.Top),
		Right:  util.Max(a.Right, b.Right),
		Bottom: util.Max(a.Bottom, b.Bottom),
	}
}
func (b Bounds) Intersects(a Bounds) bool {
	return !(a.Right < b.Left || a.Left > b.Right || a.Bottom < b.Top || a.Top > b.Bottom)
}
func (b Bounds) Contains(a Bounds) bool {
	return !(a.Left < b.Left || a.Top < b.Top || a.Right > b.Right || a.Bottom > b.Bottom)
}
func (b Bounds) Scale(s float32) Bounds {
	return Bounds{
		Left:   b.Left * s,
		Top:    b.Top * s,
		Right:  b.Right * s,
		Bottom: b.Bottom * s,
	}
}
func (b Bounds) ClipCoord(c Coord) Coord {
	c.X = util.Clamp(c.X, b.Left, b.Right)
	c.Y = util.Clamp(c.Y, b.Top, b.Bottom)
	return c
}
func (b *Bounds) Clear() {
	b.Left = 0
	b.Top = 0
	b.Right = 0
	b.Bottom = 0
}
func (b *Bounds) Include(x, y float32) {
	if b.IsZero() {
		b.Left = x
		b.Right = x
		b.Top = y
		b.Bottom = y
	} else {
		b.Left = util.Min(x, b.Left)
		b.Right = util.Max(x, b.Right)
		b.Top = util.Min(y, b.Top)
		b.Bottom = util.Max(y, b.Bottom)
	}
}
func (b Bounds) Closest(x, y float32) (float32, float32) {
	return util.Clamp(x, b.Left, b.Right), util.Clamp(y, b.Top, b.Bottom)
}
func (b Bounds) ClosestCoord(c Coord) Coord {
	x, y := b.Closest(c.X, c.Y)
	return Coord{X: x, Y: y}
}
func (b Bounds) String() string {
	return fmt.Sprintf("{L:%f T:%f R:%f B:%f W:%f H:%f}", b.Left, b.Top, b.Right, b.Bottom, b.Right-b.Left, b.Bottom-b.Top)
}

type BoundsSide int

const (
	BoundsSideLeft BoundsSide = 1 << iota
	BoundsSideTop
	BoundsSideRight
	BoundsSideBottom
)

func (b Bounds) Side(x, y float32) BoundsSide {
	var sides BoundsSide
	if x < b.Left {
		sides |= BoundsSideLeft
	} else if x > b.Right {
		sides |= BoundsSideRight
	}
	if y < b.Top {
		sides |= BoundsSideTop
	} else if y > b.Bottom {
		sides |= BoundsSideBottom
	}
	return sides
}

type ClippedLine struct {
	Start      Coord
	StartDelta float32
	End        Coord
	EndDelta   float32
	Inside     bool
}

func (b Bounds) ClipLine(x0, y0, x1, y1 float32) ClippedLine {
	side0 := b.Side(x0, y0)
	side1 := b.Side(x1, y1)
	line := ClippedLine{
		StartDelta: 0,
		EndDelta:   1,
		Inside:     false,
	}

	for {
		if (side0 | side1) == 0 {
			line.Inside = true
			break
		} else if (side0 & side1) != 0 {
			break
		} else {
			clipSide := util.Max(side0, side1)
			var x, y, delta float32
			if (clipSide & BoundsSideBottom) != 0 {
				delta = util.Delta(y0, y1, b.Bottom)
				x = util.Lerp(x0, x1, delta)
				y = b.Bottom
			} else if (clipSide & BoundsSideTop) != 0 {
				delta = util.Delta(y0, y1, b.Top)
				x = util.Lerp(x0, x1, delta)
				y = b.Top
			} else if (clipSide & BoundsSideRight) != 0 {
				delta = util.Delta(x0, x1, b.Right)
				y = util.Lerp(y0, y1, delta)
				x = b.Right
			} else if (clipSide & BoundsSideLeft) != 0 {
				delta = util.Delta(x0, x1, b.Left)
				y = util.Lerp(y0, y1, delta)
				x = b.Left
			}
			if clipSide == side0 {
				line.StartDelta = delta
				x0 = x
				y0 = y
				side0 = b.Side(x, y)
			} else {
				line.EndDelta = delta
				x1 = x
				y1 = y
				side1 = b.Side(x, y)
			}
		}
	}

	line.Start.Set(x0, y0)
	line.End.Set(x1, y1)

	return line
}
