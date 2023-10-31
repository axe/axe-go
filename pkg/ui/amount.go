package ui

type Unit int

const (
	UnitConstant Unit = iota
	UnitPercent
	UnitView   // TODO percent of view dimension
	UnitWindow // TODO percent of window dimension
	UnitScreen // TODO percent of screen dimension
	UnitFont   // TODO percent of global font size
)

func (a Unit) Get(value float32, relative float32) float32 {
	switch a {
	case UnitConstant:
		return value
	case UnitPercent:
		return value / relative
	}
	return value
}

// A unit amount
type Amount struct {
	Value float32
	Unit  Unit
}

func (a Amount) Get(relative float32) float32 {
	return a.Unit.Get(a.Value, relative)
}

func (a *Amount) Set(value float32, unit Unit) {
	a.Value = value
	a.Unit = unit
}

type AmountBounds struct {
	Left   Amount
	Right  Amount
	Top    Amount
	Bottom Amount
}

func (a AmountBounds) GetBounds(w, h float32) Bounds {
	return Bounds{
		Left:   a.Left.Get(w),
		Top:    a.Top.Get(h),
		Right:  a.Right.Get(w),
		Bottom: a.Bottom.Get(h),
	}
}

func (a *AmountBounds) Set(value float32, unit Unit) {
	a.Left.Set(value, unit)
	a.Top.Set(value, unit)
	a.Right.Set(value, unit)
	a.Bottom.Set(value, unit)
}

type AmountCorners struct {
	TopLeft     Amount
	TopRight    Amount
	BottomLeft  Amount
	BottomRight Amount
}

func (a *AmountCorners) Set(value float32, unit Unit) {
	a.TopLeft.Set(value, unit)
	a.TopRight.Set(value, unit)
	a.BottomLeft.Set(value, unit)
	a.BottomRight.Set(value, unit)
}

type AmountPoint struct {
	X Amount
	Y Amount
}

func (a AmountPoint) Get(w, h float32) (float32, float32) {
	return a.X.Get(w), a.Y.Get(h)
}

func (a *AmountPoint) Set(value float32, unit Unit) {
	a.X.Set(value, unit)
	a.Y.Set(value, unit)
}
