package ui

import "fmt"

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
func (b Bounds) Union(a Bounds) Bounds {
	if b.IsZero() {
		return a
	}
	return Bounds{
		Left:   min(a.Left, b.Left),
		Top:    min(a.Top, b.Top),
		Right:  max(a.Right, b.Right),
		Bottom: max(a.Bottom, b.Bottom),
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
		b.Left = min(x, b.Left)
		b.Right = max(x, b.Right)
		b.Top = min(y, b.Top)
		b.Bottom = max(y, b.Bottom)
	}
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
		Start:      Coord{X: x0, Y: y0},
		StartDelta: 0,
		End:        Coord{X: x1, Y: y1},
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
			clipSide := side0
			if side1 > side0 {
				clipSide = side1
			}
			var x, y, delta float32
			if (clipSide & BoundsSideBottom) != 0 {
				delta = (b.Bottom - y0) / (y1 - y0)
				x = x0 + (x1-x0)*delta
				y = b.Bottom
			} else if (clipSide & BoundsSideTop) != 0 {
				delta = (b.Top - y0) / (y1 - y0)
				x = x0 + (x1-x0)*delta
				y = b.Top
			} else if (clipSide & BoundsSideRight) != 0 {
				delta = (b.Right - x0) / (x1 - x0)
				y = y0 + (y1-y0)*delta
				x = b.Right
			} else if (clipSide & BoundsSideLeft) != 0 {
				delta = (b.Left - x0) / (x1 - x0)
				y = y0 + (y1-y0)*delta
				x = b.Left
			}
			if clipSide == side0 {
				line.StartDelta = delta
				line.Start.X = x
				line.Start.Y = y
				side0 = b.Side(x, y)
			} else {
				line.EndDelta = delta
				line.End.X = x
				line.End.Y = y
				side1 = b.Side(x, y)
			}
		}
	}

	return line
}

type ClippedPoint struct {
	Start, End int
	Delta      float32
	X, Y       float32
}

func (p *ClippedPoint) Set(start, end int, delta, x, y float32) {
	p.Start = start
	p.End = end
	p.Delta = delta
	p.X = x
	p.Y = y
}

func (b Bounds) ClipTriangle(x0, y0, x1, y1, x2, y2 float32, out [6]ClippedPoint) int {
	line0 := b.ClipLine(x0, y0, x1, y1)
	line1 := b.ClipLine(x1, y1, x2, y2)
	line2 := b.ClipLine(x2, y2, x0, y0)

	if !line0.Inside && !line1.Inside && !line2.Inside {
		return 0
	}

	if line0.Inside && line1.Inside && line2.Inside {
		out[0].Set(0, 1, 0, x0, y0)
		out[1].Set(1, 2, 0, x1, y1)
		out[2].Set(2, 0, 0, x2, y2)

		return 3
	}

	i := 0
	out[i].Set(0, 1, line0.StartDelta, line0.Start.X, line0.Start.Y)
	i++
	if line0.EndDelta < 1 {
		out[i].Set(0, 1, line0.EndDelta, line0.End.X, line0.End.Y)
		i++
	}
	out[i].Set(1, 2, line1.StartDelta, line1.Start.X, line1.Start.Y)
	i++
	if line1.EndDelta < 1 {
		out[i].Set(1, 2, line1.EndDelta, line1.End.X, line1.End.Y)
		i++
	}
	out[i].Set(2, 0, line2.StartDelta, line2.Start.X, line2.Start.Y)
	i++
	if line2.EndDelta < 1 {
		out[i].Set(2, 0, line2.EndDelta, line2.End.X, line2.End.Y)
		i++
	}
	return i
}
