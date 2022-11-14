package axe

type Path[T Attr[T]] interface {
	Set(out *T, delta float32)
	PointCount() int
	Point(index int) T
}

// Tween

type Tween[T Attr[T]] struct {
	Start T
	End   T
}

var _ Path[Vec2f] = &Tween[Vec2f]{}

func (path Tween[T]) Set(out *T, delta float32) {
	path.Start.Interpolate(path.Start, path.End, delta, out)
}
func (path Tween[T]) PointCount() int { return 2 }
func (path Tween[T]) Point(index int) T {
	switch index {
	case 0:
		return path.Start
	case 1:
		return path.End
	default:
		var empty T
		return empty
	}
}

// Point

type PathPoint[T Attr[T]] struct {
	Value T
}

var _ Path[Vec2f] = &PathPoint[Vec2f]{}

func (path PathPoint[T]) Set(out *T, delta float32) {
	path.Value.Set(out)
}
func (path PathPoint[T]) PointCount() int { return 1 }
func (path PathPoint[T]) Point(index int) T {
	return path.Value
}
