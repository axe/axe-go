package axe

type Calculator[T Attr[*T]] interface {
	Add(a T, b T, out *T)
}

type Path[T Attr[T]] interface {
	Set(out *T, delta float32)
	PointCount() int
	Point(index int) T
}

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
