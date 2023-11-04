package ui

type Bounds struct {
	Left, Top, Right, Bottom float32
}

func (b Bounds) IsZero() bool {
	return b.Left >= b.Right && b.Top >= b.Bottom
}
func (b Bounds) Width() float32 {
	return b.Right - b.Left
}
func (b Bounds) Height() float32 {
	return b.Bottom - b.Top
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
func (b Bounds) Inside(c Coord) bool {
	return !(c.X < b.Left || c.X > b.Right || c.Y < b.Top || c.Y > b.Bottom)
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

type Theme struct {
	TextStyles    TextStyles
	StateModifier map[State]VertexModifier

	// Components map[string]*ComponentTheme
	Fonts      map[string]*Font
	Animations Animations
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
