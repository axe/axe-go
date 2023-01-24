package axe

type PathFlag int

const (
	PathFlagCyclic PathFlag = 1 << iota
	PathFlagCurved
	PathFlagContinuous
)

func (flag PathFlag) Has(flags PathFlag) bool {
	return flag&flags != 0
}

type Path[T Attr[T]] interface {
	Set(out *T, delta float32)
	PointCount() int
	Point(index int) AnimValue[T]
	Flags() PathFlag
}

// Tween

type Tween[T Attr[T]] struct {
	Start AnimValue[T]
	End   AnimValue[T]
}

var _ Path[Vec2f] = &Tween[Vec2f]{}

func (path Tween[T]) Set(out *T, delta float32) {
	s := path.Start.Get()
	e := path.End.Get()
	s.Interpolate(s, e, delta, out)
}
func (path Tween[T]) PointCount() int { return 2 }
func (path Tween[T]) Point(index int) AnimValue[T] {
	switch index {
	case 0:
		return path.Start
	case 1:
		return path.End
	}
	panic("invalid index to Tween")
}
func (path Tween[T]) Flags() PathFlag { return PathFlagContinuous }

// Point

type PathPoint[T Attr[T]] struct {
	Value AnimValue[T]
}

var _ Path[Vec2f] = &PathPoint[Vec2f]{}

func (path PathPoint[T]) Set(out *T, delta float32) {
	path.Value.Get().Set(out)
}
func (path PathPoint[T]) PointCount() int { return 1 }
func (path PathPoint[T]) Point(index int) AnimValue[T] {
	return path.Value
}
func (path PathPoint[T]) Flags() PathFlag { return PathFlagContinuous | PathFlagCyclic }

// Delta

type PathDelta[T Attr[T]] struct {
	Deltas []float32
	Points []AnimValue[T]
}

var _ Path[Vec2f] = &PathDelta[Vec2f]{}

func (path PathDelta[T]) Set(out *T, delta float32) {
	var ds = path.Deltas
	var end = len(ds) - 2
	var i = 0
	for ds[i+1] < delta && i < end {
		i++
	}
	var d0 = ds[i]
	var d1 = ds[i+1]
	var pd = (delta - d0) / (d1 - d0)
	var p0 = path.Points[i].Get()
	var p1 = path.Points[i+1].Get()

	p0.Interpolate(p0, p1, pd, out)
}
func (path PathDelta[T]) PointCount() int              { return len(path.Points) }
func (path PathDelta[T]) Point(index int) AnimValue[T] { return path.Points[index] }
func (path PathDelta[T]) Flags() PathFlag              { return PathFlagContinuous }

// Linear

type PathLinear[T Attr[T]] struct {
	Inner PathDelta[T]
}

var _ Path[Vec2f] = &PathLinear[Vec2f]{}

func NewPathLinear[T Attr[T]](points []AnimValue[T]) PathLinear[T] {
	return PathLinear[T]{
		Inner: PathDelta[T]{
			Deltas: GetNormalizedDistances(points),
			Points: points,
		},
	}
}
func (path PathLinear[T]) Set(out *T, delta float32)    { path.Inner.Set(out, delta) }
func (path PathLinear[T]) PointCount() int              { return path.Inner.PointCount() }
func (path PathLinear[T]) Point(index int) AnimValue[T] { return path.Inner.Point(index) }
func (path PathLinear[T]) Flags() PathFlag              { return path.Inner.Flags() }

// Functions

func GetNormalizedDistances[T Attr[T]](points []AnimValue[T]) []float32 {
	n := len(points) - 1
	distances := make([]float32, n+1)
	distances[0] = 0

	prev := points[0].Get()

	for i := 1; i <= n; i++ {
		next := points[i].Get()
		distances[i] = distances[i-1] + prev.Distance(next)
		prev = next
	}

	var invlength = 1.0 / distances[n]

	for i := 1; i < n; i++ {
		distances[i] *= invlength
	}

	distances[n] = 1

	return distances
}

type Dynamic[T any] interface {
	Get() T
}

type Computed[T any] interface {
	Compute(current *T)
}

func Resolve[T any](value any) T {
	if dyn, ok := value.(Dynamic[T]); ok {
		return dyn.Get()
	}
	if val, ok := value.(T); ok {
		return val
	}
	panic("unable to resolve value")
}
