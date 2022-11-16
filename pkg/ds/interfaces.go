package ds

type Sized interface {
	Len() int
	Cap() int
	IsEmpty() bool
}

type Linear[V any] interface {
	Peek() V
	Pop() V
	Push(value V) bool
}

type Indexed[V any] interface {
	At(index int) V
	Len() int
}

type Clearable interface {
	Clear()
}
