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
func (b Bounds) Contains(c Coord) bool {
	return !(c.X < b.Left || c.X > b.Right || c.Y < b.Top || c.Y > b.Bottom)
}

type Theme struct {
	DefaultFontSize  float32
	DefaultFontColor Color
	DefaultFont      string
	StateModifier    map[State]VertexModifier

	// Components map[string]*ComponentTheme
	Fonts map[string]*Font
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
