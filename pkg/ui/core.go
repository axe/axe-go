package ui

import "github.com/axe/axe-go/pkg/id"

type Theme struct {
	TextStyles       TextStyles
	StatePostProcess map[State]PostProcess

	// Components map[string]*ComponentTheme
	Fonts         id.DenseMap[*Font, uint16, uint8]
	Cursors       id.DenseMap[ExtentTile, uint16, uint8]
	DefaultCursor id.Identifier
	Animations    Animations
}

type State = Flags

const (
	StateDefault State = 1 << iota
	StateHover
	StatePressed
	StateFocused
	StateDisabled
	StateDragging
	StateDragOver
	StateSelected // checked, chosen option
)

type StateFn = func(s State) bool

type Dirty = Flags

const (
	DirtyNone Dirty = (1 << iota) >> 1
	DirtyPlacement
	DirtyDeepPlacement
	DirtyVisual
)

type Watch[V comparable] struct {
	value   V
	changed bool
}

func NewWatch[V comparable](value V) Watch[V] { return Watch[V]{value: value} }

func (w Watch[V]) Get() V      { return w.value }
func (w Watch[V]) Dirty() bool { return w.changed }
func (w *Watch[V]) Clean()     { w.changed = false }
func (w *Watch[V]) Cleaned() bool {
	cleaned := w.changed
	w.changed = false
	return cleaned
}
func (w *Watch[V]) Set(value V) {
	if w.value != value {
		w.changed = true
		w.value = value
	}
}

func Override[V any](value V) *V {
	return &value
}

type Optional[V any] struct {
	value V
	set   bool
}

func NewOptional[V any](value V) Optional[V] { return Optional[V]{value: value, set: true} }
