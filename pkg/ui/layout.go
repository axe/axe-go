package ui

type Layoutable interface {
	Placement() Placement
	Margin() Bounds
	MinSize() Coord
}

type Layout interface {
	Init(b *Base, init Init)
	PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord
	Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base)
}

// A layout which places all children in a column (vertical stack).
// The children can have their width expanded to the full width or default
// to their preferred size. If they are not expanded to full width they can
// be horizontally aligned. Optionally the size inferred by the placement
// can be factored into the preferred size as well (it's not by default).
type LayoutColumn struct {
	FullWidth           bool
	HorizontalAlignment Alignment
	Spacing             Amount
	UsePlacementSizes   bool
}

func (l LayoutColumn) Init(b *Base, init Init) {}
func (l LayoutColumn) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}

	for _, child := range layoutable {
		childMargin := child.Margin.GetBounds(ctx.AmountContext)
		childMaxWidth := maxWidth
		if childMaxWidth > 0 {
			childMaxWidth -= childMargin.Left + childMargin.Right
		}
		childSize := child.PreferredSize(ctx, !l.UsePlacementSizes, childMaxWidth)
		childSize.X += childMargin.Left + childMargin.Right
		childSize.Y += childMargin.Top + childMargin.Bottom
		size.X = max(size.X, childSize.X)
		size.Y += childSize.Y
	}

	spacing := l.Spacing.Get(ctx.AmountContext, false)
	size.Y += spacing * float32(len(layoutable)-1)

	return size
}
func (l LayoutColumn) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	if len(layoutable) == 0 {
		return
	}

	offsetY := float32(0)
	deltaX := float32(l.HorizontalAlignment)
	width := bounds.Width()
	spacing := l.Spacing.Get(ctx.AmountContext, false)

	for _, child := range layoutable {
		margin := child.Margin.GetBounds(ctx.AmountContext)
		size := child.PreferredSize(ctx, !l.UsePlacementSizes, width-margin.Left-margin.Right)
		if l.FullWidth {
			size.X = width - margin.Left - margin.Right
		}
		placement := child.Placement.
			Attach(deltaX, 0, size.X, size.Y).
			Shift(margin.Left, offsetY+margin.Top)
		child.SetPlacement(placement)
		offsetY += size.Y + margin.Top + margin.Bottom + spacing
	}
}

// A layout which places all children in a row (horizontal stack).
// The children can have their height expanded to the full height or default
// to their preferred size. If they are not expanded to full height they can
// be vertically aligned. Optionally the size inferred by the placement
// can be factored into the preferred size as well (it's not by default).
type LayoutRow struct {
	FullHeight        bool
	VerticalAlignment Alignment
	Spacing           Amount
	UsePlacementSizes bool
}

func (l LayoutRow) Init(b *Base, init Init) {}
func (l LayoutRow) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}

	for _, child := range layoutable {
		childMargin := child.Margin.GetBounds(ctx.AmountContext)
		childMaxWidth := maxWidth
		if childMaxWidth > 0 {
			childMaxWidth -= childMargin.Left + childMargin.Right
		}
		childSize := child.PreferredSize(ctx, !l.UsePlacementSizes, childMaxWidth)
		childSize.X += childMargin.Left + childMargin.Right
		childSize.Y += childMargin.Top + childMargin.Bottom
		size.Y = max(size.Y, childSize.Y)
		size.X += childSize.X
	}

	spacing := l.Spacing.Get(ctx.AmountContext, false)
	size.X += spacing * float32(len(layoutable)-1)

	return size
}
func (l LayoutRow) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	if len(layoutable) == 0 {
		return
	}

	offsetX := float32(0)
	deltaY := float32(l.VerticalAlignment)
	height := bounds.Height()
	spacing := l.Spacing.Get(ctx.AmountContext, true)

	for _, child := range layoutable {
		margin := child.Margin.GetBounds(ctx.AmountContext)
		size := child.PreferredSize(ctx, !l.UsePlacementSizes, 0)
		if l.FullHeight {
			size.Y = height - margin.Top - margin.Bottom
		}
		placement := child.Placement.
			Attach(0, deltaY, size.X, size.Y).
			Shift(offsetX+margin.Left, margin.Top)
		child.SetPlacement(placement)
		offsetX += size.X + margin.Left + margin.Right + spacing
	}
}

// A layout which places all children in a grid. The number of columns in the
// grid can be given or a MinSize can be given with a non-zero X component. That
// component is used with the spacing and the width of the parent to maintain
// a dynamic number of columns based on the min width. The width of the cell
// is evenly distributed. The height of the cell can be determined by the max
// preferred size of a child in a row while also using the MinSize.Y defined
// on the layout which all can be overridden if a non-zero AspectRatio is defined
// which will make all rows have the same height respective to the cell width.
// Alignment of a component within a cell is determine by Vertical and Horizontal
// alignment, unless FullWidth & FullHeight are given which means the component
// expands to fit the cell.
type LayoutGrid struct {
	FullHeight          bool
	FullWidth           bool
	VerticalAlignment   Alignment
	HorizontalAlignment Alignment
	VerticalSpacing     Amount
	HorizontalSpacing   Amount
	UsePlacementSizes   bool
	MinSize             Coord
	Columns             int
	AspectRatio         float32
	EqualHeights        bool
}

const (
	LayoutGridWidthRoundingError = 0.001
)

func (l LayoutGrid) Init(b *Base, init Init) {}
func (l LayoutGrid) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}
	// TODO
	return size
}
func (l LayoutGrid) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	if len(layoutable) == 0 {
		return
	}

	width := bounds.Width()
	spacingX := l.HorizontalSpacing.Get(ctx.AmountContext, true)
	spacingY := l.VerticalSpacing.Get(ctx.AmountContext, false)

	columns := l.Columns
	if columns == 0 && l.MinSize.X != 0 {
		columns = int((width + spacingX) / (l.MinSize.X + spacingX))
	}
	if columns < 1 {
		columns = 1
	}
	cellWidth := (width - (float32(columns-1) * spacingX)) / float32(columns)
	if cellWidth < l.MinSize.X {
		columns = int((width + spacingX) / (l.MinSize.X + spacingX))
		if columns < 1 {
			columns = 1
		}
		cellWidth = (width - (float32(columns-1) * spacingX)) / float32(columns)
	}
	cellHeights := make([]float32, (len(layoutable)+columns-1)/int(columns))
	cellWidth -= LayoutGridWidthRoundingError

	sizes, fullSizes, margins := getLayoutableDimensions(ctx, l.UsePlacementSizes, cellWidth, layoutable)

	if l.AspectRatio > 0 {
		for i := range cellHeights {
			cellHeights[i] = cellWidth * l.AspectRatio
		}
	} else {
		for childIndex := range layoutable {
			row := childIndex / columns
			cellHeights[row] = max(fullSizes[childIndex].Y, cellHeights[row])
		}
	}
	maxHeight := float32(0)
	for i := range cellHeights {
		cellHeights[i] = max(cellHeights[i], l.MinSize.Y)
		maxHeight = max(maxHeight, cellHeights[i])
	}
	if l.EqualHeights {
		for i := range cellHeights {
			cellHeights[i] = maxHeight
		}
	}

	offsetY := float32(0)

	for childIndex, child := range layoutable {
		size := sizes[childIndex]
		fullSize := fullSizes[childIndex]
		margin := margins[childIndex]
		row := childIndex / columns
		col := childIndex % columns
		cellHeight := cellHeights[row]

		if l.FullWidth || fullSize.X > cellWidth {
			size.X = cellWidth - margin.Left - margin.Right
			fullSize.X = cellWidth
		}
		if l.FullHeight || fullSize.Y > cellHeight {
			size.Y = cellHeight - margin.Top - margin.Bottom
			fullSize.Y = cellHeight
		}

		offsetX := float32(col) * (cellWidth + spacingX)
		alignX := l.HorizontalAlignment.Compute(cellWidth - fullSize.X)
		alignY := l.VerticalAlignment.Compute(cellHeight - fullSize.Y)

		left := offsetX + alignX + margin.Left
		top := offsetY + alignY + margin.Top

		child.SetPlacement(Absolute(left, top, size.X, size.Y))

		if col == columns-1 {
			offsetY += cellHeight + spacingY
		}
	}
}

// A layout which places all children on after another and wraps them
// if it can't fit in a line.
type LayoutInline struct {
	VerticalAlignment   Alignment
	HorizontalAlignment Alignment
	VerticalSpacing     Amount
	HorizontalSpacing   Amount
	UsePlacementSizes   bool
}

func (l LayoutInline) Init(b *Base, init Init) {}
func (l LayoutInline) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}
	// TODO
	return size
}
func (l LayoutInline) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	if len(layoutable) == 0 {
		return
	}

	type line struct {
		start, end    int
		width, height float32
	}

	width := bounds.Width()
	spacingX := l.HorizontalSpacing.Get(ctx.AmountContext, true)
	spacingY := l.VerticalSpacing.Get(ctx.AmountContext, false)
	sizes, fullSizes, margins := getLayoutableDimensions(ctx, l.UsePlacementSizes, width, layoutable)
	lines := make([]line, 0, len(layoutable))
	currentLine := line{}

	for childIndex := range layoutable {
		size := fullSizes[childIndex]
		nextWidth := currentLine.width + size.X
		if childIndex > currentLine.start {
			nextWidth += spacingX
		}
		if nextWidth > width && childIndex > currentLine.start {
			currentLine.end = childIndex
			lines = append(lines, currentLine)
			currentLine = line{start: childIndex}
		}
		currentLine.width += size.X
		if childIndex > currentLine.start {
			currentLine.width += spacingX
		}
		currentLine.height = max(currentLine.height, size.Y)
	}
	if currentLine.width > 0 {
		currentLine.end = len(layoutable)
		lines = append(lines, currentLine)
	}

	offsetY := float32(0)

	for _, line := range lines {
		offsetX := max(0, l.HorizontalAlignment.Compute(width-line.width))
		for i := line.start; i < line.end; i++ {
			child := layoutable[i]
			size := sizes[i]
			fullSize := fullSizes[i]
			margin := margins[i]
			alignY := l.VerticalAlignment.Compute(line.height - fullSize.Y)

			child.SetPlacement(Absolute(offsetX+margin.Left, offsetY+alignY+margin.Top, size.X, size.Y))
			offsetX += fullSize.X + spacingX
		}
		offsetY += line.height + spacingY
	}
}

func getLayoutableDimensions(ctx *RenderContext, usePlacementSizes bool, width float32, children []*Base) (sizes []Coord, fullSizes []Coord, margins []Bounds) {
	sizes = make([]Coord, len(children))
	fullSizes = make([]Coord, len(children))
	margins = make([]Bounds, len(children))

	for childIndex, child := range children {
		margin := child.Margin.GetBounds(ctx.AmountContext)
		size := child.PreferredSize(ctx, !usePlacementSizes, width-margin.Left-margin.Right)

		sizes[childIndex] = size
		fullSizes[childIndex] = Coord{
			X: size.X + margin.Left + margin.Right,
			Y: size.Y + margin.Top + margin.Bottom,
		}
		margins[childIndex] = margin
	}

	return
}
