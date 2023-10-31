package ui

type Layoutable interface {
	Placement() Placement
	Margin() Bounds
}

type Layout interface {
	Init(init Init)
	Layout(b Bounds, layoutable []Layoutable)
}

type LayoutRow struct {
}
