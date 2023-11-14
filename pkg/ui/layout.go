package ui

import "github.com/axe/axe-go/pkg/util"

type Layoutable interface {
	Placement() Placement
	Margin() Bounds
	MinSize() Coord
}

type Layout interface {
	Init(b *Base)
	PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord
	Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base)
}

var _ Layout = LayoutColumn{}
var _ Layout = LayoutRow{}
var _ Layout = LayoutGrid{}
var _ Layout = LayoutInline{}
var _ Layout = LayoutStatic{}

// A way to control the weight of a metric for a component at a given index.
// All weights for N components will add up to 1
// If the number of weights is <= 1 then all weights returned are equally distributed.
type LayoutWeights []float32

func (w LayoutWeights) GetWeight(i, n int) float32 {
	if len(w) <= 1 {
		return float32(1) / float32(n)
	} else {
		total := w.TotalWeight(n)
		if total != 0 {
			return w.WeightAt(i) / total
		} else {
			return float32(1) / float32(n)
		}
	}
}
func (w LayoutWeights) WeightAt(i int) float32 {
	last := len(w) - 1
	if i > last {
		return w[last]
	} else {
		return w[i]
	}
}
func (w LayoutWeights) TotalWeight(n int) float32 {
	last := len(w) - 1
	totalWeight := float32(0)
	for i := 0; i < n; i++ {
		if i <= last {
			totalWeight += w[i]
		} else {
			totalWeight += w[last]
		}
	}
	return totalWeight
}

// A way to control the dimensions of a metric in a layout.
type LayoutDimensions []float32

func (d LayoutDimensions) Get(i int) float32 {
	last := len(d) - 1
	if last == -1 {
		return 0
	} else if i <= last {
		return d[i]
	} else {
		return d[last]
	}
}

// A layout which places all children in a column (vertical stack).
// The children can have their width expanded to the full width or default
// to their preferred size. If they are not expanded to full width they can
// be horizontally aligned. Optionally the size inferred by the placement
// can be factored into the preferred size as well (it's not by default).
type LayoutColumn struct {
	FullWidth           bool
	FullHeight          bool
	FullHeightWeights   LayoutWeights
	EqualWidths         bool
	HorizontalAlignment Alignment
	Spacing             Amount
}

func (l LayoutColumn) Init(b *Base) {}
func (l LayoutColumn) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}

	spacing := l.Spacing.Get(ctx.AmountContext, false)
	spacingTotal := spacing * float32(len(layoutable)-1)
	sizings := getLayoutSizings(ctx, maxWidth, layoutable)

	size.X = sizings.MaxWidth
	size.Y = sizings.TotalHeight + spacingTotal

	return size
}
func (l LayoutColumn) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	n := len(layoutable)
	if n == 0 {
		return
	}

	offsetY := float32(0)
	deltaX := float32(l.HorizontalAlignment)
	width, maxHeight := bounds.Dimensions()
	sizings := getLayoutSizings(ctx, width, layoutable)
	spacing := l.Spacing.Get(ctx.AmountContext, false)
	spacingTotal := spacing * float32(n-1)
	totalHeight := sizings.TotalHeight + spacingTotal

	if l.FullHeight && maxHeight > totalHeight {
		remaining := maxHeight - totalHeight
		for childIndex := range sizings.Sizings {
			sizing := &sizings.Sizings[childIndex]
			add := remaining * l.FullHeightWeights.GetWeight(childIndex, n)
			sizing.FullHeight += add
			sizing.Height += add
		}
	}

	for childIndex, child := range layoutable {
		sizing := sizings.Sizings[childIndex]
		paddingWidth := sizing.FullWidth - sizing.Width
		if l.FullWidth {
			sizing.FullWidth = width
			sizing.Width = width - paddingWidth
		} else if l.EqualWidths {
			sizing.FullWidth = sizings.MaxWidth
			sizing.Width = sizings.MaxWidth - paddingWidth
		}
		placement := child.Placement.
			Attach(deltaX, 0, sizing.Width, sizing.Height).
			Shift(sizing.OffsetX, offsetY+sizing.OffsetY)
		child.SetPlacement(placement)
		offsetY += sizing.FullHeight + spacing
	}
}

// A layout which places all children in a row (horizontal stack).
// The children can have their height expanded to the full height or default
// to their preferred size. If they are not expanded to full height they can
// be vertically aligned. Optionally the size inferred by the placement
// can be factored into the preferred size as well (it's not by default).
type LayoutRow struct {
	FullHeight        bool
	FullWidth         bool
	FullWidthWeights  LayoutWeights
	EqualHeights      bool
	VerticalAlignment Alignment
	Spacing           Amount
	ExpandRight       bool
}

func (l LayoutRow) Init(b *Base) {}
func (l LayoutRow) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}

	n := len(layoutable)
	spacing := l.Spacing.Get(ctx.AmountContext, false)
	spacingTotal := spacing * float32(n-1)
	minSizings := getLayoutSizings(ctx, 0, layoutable)
	minSizings.TotalWidth += spacingTotal

	// If the minimum sizings already goes beyond the maxWidth...
	if minSizings.TotalWidth >= maxWidth {
		// We can't fit them in the maxWidth, so report the minimum size (by width) possible.
		size.X = minSizings.TotalWidth
		size.Y = minSizings.MaxHeight
	} else {
		// How much width can we give each layoutable
		availableMaxWidth := maxWidth - spacingTotal

		// Compute max sizings
		maxSizings := getLayoutSizings(ctx, availableMaxWidth, layoutable)
		maxSizings.TotalWidth += spacingTotal

		// If their maximum width fits in the desired maxWidth...
		if maxSizings.TotalWidth <= maxWidth {
			// We can fit them in maxWidth at their max sizes...
			size.X = maxSizings.TotalWidth
			size.Y = maxSizings.MaxHeight
		} else {
			// We need to look at the max sizings. If any of them take up less width then what we can evenly
			// divide up then keep track how much extra width we can give to the potentially greedier components.
			// We need to simulate layout to compute a real height
			targetWidth := availableMaxWidth / float32(len(layoutable))
			extraWidth := float32(0)
			maxHeight := float32(0)
			totalWidth := float32(0)

			for _, maxSizing := range maxSizings.Sizings {
				// If the max size of this child can fit in its available width, consider it layed out at its max size
				if maxSizing.FullWidth <= targetWidth {
					extraWidth += targetWidth - maxSizing.FullWidth
					maxHeight = util.Max(maxHeight, maxSizing.FullHeight)
					totalWidth += maxSizing.FullWidth
				}
			}

			start, end, move := iteratorRange(n, l.ExpandRight)
			for i := start; i != end; i += move {
				maxSizing := maxSizings.Sizings[i]
				if maxSizing.FullWidth > targetWidth {
					if extraWidth < 0 {
						newSizing := getLayoutSizing(ctx, targetWidth, layoutable[i])
						maxHeight = util.Max(maxHeight, newSizing.FullHeight)
						totalWidth += newSizing.FullWidth
					} else {
						availableWidth := targetWidth + extraWidth
						newSizing := getLayoutSizing(ctx, availableWidth, layoutable[i])
						maxHeight = util.Max(maxHeight, newSizing.FullHeight)
						extraWidth -= newSizing.FullWidth - targetWidth
						totalWidth += newSizing.FullWidth
					}
				}
			}

			size.X = totalWidth + spacingTotal
			size.Y = maxHeight
		}
	}

	return size
}
func (l LayoutRow) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	n := len(layoutable)
	if n == 0 {
		return
	}

	maxWidth, boundsHeight := bounds.Dimensions()
	spacing := l.Spacing.Get(ctx.AmountContext, false)
	spacingTotal := spacing * float32(n-1)
	minSizings := getLayoutSizings(ctx, 0, layoutable)
	minSizings.TotalWidth += spacingTotal

	sizings := minSizings.Sizings
	maxHeight := minSizings.MaxHeight

	// If the minimum sizings do not go beyond the maxWidth...
	if minSizings.TotalWidth < maxWidth {
		// How much width can we give each layoutable
		availableMaxWidth := maxWidth - spacingTotal

		// Compute max sizings
		maxSizings := getLayoutSizings(ctx, availableMaxWidth, layoutable)
		maxSizings.TotalWidth += spacingTotal

		// If their maximum width fits in the desired maxWidth...
		if maxSizings.TotalWidth <= maxWidth {
			// We can fit them in maxWidth at their max sizes...
			sizings = maxSizings.Sizings
			maxHeight = maxSizings.MaxHeight
		} else {

			// We need to look at the max sizings. If any of them take up less width then what we can evenly
			// divide up then keep track how much extra width we can give to the potentially greedier components.
			// We need to simulate layout to compute a real height
			targetWidth := availableMaxWidth / float32(len(layoutable))
			extraWidth := float32(0)
			maxHeight = 0

			for childIndex, maxSizing := range maxSizings.Sizings {
				if maxSizing.FullWidth <= targetWidth {
					extraWidth += targetWidth - maxSizing.FullWidth
					maxHeight = util.Max(maxHeight, maxSizing.FullHeight)
					sizings[childIndex] = maxSizing
				}
			}

			start, end, move := iteratorRange(n, l.ExpandRight)
			for i := start; i != end; i += move {
				maxSizing := maxSizings.Sizings[i]
				if maxSizing.FullWidth > targetWidth {
					if extraWidth <= 0 {
						newSizing := getLayoutSizing(ctx, targetWidth, layoutable[i])
						maxHeight = util.Max(maxHeight, newSizing.FullHeight)
						sizings[i] = newSizing
					} else {
						availableWidth := targetWidth + extraWidth
						newSizing := getLayoutSizing(ctx, availableWidth, layoutable[i])
						maxHeight = util.Max(maxHeight, newSizing.FullHeight)
						extraWidth -= newSizing.FullWidth - targetWidth
						sizings[i] = newSizing
					}
				}
			}
		}
	}

	if l.FullWidth {
		totalWidth := spacingTotal
		for _, sizing := range sizings {
			totalWidth += sizing.FullWidth
		}
		if totalWidth < maxWidth {
			remaining := maxWidth - totalWidth
			for i := range sizings {
				sizing := &sizings[i]
				add := remaining * l.FullWidthWeights.GetWeight(i, n)
				sizing.Width += add
				sizing.FullWidth += add
			}
		}
	}

	deltaY := float32(l.VerticalAlignment)
	offsetX := float32(0)

	for childIndex, child := range layoutable {
		sizing := sizings[childIndex]
		heightPadding := sizing.HeightPadding()
		if l.FullHeight {
			sizing.FullHeight = boundsHeight
			sizing.Height = boundsHeight - heightPadding
		} else if l.EqualHeights {
			sizing.FullHeight = maxHeight
			sizing.Height = maxHeight - heightPadding
		}
		placement := child.Placement.
			Attach(0, deltaY, sizing.Width, sizing.Height).
			Shift(offsetX+sizing.OffsetX, sizing.OffsetY)
		child.SetPlacement(placement)
		offsetX += sizing.FullWidth + spacing
	}
}

// A layout which places all children in a grid. The number of columns in the
// grid can be given or MinWidths can be given with a non-zero X component. That
// component is used with the spacing and the width of the parent to maintain
// a dynamic number of columns based on the min width. The width of the cell
// is evenly distributed is EqualWidths. The height of the cell can be determined
// by the max preferred size of a child in a row while also using the MinHeights defined
// on the layout which all can be overridden if a non-zero AspectRatio is defined
// which will make all rows have the same height respective to the cell width.
// Alignment of a component within a cell is determine by Vertical and Horizontal
// alignment, unless FullWidth & FullHeight are given which means the component
// expands to fit the cell.
type LayoutGrid struct {
	// The grid will take up the full width of the bounds. May result in ignoring other width settings.
	GridFullWidth bool
	// When stretching the columns to match the width of the bounds - how should each column
	// be stretched? A value of zero will cause no stretching, and non-zero values will be
	// divided up based on the value / the sum of values. One or zero weights will cause
	// the widths to be equally stretched.
	GridFullWidthWeights LayoutWeights
	// The grid will take up the full height of the bounds. May result in ignoring other height settings.
	GridFullHeight bool
	// When stretching the rows to match the height of the bounds - how should each row
	// be stretched? A value of zero will cause no stretching, and non-zero values will be
	// divided up based on the value / the sum of values. One or zero weights will cause
	// the heights to be equally stretched.
	GridFullHeightWeights LayoutWeights
	// Each cell component is expanded to the row height
	FullHeight bool
	// Each cell component is expanded to the column width
	FullWidth bool
	// When FullHeight is false, how to align the cell component veritcally in the cell
	VerticalAlignment Alignment
	// When FullWidth is false, how to align the cell component horizontally in the cell
	HorizontalAlignment Alignment
	// How much space between rows
	VerticalSpacing Amount
	// How much space between columns
	HorizontalSpacing Amount
	// Defines the min widths for one or all columns. If the number of columns extends beyond the number
	// of min widths defined, they will use the last min width defined.
	MinWidths LayoutDimensions
	// Defines the max widths for one or all columns. If the number of columns extends beyond the number
	// of max widths defined, they will use the last max width defined.
	MaxWidths LayoutDimensions
	// Defines the min heights for one or all rows. If the number of rows extends beyond the number
	// of min heights defined, they will use the last min height defined.
	MinHeights LayoutDimensions
	// Defines the max heights for one or all rows. If the number of rows extends beyond the number
	// of max heights defined, they will use the last max height defined.
	MaxHeights LayoutDimensions
	// The number of columns in the grid. A zero value will have the grid calculate the number of columns.
	Columns int
	// A value used in determining how many columns should be in the grid when not defined.
	// A value of 0.0 will try to fit as many columns as possible while a value of 1.0 will try to fit as few columns as possible.
	ColumnsDynamicDelta float32
	// If the cell heights should be a ratio of their width. A value of 0.0 has no affect.
	// A value of 0.5 will make the height half the width.
	AspectRatio float32
	// If all the rows in the grid should try to have equal heights. Min/Max heights may interfere with this.
	EqualHeights bool
	// If all the columns in the grid should try to have equal widths. Min/Max widths may interfere with this.
	EqualWidths bool
}

const (
	LayoutGridWidthRoundingError = 0.001
)

type layoutGridInfo struct {
	sizings            LayoutSizings
	spacingX, spacingY float32
	layoutGridSizings
}

type layoutGridSizings struct {
	heights      []float32
	widths       []float32
	totalSpacing Coord
	totalSize    Coord
	rows         int
	columns      int
}

func (info layoutGridInfo) getSizingsFor(columns int, grid *LayoutGrid) layoutGridSizings {
	last := len(info.sizings.Sizings) - 1
	rows := (last + columns) / columns
	sizings := layoutGridSizings{
		columns: columns,
		rows:    rows,
		widths:  make([]float32, columns),
		heights: make([]float32, rows),
		totalSpacing: Coord{
			X: float32(columns-1) * info.spacingX,
			Y: float32(rows-1) * info.spacingY,
		},
	}
	if len(grid.MinWidths) > 0 {
		for i := range sizings.widths {
			sizings.widths[i] = grid.MinWidths.Get(i)
		}
	}
	if len(grid.MinHeights) > 0 {
		for i := range sizings.heights {
			sizings.heights[i] = grid.MinHeights.Get(i)
		}
	}
	for childIndex, sizing := range info.sizings.Sizings {
		col := childIndex % columns
		row := childIndex / columns
		sizings.widths[col] = util.Max(sizings.widths[col], sizing.FullWidth)
		sizings.heights[row] = util.Max(sizings.heights[row], sizing.FullHeight)
	}
	maxHeight := float32(0)
	maxWidth := float32(0)
	for _, width := range sizings.widths {
		maxWidth = util.Max(maxWidth, width)
	}
	for _, height := range sizings.heights {
		maxHeight = util.Max(maxHeight, height)
	}
	if grid.EqualWidths {
		for i := range sizings.widths {
			sizings.widths[i] = maxWidth
		}
	}
	if grid.AspectRatio != 0 {
		for i := range sizings.heights {
			sizings.heights[i] = maxWidth * grid.AspectRatio
		}
	} else if grid.EqualHeights {
		for i := range sizings.heights {
			sizings.heights[i] = maxHeight
		}
	}
	if len(grid.MaxWidths) > 0 {
		for i := range sizings.widths {
			maxWidth := grid.MaxWidths.Get(i)
			if maxWidth > 0 {
				sizings.widths[i] = util.Min(sizings.widths[i], maxWidth)
			}
		}
	}
	if len(grid.MaxHeights) > 0 {
		for i := range sizings.heights {
			maxHeight := grid.MaxHeights.Get(i)
			if maxHeight > 0 {
				sizings.heights[i] = util.Min(sizings.heights[i], maxHeight)
			}
		}
	}
	sizings.totalSize = sizings.totalSpacing
	for _, columnWidth := range sizings.widths {
		sizings.totalSize.X += columnWidth
	}
	for _, rowHeight := range sizings.heights {
		sizings.totalSize.Y += rowHeight
	}

	return sizings
}

func (l LayoutGrid) Init(b *Base) {}
func (l *LayoutGrid) getSizingInfo(ctx *RenderContext, maxWidth float32, layoutable []*Base) layoutGridInfo {
	info := layoutGridInfo{
		spacingX: l.HorizontalSpacing.Get(ctx.AmountContext, true),
		spacingY: l.VerticalSpacing.Get(ctx.AmountContext, false),
	}

	columns := l.Columns
	if columns != 0 {
		if l.EqualWidths {
			// equal width with defined number of columns is easiest, we calculate column width add calculate sizings based on that.
			columnWidth := ((maxWidth + info.spacingX) / float32(columns)) - info.spacingX
			info.sizings = getLayoutSizings(ctx, columnWidth, layoutable)
		} else {
			// we look at the min size, find out how much "extra" space we have to work with, then divide it up evenly.
			info.sizings = getLayoutSizings(ctx, 0, layoutable)
			minWidths := make([]float32, columns)
			for childIndex, sizing := range info.sizings.Sizings {
				col := childIndex % columns
				minWidths[col] = util.Max(minWidths[col], sizing.FullWidth)
			}
			totalMinWidths := float32(0)
			for _, minWidth := range minWidths {
				totalMinWidths += minWidth
			}
			totalMinWidths += info.spacingX * float32(columns-1)
			extra := maxWidth - totalMinWidths
			if extra > 0 {
				extraSpacePerColumn := extra / float32(columns)
				for childIndex, sizing := range info.sizings.Sizings {
					col := childIndex % columns
					columnWidth := minWidths[col]
					if sizing.FullWidth < columnWidth {
						info.sizings.Sizings[childIndex] = getLayoutSizing(ctx, columnWidth+extraSpacePerColumn, layoutable[childIndex])
					}
				}
			}
		}
	} else {
		// we calculate columns based on max column width, and with the real
		// column width we calculate we determine sizings.
		layoutWidth := maxWidth
		if l.Columns == 0 {
			layoutWidth = l.ColumnsDynamicDelta * maxWidth
		}
		info.sizings = getLayoutSizings(ctx, layoutWidth, layoutable)
		maxChildWidth := float32(0)
		for _, width := range l.MinWidths {
			maxChildWidth = util.Max(maxChildWidth, width)
		}
		for _, sizing := range info.sizings.Sizings {
			maxChildWidth = util.Max(maxChildWidth, sizing.FullWidth)
		}
		columns = int((maxWidth + info.spacingX) / (maxChildWidth + info.spacingX))
		if columns == 0 {
			columns = 1
		}
		columnWidth := ((maxWidth + info.spacingX) / float32(columns)) - info.spacingX
		info.sizings = getLayoutSizings(ctx, columnWidth, layoutable)
	}
	info.layoutGridSizings = info.getSizingsFor(columns, l)

	return info
}
func (l LayoutGrid) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}

	if len(layoutable) == 0 {
		return size
	}

	info := l.getSizingInfo(ctx, maxWidth, layoutable)
	size = info.totalSize

	return size
}
func (l LayoutGrid) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	n := len(layoutable)
	if n == 0 {
		return
	}

	maxWidth, maxHeight := bounds.Dimensions()
	info := l.getSizingInfo(ctx, maxWidth, layoutable)

	if info.totalSize.X < maxWidth && l.GridFullWidth {
		remaining := maxWidth - info.totalSize.X
		for i := range info.widths {
			add := remaining * l.GridFullWidthWeights.GetWeight(i, info.columns)
			info.widths[i] += add
		}
	}
	if info.totalSize.Y < maxHeight && l.GridFullHeight {
		remaining := maxHeight - info.totalSize.Y
		for i := range info.heights {
			add := remaining * l.GridFullHeightWeights.GetWeight(i, info.rows)
			info.heights[i] += add
		}
	}

	offsetX := float32(0)
	offsetY := float32(0)
	for childIndex, child := range layoutable {
		col := childIndex % info.columns
		row := childIndex / info.columns
		sizing := info.sizings.Sizings[childIndex]
		cellWidth := info.widths[col]
		cellHeight := info.heights[row]

		if l.FullWidth || sizing.FullWidth > cellWidth {
			sizing.Width = cellWidth - sizing.WidthPadding()
			sizing.FullWidth = cellWidth
		}
		if l.FullHeight || sizing.FullHeight > cellHeight {
			sizing.Height = cellHeight - sizing.HeightPadding()
			sizing.FullHeight = cellHeight
		}

		alignX := l.HorizontalAlignment.Compute(cellWidth - sizing.FullWidth)
		alignY := l.VerticalAlignment.Compute(cellHeight - sizing.FullHeight)

		left := offsetX + alignX + sizing.OffsetX
		top := offsetY + alignY + sizing.OffsetY

		child.SetPlacement(Absolute(left, top, sizing.Width, sizing.Height))

		offsetX += cellWidth + info.spacingX
		if col == info.columns-1 {
			offsetY += cellHeight + info.spacingY
			offsetX = 0
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
}

type layoutInlineLine struct {
	start, endExclusive int
	width, height       float32
}

func (l LayoutInline) Init(b *Base) {}
func (l LayoutInline) getLines(ctx *RenderContext, maxWidth float32, layoutable []*Base) ([]layoutInlineLine, LayoutSizings) {
	spacingX := l.HorizontalSpacing.Get(ctx.AmountContext, true)
	sizings := getLayoutSizings(ctx, maxWidth, layoutable)
	lines := make([]layoutInlineLine, 0, len(layoutable))
	currentLine := layoutInlineLine{}

	for childIndex := range layoutable {
		sizing := sizings.Sizings[childIndex]
		nextWidth := currentLine.width + sizing.FullWidth
		if childIndex > currentLine.start {
			nextWidth += spacingX
		}
		if nextWidth > maxWidth && childIndex > currentLine.start {
			currentLine.endExclusive = childIndex
			lines = append(lines, currentLine)
			currentLine = layoutInlineLine{start: childIndex}
		}
		currentLine.width += sizing.FullWidth
		if childIndex > currentLine.start {
			currentLine.width += spacingX
		}
		currentLine.height = util.Max(currentLine.height, sizing.FullHeight)
	}

	if currentLine.width > 0 {
		currentLine.endExclusive = len(layoutable)
		lines = append(lines, currentLine)
	}

	return lines, sizings
}
func (l LayoutInline) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}

	if len(layoutable) == 0 {
		return size
	}

	spacingY := l.VerticalSpacing.Get(ctx.AmountContext, false)
	lines, _ := l.getLines(ctx, maxWidth, layoutable)

	for _, line := range lines {
		size.Y += line.height
		size.X = util.Max(size.X, line.width)
	}
	size.Y += spacingY * float32(len(lines)-1)

	return size
}
func (l LayoutInline) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	if len(layoutable) == 0 {
		return
	}

	maxWidth := bounds.Width()
	spacingX := l.HorizontalSpacing.Get(ctx.AmountContext, true)
	spacingY := l.VerticalSpacing.Get(ctx.AmountContext, false)
	lines, sizings := l.getLines(ctx, maxWidth, layoutable)
	offsetY := float32(0)

	for _, line := range lines {
		offsetX := util.Max(0, l.HorizontalAlignment.Compute(maxWidth-line.width))
		for i := line.start; i < line.endExclusive; i++ {
			child := layoutable[i]
			sizing := sizings.Sizings[i]
			alignY := l.VerticalAlignment.Compute(line.height - sizing.FullHeight)

			child.SetPlacement(Absolute(
				offsetX+sizing.OffsetX,
				offsetY+alignY+sizing.OffsetY,
				sizing.Width,
				sizing.Height,
			))
			offsetX += sizing.FullWidth + spacingX
		}
		offsetY += line.height + spacingY
	}
}

// A layout which uses the defined placements on each child but can help with enforcing
// their preferred size and keeping them in the parent. Perfect for open spaces with
// draggable components.
type LayoutStatic struct {
	EnforcePreferredSize    bool
	KeepInside              bool
	KeepInsideForgetSize    bool
	KeepInsideIgnoreMargins bool
}

func (l LayoutStatic) Init(b *Base) {}
func (l LayoutStatic) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32, layoutable []*Base) Coord {
	size := Coord{}

	if len(layoutable) == 0 {
		return size
	}

	minWidth := float32(0)
	for _, child := range layoutable {
		minSize := child.PreferredSize(ctx, minWidth)
		padding := child.Placement.Padding()
		if minSize.X+padding.X < maxWidth {
			minSize = child.PreferredSize(ctx, maxWidth-padding.X)
		}

		size = size.Max(Coord{
			X: minSize.X + padding.X,
			Y: minSize.Y + padding.Y,
		})
		minWidth = util.Max(minWidth, size.X)
	}

	return size
}
func (l LayoutStatic) Layout(b *Base, ctx *RenderContext, bounds Bounds, layoutable []*Base) {
	if len(layoutable) == 0 {
		return
	}

	if !l.EnforcePreferredSize || !l.KeepInside {
		return
	}

	width, height := bounds.Dimensions()

	for _, child := range layoutable {
		placement := child.Placement
		keepSize := !l.KeepInsideForgetSize

		if l.EnforcePreferredSize {
			// Get the placement width to be at least the min size
			minSize := child.PreferredSize(ctx, 0)
			placementBounds := placement.GetBoundsIn(bounds)
			if placementBounds.Width() < minSize.X {
				placement = placement.WithWidth(minSize.X)
				placementBounds = placement.GetBoundsIn(bounds)
				keepSize = true
			}
			// Once we're at the min size, what is the preferred height of the child based on
			// its current width.
			preferredSize := child.PreferredSize(ctx, placementBounds.Width())
			if placementBounds.Height() < preferredSize.Y {
				placement = placement.WithHeight(preferredSize.Y)
				keepSize = true
			}
		}

		if l.KeepInside {
			margins := child.Margin.GetBounds(ctx.AmountContext)

			fitInsideWidth := width
			fitInsideHeight := height
			if !l.KeepInsideIgnoreMargins {
				fitInsideWidth -= margins.Left + margins.Right
				fitInsideHeight -= margins.Top + margins.Bottom
				placement = placement.Shift(-margins.Left, -margins.Top)
			}
			placement = placement.FitInside(fitInsideWidth, fitInsideHeight, keepSize)
			if !l.KeepInsideIgnoreMargins {
				placement = placement.Shift(margins.Left, margins.Top)
			}

			// We can't keep size if any of this is true, but we're going to force it inside the parent.
			actual := placement.GetBoundsIn(bounds)
			if actual.Bottom > height {
				placement.Bottom.Base -= actual.Bottom - height
			}
			if actual.Top < 0 {
				placement.Top.Base -= actual.Top
			}
			if actual.Right > width {
				placement.Right.Base -= actual.Right - width
			}
			if actual.Left < 0 {
				placement.Left.Base -= actual.Left
			}
		}

		child.SetPlacement(placement)
	}
}

type LayoutSizing struct {
	FullWidth, FullHeight float32
	Width, Height         float32
	OffsetX, OffsetY      float32
}

func (ls LayoutSizing) HeightPadding() float32 { return ls.FullHeight - ls.Height }
func (ls LayoutSizing) WidthPadding() float32  { return ls.FullWidth - ls.Width }

func getLayoutSizing(ctx *RenderContext, width float32, child *Base) LayoutSizing {
	margin := child.Margin.GetBounds(ctx.AmountContext)
	extraWidth := margin.Left + margin.Right
	size := child.PreferredSize(ctx, width-extraWidth)
	return LayoutSizing{
		Width:      size.X,
		Height:     size.Y,
		FullWidth:  size.X + extraWidth,
		FullHeight: size.Y + margin.Top + margin.Bottom,
		OffsetX:    margin.Left,
		OffsetY:    margin.Top,
	}
}

type LayoutSizings struct {
	Sizings                 []LayoutSizing
	MaxWidth, MaxHeight     float32
	TotalWidth, TotalHeight float32
}

func getLayoutSizings(ctx *RenderContext, width float32, layoutable []*Base) (sizings LayoutSizings) {
	sizings.Sizings = make([]LayoutSizing, len(layoutable))

	for childIndex, child := range layoutable {
		sizing := getLayoutSizing(ctx, width, child)
		sizings.Sizings[childIndex] = sizing
		sizings.MaxWidth = util.Max(sizings.MaxWidth, sizing.FullWidth)
		sizings.MaxHeight = util.Max(sizings.MaxHeight, sizing.FullHeight)
		sizings.TotalWidth += sizing.FullWidth
		sizings.TotalHeight += sizing.FullHeight
	}

	return
}

func iteratorRange(count int, reverse bool) (start, end, move int) {
	start = 0
	end = count
	move = 1
	if reverse {
		start = count - 1
		end = -1
		move = -1
	}
	return
}
