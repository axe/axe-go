package axe

type AnimValue[T any] interface {
	Get() T
}

type AnimConstant[T any] struct {
	Value T
}

var _ AnimValue[int] = AnimConstant[int]{}

func (con AnimConstant[T]) Get() T {
	return con.Value
}
