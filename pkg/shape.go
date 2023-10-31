package axe

type Shape[A Attr[A]] interface {
	Finite() bool
	Distance(point A) float32
	Normal(point A, out *A) bool
	Raytrace(point A, direction A) bool
	Bounds(bounds *Bounds[A]) bool
	Round(round *Round[A]) bool
	PlaneSign(plane Plane[A]) PlaneSign
	Transform(shapeTransform Trans[A]) Shape[A]
}

type Bounds[A Attr[A]] struct {
	Min       A
	Max       A
	Thickness float32
}

type Bounds2f = Bounds[Vec2f]

var _ Shape[Vec2f] = &Bounds2f{}

func (b Bounds[A]) Finite() bool {
	return true
}

func (b Bounds[A]) Distance(point A) float32 {
	return 0
}
func (b Bounds[A]) DistanceSq(point A) float32 {
	return 0
}
func (b Bounds[A]) Normal(point A, out *A) bool {
	return true
}
func (b Bounds[A]) Raytrace(point A, direction A) bool {
	return true
}
func (b Bounds[A]) Bounds(bounds *Bounds[A]) bool {
	bounds.Min = b.Min
	bounds.Max = b.Max
	return true
}
func (b Bounds[A]) Round(round *Round[A]) bool {
	round.Radius = b.Max.Distance(b.Min)*0.5 + float32(b.Thickness)
	b.Min.Lerp(b.Max, 0.5, &round.Center)
	return true
}
func (b Bounds[A]) CornerCount() int {
	var empty A
	c := empty.Components()
	return 1 << c
}
func (b Bounds[A]) Corner(cornerIndex int) A {
	var corner A
	componentCount := corner.Components()
	for i := 0; i < componentCount; i++ {
		if cornerIndex&(1<<i) == 0 {
			corner.SetComponent(i, b.Min.GetComponent(i), &corner)
		} else {
			corner.SetComponent(i, b.Max.GetComponent(i), &corner)
		}
	}
	return corner
}
func (b Bounds[A]) Corners() []A {
	n := b.CornerCount()
	corners := make([]A, n)
	for i := 0; i < n; i++ {
		corners[i] = b.Corner(i)
	}
	return corners
}
func (b Bounds[A]) PlaneSign(plane Plane[A]) PlaneSign {
	cornerCount := b.CornerCount()
	s0 := plane.Sign(b.Corner(0))
	for i := 1; i < cornerCount; i++ {
		if s1 := plane.Sign(b.Corner(i)); s1 != s0 {
			return PlaneSignIntersects
		}
	}
	return s0
}
func (b Bounds[A]) Transform(shapeTransform Trans[A]) Shape[A] {
	return &Bounds[A]{
		Min: shapeTransform.Transform(b.Min),
		Max: shapeTransform.Transform(b.Max),
	}
}

type Line[A Attr[A]] struct {
	Start A
	End   A
}

func (l Line[A]) Diff() A {
	var diff A
	l.End.Sub(l.Start, &diff)
	return diff
}
func (l Line[A]) PlaneSign(plane Plane[A]) PlaneSign {
	s := plane.Sign(l.Start)
	e := plane.Sign(l.End)
	if s == 0 {
		return e
	}
	if e == 0 || s == e {
		return s
	}
	return PlaneSignIntersects
}

type Segment[A Attr[A]] struct {
	Start A
	End   A
}

type Round[A Attr[A]] struct {
	Center A
	Radius float32
}

func (round Round[A]) PlaneSign(plane Plane[A]) PlaneSign {
	d := plane.Distance(round.Center)
	return PlaneSign(Sign(d - float32(Sign(d))*round.Radius))
}

type Circle = Round[Vec2f]
type Sphere = Round[Vec3f]
