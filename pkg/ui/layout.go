package ui

type Layoutable interface {
	Placement() Placement
	Margin() Bounds
	MinSize() Coord
}

type Layout interface {
	Init(init Init)
	Layout(b Bounds, layoutable []Layoutable)
}

type LayoutRow struct {
}
