package shape

import (
	"math"

	"github.com/axe/axe-go/pkg/util"
)

var (
	InfiniteBounds Bounds2 = Bounds2{
		Min: Point2{X: -math.MaxFloat32, Y: -math.MaxFloat32},
		Max: Point2{X: math.MaxFloat32, Y: math.MaxFloat32},
	}
	InfiniteCircle Circle2 = Circle2{Radius: math.MaxFloat32}
)

type HasShape2 interface {
	GetShape() Shape2
}

type Shape2 interface {
	Closest(p Point2, distance *float32, normal *Point2) Point2
	// Translate(t Point2) Shape2
	Bounds() Bounds2
	Circle() Circle2
	// Bounding(p Point2) Shape2
	PolygonPoints(target []Point2) int
	PolygonSize() int
}

var _ Shape2 = Point2{}
var _ Shape2 = Circle2{}
var _ Shape2 = Bounds2{}
var _ Shape2 = Line2{}
var _ Shape2 = Segment2{}
var _ Shape2 = Plane2{}
var _ Shape2 = Ellipse2{}
var _ Shape2 = Polygon2{}
var _ Shape2 = Inverted2{}
var _ Shape2 = Padded2{}

type Point2 struct {
	X, Y float32
}

func (s Point2) Bounds() Bounds2 { return Bounds2{Min: s, Max: s} }
func (s Point2) Circle() Circle2 { return Circle2{Center: s} }
func (s Point2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	if distance != nil || normal != nil {
		n, d := s.NormalTo(p)
		if distance != nil {
			*distance = d
		}
		if normal != nil {
			*normal = n
		}
	}
	return p
}
func (s Point2) Move(dx, dy float32) Point2 { return Point2{X: s.X + dx, Y: s.Y + dy} }
func (s Point2) Add(p Point2) Point2        { return Point2{X: s.X + p.X, Y: s.Y + p.Y} }
func (s Point2) Sub(p Point2) Point2        { return Point2{X: s.X - p.X, Y: s.Y - p.Y} }
func (s Point2) Scale(scale float32) Point2 { return Point2{X: s.X * scale, Y: s.Y * scale} }
func (s Point2) Neg() Point2                { return Point2{X: -s.X, Y: -s.Y} }
func (s Point2) Min(p Point2) Point2        { return Point2{X: util.Min(s.X, p.X), Y: util.Min(s.Y, p.Y)} }
func (s Point2) Max(p Point2) Point2        { return Point2{X: util.Max(s.X, p.X), Y: util.Max(s.Y, p.Y)} }
func (s Point2) Tangent() Point2            { return Point2{X: -s.Y, Y: s.X} }
func (s Point2) Dot(p Point2) float32       { return s.X*p.X + s.Y*p.Y }
func (s Point2) LengthSq() float32          { return s.X*s.X + s.Y*s.Y }
func (s Point2) Length() float32            { return util.Sqrt(s.X*s.X + s.Y*s.Y) }
func (s Point2) NormalTo(p Point2) (Point2, float32) {
	dx := p.X - s.X
	dy := p.Y - s.Y
	sq := dx*dx + dy*dy
	d := sq
	if d != 0 && d != 1 {
		d = util.Sqrt(d)
		invD := 1.0 / d
		dx *= invD
		dy *= invD
	}
	return Point2{X: dx, Y: dy}, d
}
func (s Point2) Lerp(e Point2, d float32) Point2 {
	return Point2{X: (e.X-s.X)*d + s.X, Y: (e.Y-s.Y)*d + s.Y}
}
func (s Point2) PolygonSize() int { return 1 }
func (s Point2) PolygonPoints(target []Point2) int {
	target[0] = s
	return 1
}

type Circle2 struct {
	Center Point2
	Radius float32
}

func (s Circle2) Bounds() Bounds2 {
	return Bounds2{Min: s.Center.Move(-s.Radius, -s.Radius), Max: s.Center.Move(s.Radius, s.Radius)}
}
func (s Circle2) Circle() Circle2 { return s }
func (s Circle2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	n, d := s.Center.NormalTo(p)
	if distance != nil {
		*distance = d - s.Radius
	}
	if normal != nil {
		*normal = n
	}
	r := util.Min(d, s.Radius)
	return Point2{X: n.X * r, Y: n.Y * r}
}
func (s Circle2) PolygonSize() int { return -1 }
func (s Circle2) PolygonPoints(target []Point2) int {
	n := len(target)
	angleIncrease := math.Pi * 2 / float32(n)
	angle := float32(0)
	for i := range target {
		cos, sin := util.CosSin(angle)
		angle += angleIncrease
		target[i] = s.Center.Move(cos*s.Radius, sin*s.Radius)
	}
	return n
}

type Bounds2 struct {
	Min, Max Point2
}

func (s Bounds2) Bounds() Bounds2 { return s }
func (s Bounds2) Circle() Circle2 {
	half := s.Max.Sub(s.Min).Scale(0.5)
	return Circle2{
		Center: s.Min.Add(half),
		Radius: half.Length(),
	}
}
func (s Bounds2) Center() Point2 {
	return s.Max.Add(s.Min).Scale(0.5)
}
func (s Bounds2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	closest := Point2{
		X: util.Clamp(p.X, s.Min.X, s.Max.X),
		Y: util.Clamp(p.Y, s.Min.Y, s.Max.Y),
	}
	if distance != nil || normal != nil {
		if closest.X != p.X || closest.Y != p.Y {
			n, d := closest.NormalTo(p)
			if distance != nil {
				*distance = d
			}
			if normal != nil {
				*normal = n
			}
		} else {
			dt := p.Y - s.Min.Y
			dl := p.X - s.Min.X
			dr := s.Max.X - p.X
			db := s.Max.Y - p.Y
			min := util.Min(dt, util.Min(dl, util.Min(dr, db)))
			n := Point2{}
			switch min {
			case dt:
				n.Y = -1
			case db:
				n.Y = 1
			case dl:
				n.X = -1
			case dr:
				n.X = 1
			}
			if normal != nil {
				*normal = n
			}
			if distance != nil {
				*distance = -min
			}
		}
	}
	return closest
}
func (s Bounds2) Intersects(b Bounds2) bool {
	return !(s.Min.X > b.Max.X || s.Min.Y > b.Max.Y || s.Max.X < b.Min.X || s.Max.Y < b.Min.Y)
}
func (s Bounds2) PolygonSize() int { return 4 }
func (s Bounds2) PolygonPoints(target []Point2) int {
	target[0] = s.Min
	target[1] = Point2{X: s.Max.X, Y: s.Min.Y}
	target[2] = s.Max
	target[3] = Point2{X: s.Min.X, Y: s.Max.Y}
	return 4
}
func (s Bounds2) Union(b Bounds2) Bounds2 {
	return Bounds2{Min: s.Min.Min(b.Min), Max: s.Max.Max(b.Max)}
}
func (s Bounds2) Intersection(b Bounds2) Bounds2 {
	return Bounds2{Min: s.Min.Max(b.Min), Max: s.Max.Min(b.Max)}
}
func (s Bounds2) Area() float32 {
	return (s.Max.X - s.Min.X) * (s.Max.Y - s.Min.Y)
}
func (s Bounds2) Empty() bool {
	return s.Max.X <= s.Min.X || s.Max.Y <= s.Min.Y
}

type Line2 struct {
	Start, End Point2
}

func (s Line2) Bounds() Bounds2 {
	return InfiniteBounds
}
func (s Line2) Circle() Circle2 {
	return InfiniteCircle
}
func (s Line2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	p0 := s.End.Sub(s.Start)
	p1 := p.Sub(s.Start)
	delta := util.Div(p0.Dot(p1), p0.LengthSq())
	closest := s.Start.Lerp(s.End, delta)
	if distance != nil || normal != nil {
		n, d := closest.NormalTo(p)
		if distance != nil {
			*distance = d
		}
		if normal != nil {
			*normal = n
		}
	}
	return closest
}
func (s Line2) PolygonSize() int { return 2 }
func (s Line2) PolygonPoints(target []Point2) int {
	diff := s.End.Sub(s.Start)
	if diff.X == 0 {
		target[0] = Point2{X: s.Start.X, Y: -math.MaxFloat32}
		target[1] = Point2{X: s.Start.X, Y: math.MaxFloat32}
	} else if diff.Y == 0 {
		target[0] = Point2{X: -math.MaxFloat32, Y: s.Start.Y}
		target[1] = Point2{X: math.MaxFloat32, Y: s.Start.Y}
	} else {
		m := float64(diff.Y) / float64(diff.X)
		b := float64(s.Start.Y) - float64(s.Start.X)*m

		s0x := math.MaxFloat32
		s0y := m*s0x + b
		s1y := math.MaxFloat32
		s1x := (s1y - b) / m

		e0x := -math.MaxFloat32
		e0y := m*s0x + b
		e1y := -math.MaxFloat32
		e1x := (s1y - b) / m

		if s0y < s1x {
			target[0].X = float32(s0x)
			target[0].Y = float32(s0y)
		} else {
			target[0].X = float32(s1x)
			target[0].Y = float32(s1y)
		}
		if e0y < e1x {
			target[1].X = float32(e0x)
			target[1].Y = float32(e0y)
		} else {
			target[1].X = float32(e1x)
			target[1].Y = float32(e1y)
		}
	}

	return 2
}

type Segment2 struct {
	Start, End Point2
}

func (s Segment2) Bounds() Bounds2 {
	return Bounds2{
		Min: s.Start.Min(s.End),
		Max: s.Start.Max(s.End),
	}
}
func (s Segment2) Circle() Circle2 {
	half := s.End.Sub(s.Start).Scale(0.5)
	center := half.Add(s.Start)
	return Circle2{Center: center, Radius: half.Length()}
}
func (s Segment2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	p0 := s.End.Sub(s.Start)
	p1 := p.Sub(s.Start)
	delta := util.Div(p0.Dot(p1), p0.LengthSq())
	clampedDelta := util.Clamp(delta, 0, 1)
	closest := s.Start.Lerp(s.End, clampedDelta)
	if distance != nil || normal != nil {
		n, d := closest.NormalTo(p)
		if distance != nil {
			*distance = d
		}
		if normal != nil {
			*normal = n
		}
	}
	return closest
}
func (s Segment2) PolygonSize() int { return 2 }
func (s Segment2) PolygonPoints(target []Point2) int {
	target[0] = s.Start
	target[1] = s.End
	return 2
}

type Plane2 struct {
	A, B, C float32
}

func Plane2PointNormal(point, normal Point2) Plane2 {
	return Plane2{
		A: normal.X,
		B: normal.Y,
		C: -normal.Dot(point),
	}
}

func Plane2Line(start, end Point2) Plane2 {
	n, _ := start.NormalTo(end)
	return Plane2PointNormal(start, n.Tangent())
}

func (s Plane2) Bounds() Bounds2 {
	return InfiniteBounds
}
func (s Plane2) Circle() Circle2 {
	return InfiniteCircle
}
func (s Plane2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	d := s.Distance(p)
	closest := Point2{
		X: p.X - s.A*d,
		Y: p.Y - s.B*d,
	}
	if distance != nil {
		*distance = d
	}
	if normal != nil {
		*normal = Point2{X: s.A, Y: s.B}
	}
	return closest
}
func (s Plane2) Distance(p Point2) float32 {
	return (s.A*p.X + s.B*p.Y + s.C)
}
func (s Plane2) Normal() Point2   { return Point2{X: s.A, Y: s.B} }
func (s Plane2) Origin() Point2   { return s.Normal().Scale(-s.C) }
func (s Plane2) PolygonSize() int { return 2 }
func (s Plane2) PolygonPoints(target []Point2) int {
	normal := s.Normal()
	p0 := normal.Scale(-s.C)
	p1 := p0.Add(normal.Tangent())
	return Line2{Start: p0, End: p1}.PolygonPoints(target)
}

type Ellipse2 struct {
	Center        Point2
	Width, Height float32
}

func (s Ellipse2) Bounds() Bounds2 {
	hw := s.Width * 0.5
	hh := s.Height * 0.5
	return Bounds2{
		Min: s.Center.Move(-hw, -hh),
		Max: s.Center.Move(hw, hh),
	}
}
func (s Ellipse2) Circle() Circle2 {
	return Circle2{
		Center: s.Center,
		Radius: util.Max(s.Width, s.Height) * 0.5,
	}
}
func (s Ellipse2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	n, d := s.Center.NormalTo(p)
	bX := n.X * s.Width
	bY := n.Y * s.Height
	radiusAtAngle := util.Sqrt(bX*bX + bY*bY)
	if distance != nil {
		*distance = d - radiusAtAngle
	}
	if normal != nil {
		*normal = n
	}
	r := util.Min(d, radiusAtAngle)
	return Point2{X: n.X * r, Y: n.Y * r}
}
func (s Ellipse2) PolygonSize() int { return -1 }
func (s Ellipse2) PolygonPoints(target []Point2) int {
	n := len(target)
	angleIncrease := math.Pi * 2 / float32(n)
	angle := float32(0)
	hw := s.Width * 0.5
	hh := s.Height * 0.5
	for i := range target {
		cos, sin := util.CosSin(angle)
		angle += angleIncrease
		target[i] = s.Center.Move(cos*hw, sin*hh)
	}
	return n
}

type Polygon2 struct {
	Points []Point2
}

func (s Polygon2) Bounds() Bounds2 {
	min := s.Points[0]
	max := s.Points[0]
	for i := 1; i < len(s.Points); i++ {
		p := s.Points[i]
		min = min.Min(p)
		max = max.Max(p)
	}
	return Bounds2{Min: min, Max: max}
}
func (s Polygon2) Circle() Circle2 {
	return s.Bounds().Circle()
}
func (s Polygon2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	inside := s.Inside(p)
	if distance == nil && normal == nil {
		if inside {
			return p
		}
	}

	minDistance := float32(math.MaxFloat32)
	minNormal := Point2{}
	minClosest := Point2{}
	p0 := s.Points[len(s.Points)-1]
	for _, p1 := range s.Points {
		line := Segment2{Start: p0, End: p1}
		lineDistance := float32(0)
		lineNormal := Point2{}
		lineClosest := line.Closest(p, &lineDistance, &lineNormal)
		if lineDistance < minDistance {
			minDistance = lineDistance
			minNormal = lineNormal
			minClosest = lineClosest
		}
		p0 = p1
	}
	if inside {
		minDistance = -minDistance
		minNormal.X = -minNormal.X
		minNormal.Y = -minNormal.Y
	}
	if distance != nil {
		*distance = minDistance
	}
	if normal != nil {
		*normal = minNormal
	}
	if inside {
		minClosest = p
	}
	return minClosest
}
func (s Polygon2) Inside(p Point2) bool {
	in := false
	n := len(s.Points)
	i := 0
	j := n - 1
	for i < n {
		jp := s.Points[j]
		ip := s.Points[i]
		if ((ip.Y > p.Y) != (jp.Y > p.Y)) && (p.X < (jp.X-ip.X)*(p.Y-ip.Y)/(jp.Y-ip.Y)+ip.X) {
			in = !in
		}
		j = i
		i++
	}
	return in
}
func (s Polygon2) PolygonSize() int { return len(s.Points) }
func (s Polygon2) PolygonPoints(target []Point2) int {
	n := util.Min(len(target), len(s.Points))
	copy(target, s.Points[:n])
	return n
}

type Inverted2 struct {
	Shape Shape2
}

func (s Inverted2) Bounds() Bounds2 {
	return InfiniteBounds
}
func (s Inverted2) Circle() Circle2 {
	return InfiniteCircle
}
func (s Inverted2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	d := float32(0)
	n := Point2{}
	closest := s.Shape.Closest(p, &d, &n)
	if distance != nil {
		*distance = -d
	}
	if normal != nil {
		*normal = Point2{X: -n.X, Y: -n.Y}
	}
	if d < 0 {
		closest.X -= n.X * d
		closest.Y -= n.Y * d
	}
	return closest
}
func (s Inverted2) PolygonSize() int { return s.Shape.PolygonSize() }
func (s Inverted2) PolygonPoints(target []Point2) int {
	n := s.Shape.PolygonPoints(target)
	util.SliceReverse(target[:n])
	return n
}

type Padded2 struct {
	Shape   Shape2
	Padding float32
	Rounded bool
}

func (s Padded2) Bounds() Bounds2 {
	b := s.Shape.Bounds()
	if b != InfiniteBounds {
		b.Min.Move(-s.Padding, -s.Padding)
		b.Max.Move(s.Padding, s.Padding)
	}
	return b
}
func (s Padded2) Circle() Circle2 {
	c := s.Shape.Circle()
	if c != InfiniteCircle {
		c.Radius += s.Padding
	}
	return c
}
func (s Padded2) Closest(p Point2, distance *float32, normal *Point2) Point2 {
	closestDistance := float32(0)
	closestNormal := Point2{}
	c := s.Shape.Closest(p, &closestDistance, &closestNormal)
	if closestDistance > s.Padding {
		c.X += closestNormal.X * s.Padding
		c.Y += closestNormal.Y * s.Padding
	}
	if distance != nil {
		*distance = closestDistance - s.Padding
	}
	if normal != nil {
		*normal = closestNormal
	}
	return c
}
func (s Padded2) PolygonSize() int {
	if s.Rounded {
		return -s.PolygonSize()
	}
	return s.PolygonSize()
}
func (s Padded2) PolygonPoints(target []Point2) int {
	n := s.Shape.PolygonPoints(target)
	if s.Rounded {
		// TODO capsule
	} else {
		outer := make([]Point2, n)
		last := n - 1
		i0 := last - 1
		i1 := last
		w := s.Padding * 0.5
		for i2 := 0; i2 <= last; i2++ {
			p0 := target[i0]
			p1 := target[i1]
			p2 := target[i2]
			n1, _ := p0.NormalTo(p1)
			n2, _ := p2.NormalTo(p1)
			nx := (n1.Y + -n2.Y)
			ny := (-n1.X + n2.X)
			outer[i1].X = nx*w + p1.X
			outer[i1].Y = ny*w + p1.Y
			i0 = i1
			i1 = i2
		}
		copy(target, outer)
	}
	return n
}
