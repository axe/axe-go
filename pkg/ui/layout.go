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

// Layout information that can be specified at the component level that a layout may
// or may not use during placement.
type LayoutData struct {
	// The weight for the width when the layout is stretched.
	WidthWeight float32
	// Each cell component is expanded to the column width
	FullWidth *bool
	// When FullWidth is false, how to align the cell component horizontally in the cell
	HorizontalAlignment *Alignment
	// The weight for the height when the layout is stretched.
	HeightWeight float32
	// Each cell component is expanded to the row height
	FullHeight *bool
	// When FullHeight is false, how to align the cell component veritcally in the cell
	VerticalAlignment *Alignment
	// For certain layouts, should the preferred size be enforced?
	EnforcePreferredSize *bool
	// For certain layouts, should the component be moved within the parent's bounds?
	KeepInside              *bool
	KeepInsideForgetSize    *bool
	KeepInsideIgnoreMargins *bool
}

func (ld LayoutData) WithWidthWeight(widthWeight float32) LayoutData {
	ld.WidthWeight = widthWeight
	return ld
}

func (ld LayoutData) WithFullWidth(fullWidth bool) LayoutData {
	ld.FullWidth = &fullWidth
	return ld
}

func (ld LayoutData) WithHorizontalAlignment(alignment Alignment) LayoutData {
	ld.HorizontalAlignment = &alignment
	return ld
}

func (ld LayoutData) WithHeightWeight(heightWeight float32) LayoutData {
	ld.HeightWeight = heightWeight
	return ld
}

func (ld LayoutData) WithFullHeight(fullHeight bool) LayoutData {
	ld.FullHeight = &fullHeight
	return ld
}

func (ld LayoutData) WithVerticalAlignment(alignment Alignment) LayoutData {
	ld.VerticalAlignment = &alignment
	return ld
}

func (ld LayoutData) WithEnforcePreferredSize(enforce bool) LayoutData {
	ld.EnforcePreferredSize = &enforce
	return ld
}

func (ld LayoutData) WithKeepInside(keepInside bool) LayoutData {
	ld.KeepInside = &keepInside
	return ld
}

func (ld LayoutData) WithKeepInsideIgnoreMargins(ignoreMargins bool) LayoutData {
	ld.KeepInsideIgnoreMargins = &ignoreMargins
	return ld
}

func (ld LayoutData) WithKeepInsideForgetSize(forgetSize bool) LayoutData {
	ld.KeepInsideForgetSize = &forgetSize
	return ld
}

// A layout which places all children in a column (vertical stack).
// The children can have their width expanded to the full width or default
// to their preferred size. If they are not expanded to full width they can
// be horizontally aligned. Optionally the size inferred by the placement
// can be factored into the preferred size as well (it's not by default).
type LayoutColumn struct {
	FullWidth           bool
	FullHeight          bool
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
	width, maxHeight := bounds.Dimensions()
	sizings := getLayoutSizings(ctx, width, layoutable)
	spacing := l.Spacing.Get(ctx.AmountContext, false)
	spacingTotal := spacing * float32(n-1)
	totalHeight := sizings.TotalHeight + spacingTotal

	if l.FullHeight && maxHeight > totalHeight {
		extraHeight := maxHeight - totalHeight
		for childIndex := range sizings.Sizings {
			sizing := &sizings.Sizings[childIndex]
			weight := sizings.GetHeightWeight(childIndex)
			add := extraHeight * weight
			sizing.FullHeight += add
			sizing.Height += add
		}
	}

	for childIndex, child := range layoutable {
		sizing := sizings.Sizings[childIndex]
		if util.Coalesce(sizing.Data.FullWidth, l.FullWidth) {
			sizing.SetFullWidth(width)
		} else if l.EqualWidths {
			sizing.SetFullWidth(sizings.MaxWidth)
		}
		halign := util.Coalesce(sizing.Data.HorizontalAlignment, l.HorizontalAlignment)
		placement := child.Placement.
			Attach(float32(halign), 0, sizing.Width, sizing.Height).
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
				weight := minSizings.GetWidthWeight(i)
				add := remaining * weight
				sizing.Width += add
				sizing.FullWidth += add
			}
		}
	}

	offsetX := float32(0)

	for childIndex, child := range layoutable {
		sizing := sizings[childIndex]
		if util.Coalesce(sizing.Data.FullHeight, l.FullHeight) {
			sizing.SetFullHeight(boundsHeight)
		} else if l.EqualHeights {
			sizing.SetFullHeight(maxHeight)
		}
		alignment := util.Coalesce(sizing.Data.VerticalAlignment, l.VerticalAlignment)
		placement := child.Placement.
			Attach(0, float32(alignment), sizing.Width, sizing.Height).
			Shift(offsetX+sizing.OffsetX, sizing.OffsetY)
		child.SetPlacement(placement)
		offsetX += sizing.FullWidth + spacing
	}
}

// A definition for a single column in the table layout.
type LayoutGridColumn struct {
	// When stretching the columns to match the width of the bounds - how should each column
	// be stretched? A value of zero will cause no stretching, and non-zero values will be
	// divided up based on the value / the sum of values. One or zero weights will cause
	// the widths to be equally stretched.
	Weight float32
	// Each cell component is expanded to the column width
	FullWidth bool
	// When FullWidth is false, how to align the cell component horizontally in the cell
	HorizontalAlignment Alignment
	// Each cell component is expanded to the row height
	FullHeight *bool
	// When FullHeight is false, how to align the cell component vertically in the cell
	VerticalAlignment *Alignment
	// Defines the min widths for one or all columns.
	Min float32
	// Defines the max widths for one or all columns.
	Max float32
}

// A definition for a single row in the table layout.
type LayoutGridRow struct {
	// When stretching the rows to match the height of the bounds - how should each row
	// be stretched? A value of zero will cause no stretching, and non-zero values will be
	// divided up based on the value / the sum of values. One or zero weights will cause
	// the heights to be equally stretched.
	Weight float32
	// Each cell component is expanded to the row height
	FullHeight bool
	// When FullHeight is false, how to align the cell component veritcally in the cell
	VerticalAlignment Alignment
	// Each cell component is expanded to the column width
	FullWidth *bool
	// When FullWidth is false, how to align the cell component horizontally in the cell
	HorizontalAlignment *Alignment
	// Defines the min heights for one or all rows. If the number of rows extends beyond the number
	// of min heights defined, they will use the last min height defined.
	Min float32
	// Defines the max heights for one or all rows. If the number of rows extends beyond the number
	// of max heights defined, they will use the last max height defined.
	Max float32
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
	FullWidth bool
	// The grid will take up the full height of the bounds. May result in ignoring other height settings.
	FullHeight bool
	// How much space between rows
	VerticalSpacing Amount
	// How much space between columns
	HorizontalSpacing Amount
	// Whether the number of columns should be dynamically calculated.
	ColumnsDynamic bool
	// When dynamically calculating columns, this will be the min number of columns in the grid.
	ColumnsMin int
	// When dynamically calculating columns, this will be the max number of columns in the grid.
	ColumnsMax int
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
	// The columns in the grid. You can use Columns or ColumnsMin to control the number of columns in the grid.
	// When the number of column definitions is below the number of columns the last column definition will be used.
	// At least one column is typically defined.
	Columns []LayoutGridColumn
	// The rows in the table. If the table has anymore rows beyond what's defined the last row will be used.
	// At least one row is typically defined.
	Rows []LayoutGridRow
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
	heights       []float32
	heightWeights []float32
	widths        []float32
	widthWeights  []float32
	totalSpacing  Coord
	totalSize     Coord
	rows          int
	columns       int
	totalWeight   Coord
}

func (lgs layoutGridSizings) getWidthWeight(col int) float32 {
	if lgs.totalWeight.X != 0 {
		return lgs.widthWeights[col] / lgs.totalWeight.X
	} else {
		return 1.0 / float32(len(lgs.widthWeights))
	}
}
func (lgs layoutGridSizings) getHeightWeight(row int) float32 {
	if lgs.totalWeight.Y != 0 {
		return lgs.heightWeights[row] / lgs.totalWeight.Y
	} else {
		return 1.0 / float32(len(lgs.heightWeights))
	}
}

func (info layoutGridInfo) getSizingsFor(columns int, grid *LayoutGrid) layoutGridSizings {
	last := len(info.sizings.Sizings) - 1
	rows := (last + columns) / columns
	sizings := layoutGridSizings{
		columns:       columns,
		rows:          rows,
		widths:        make([]float32, columns),
		widthWeights:  make([]float32, columns),
		heights:       make([]float32, rows),
		heightWeights: make([]float32, rows),
		totalSpacing: Coord{
			X: float32(columns-1) * info.spacingX,
			Y: float32(rows-1) * info.spacingY,
		},
	}
	for i := range sizings.widths {
		col := grid.columnAt(i)
		sizings.widths[i] = col.Min
		sizings.widthWeights[i] = col.Weight
	}
	for i := range sizings.heights {
		row := grid.rowAt(i)
		sizings.heights[i] = row.Min
		sizings.heightWeights[i] = row.Weight
	}
	for childIndex, sizing := range info.sizings.Sizings {
		col := childIndex % columns
		row := childIndex / columns
		sizings.widths[col] = util.Max(sizings.widths[col], sizing.FullWidth)
		sizings.widthWeights[col] = util.Max(sizings.widthWeights[col], sizing.Data.WidthWeight)
		sizings.heights[row] = util.Max(sizings.heights[row], sizing.FullHeight)
		sizings.heightWeights[row] = util.Max(sizings.heightWeights[row], sizing.Data.HeightWeight)
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
	for i := range sizings.widths {
		if maxWidth := grid.columnAt(i).Max; maxWidth > 0 {
			sizings.widths[i] = util.Min(sizings.widths[i], maxWidth)
		}
	}
	for i := range sizings.heights {
		if maxHeight := grid.rowAt(i).Max; maxHeight > 0 {
			sizings.heights[i] = util.Min(sizings.heights[i], maxHeight)
		}
	}
	sizings.totalSize = sizings.totalSpacing
	for _, columnWidth := range sizings.widths {
		sizings.totalSize.X += columnWidth
	}
	for _, rowHeight := range sizings.heights {
		sizings.totalSize.Y += rowHeight
	}
	for _, columnWeight := range sizings.widthWeights {
		sizings.totalWeight.X += columnWeight
	}
	for _, rowWeight := range sizings.heightWeights {
		sizings.totalWeight.Y += rowWeight
	}

	return sizings
}

func (l LayoutGrid) columnAt(i int) *LayoutGridColumn {
	last := len(l.Columns) - 1
	if i <= last {
		return &l.Columns[i]
	} else if last == -1 {
		var empty LayoutGridColumn
		return &empty
	} else {
		return &l.Columns[last]
	}
}

func (l LayoutGrid) rowAt(i int) *LayoutGridRow {
	last := len(l.Rows) - 1
	if i <= last {
		return &l.Rows[i]
	} else if last == -1 {
		var empty LayoutGridRow
		return &empty
	} else {
		return &l.Rows[last]
	}
}

func (l LayoutGrid) Init(b *Base) {}
func (l *LayoutGrid) getSizingInfo(ctx *RenderContext, maxWidth float32, layoutable []*Base) layoutGridInfo {
	info := layoutGridInfo{
		spacingX: l.HorizontalSpacing.Get(ctx.AmountContext, true),
		spacingY: l.VerticalSpacing.Get(ctx.AmountContext, false),
	}

	columns := len(l.Columns)
	if !l.ColumnsDynamic {
		columns = util.Max(columns, l.ColumnsMin)
		if l.EqualWidths {
			// equal width with defined number of columns is easiest, we calculate column width add calculate sizings based on that.
			columnWidth := ((maxWidth + info.spacingX) / float32(columns)) - info.spacingX
			info.sizings = getLayoutSizings(ctx, columnWidth, layoutable)
		} else {
			// we look at the min size, find out how much "extra" space we have to work with, then divide it up evenly.
			info.sizings = getLayoutSizings(ctx, 0, layoutable)
			minWidths := make([]float32, columns)
			colWeights := make([]float32, columns)
			for col := range minWidths {
				c := l.columnAt(col)
				minWidths[col] = c.Min
				colWeights[col] = c.Weight
			}
			for childIndex, sizing := range info.sizings.Sizings {
				col := childIndex % columns
				minWidths[col] = util.Max(minWidths[col], sizing.FullWidth)
				colWeights[col] = util.Max(colWeights[col], sizing.Data.WidthWeight)
			}
			totalMinWidths := float32(0)
			for _, minWidth := range minWidths {
				totalMinWidths += minWidth
			}
			totalMinWidths += info.spacingX * float32(columns-1)
			extra := maxWidth - totalMinWidths
			if extra > 0 {
				totalColWeights := float32(0)
				for _, weight := range colWeights {
					totalColWeights += weight
				}
				if totalColWeights == 0 {
					totalColWeights = float32(columns)
					for i := range colWeights {
						colWeights[i] = 1
					}
				}
				for childIndex, sizing := range info.sizings.Sizings {
					col := childIndex % columns
					columnWidth := minWidths[col]
					if sizing.FullWidth < columnWidth {
						extraSize := extra * (colWeights[col] / totalColWeights)
						info.sizings.Sizings[childIndex] = getLayoutSizing(ctx, columnWidth+extraSize, layoutable[childIndex])
					}
				}
			}
		}
	} else {
		// we calculate columns based on max column width, and with the real
		// column width we calculate we determine sizings.
		layoutWidth := maxWidth
		if columns == 0 {
			layoutWidth = l.ColumnsDynamicDelta * maxWidth
		}
		info.sizings = getLayoutSizings(ctx, layoutWidth, layoutable)
		maxChildWidth := float32(0)
		for _, col := range l.Columns {
			maxChildWidth = util.Max(maxChildWidth, col.Min)
		}
		for _, sizing := range info.sizings.Sizings {
			maxChildWidth = util.Max(maxChildWidth, sizing.FullWidth)
		}
		columns = util.Max(l.ColumnsMin, int((maxWidth+info.spacingX)/(maxChildWidth+info.spacingX)))
		if columns == 0 {
			columns = 1
		}
		if l.ColumnsMax > 0 && columns > l.ColumnsMax {
			columns = l.ColumnsMax
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

	if info.totalSize.X < maxWidth && l.FullWidth {
		remaining := maxWidth - info.totalSize.X
		for i := range info.widths {
			add := remaining * info.getWidthWeight(i)
			info.widths[i] += add
		}
	}
	if info.totalSize.Y < maxHeight && l.FullHeight {
		remaining := maxHeight - info.totalSize.Y
		for i := range info.heights {
			add := remaining * info.getHeightWeight(i)
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
		colLayout := l.columnAt(col)
		rowLayout := l.rowAt(row)

		if util.Coalesce2(sizing.Data.FullWidth, rowLayout.FullWidth, colLayout.FullWidth) || sizing.FullWidth > cellWidth {
			sizing.SetFullWidth(cellWidth)
		}
		if util.Coalesce2(sizing.Data.FullHeight, colLayout.FullHeight, rowLayout.FullHeight) || sizing.FullHeight > cellHeight {
			sizing.SetFullHeight(cellHeight)
		}

		halign := util.Coalesce2(sizing.Data.HorizontalAlignment, rowLayout.HorizontalAlignment, colLayout.HorizontalAlignment)
		valign := util.Coalesce2(sizing.Data.VerticalAlignment, colLayout.VerticalAlignment, rowLayout.VerticalAlignment)

		alignX := halign.Compute(cellWidth - sizing.FullWidth)
		alignY := valign.Compute(cellHeight - sizing.FullHeight)

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
	// The default vertical alignment of a component in a line.
	LineVerticalAlignment Alignment
	// The default horizontal alignment of a component in a stretched out
	// line in the parent component. This is only used when FullWidth=true
	// and LineFullWidth=false.
	LineHorizontalAlignment Alignment
	// The alignment of all the lines in the parent component.
	VerticalAlignment Alignment
	// The alignment of all the lines in the parent component.
	HorizontalAlignment Alignment
	// The spacing between each line.
	VerticalSpacing Amount
	// The spacing between each component in a line.
	HorizontalSpacing Amount
	// If all the components in a given line share the same height.
	LineFullHeight bool
	// If all the components in a given line take up the full width available.
	// This only has an effect when FullWidth is true.
	LineFullWidth bool
	// If all the lines share the same height.
	EqualHeights bool
	// If all components in a line should share the same width
	LineEqualWidths bool
	// If the layout should stretch the lines to the height of the parent.
	// The height weights of the components are used to stretch and if none are
	// defined then they are all equally stretched.
	FullHeight bool
	// If the layout should stretch the lines to the width of the parent.
	// The width weights of the components are used to stretch and if none are
	// defined then they are all equally stretched.
	FullWidth bool
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

	if l.EqualHeights {
		maxHeight := float32(0)
		for _, line := range lines {
			maxHeight = util.Max(maxHeight, line.height)
		}
		for i := range lines {
			lines[i].height = maxHeight
		}
	}
	if l.LineEqualWidths {
		for _, line := range lines {
			maxWidth := float32(0)
			for i := line.start; i < line.endExclusive; i++ {
				maxWidth = util.Max(maxWidth, sizings.Sizings[i].FullWidth)
			}
			for i := line.start; i < line.endExclusive; i++ {
				sizings.Sizings[i].SetFullWidth(maxWidth)
			}
		}
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
	n := len(layoutable)
	if n == 0 {
		return
	}

	maxWidth, maxHeight := bounds.Dimensions()
	spacingX := l.HorizontalSpacing.Get(ctx.AmountContext, true)
	spacingY := l.VerticalSpacing.Get(ctx.AmountContext, false)
	lines, sizings := l.getLines(ctx, maxWidth, layoutable)
	offsetY := float32(0)

	totalHeight := spacingY * float32(n-1)
	for _, line := range lines {
		totalHeight += line.height
	}

	if !l.FullHeight {
		offsetY = l.VerticalAlignment.Compute(maxHeight - totalHeight)
	}

	if l.FullWidth {
		for _, line := range lines {
			extraWidth := maxWidth - line.width
			if extraWidth < 0 {
				continue
			}
			totalWidthWeight := float32(0)
			for i := line.start; i < line.endExclusive; i++ {
				sizing := sizings.Sizings[i]
				totalWidthWeight += sizing.Data.WidthWeight
			}
			equalExtraWidth := extraWidth / float32(line.endExclusive-line.start)
			offsetX := float32(0)
			for i := line.start; i < line.endExclusive; i++ {
				sizing := &sizings.Sizings[i]
				add := equalExtraWidth
				if totalWidthWeight > 0 {
					add = extraWidth * (sizing.Data.WidthWeight / totalWidthWeight)
				}
				if l.LineFullWidth {
					sizing.Width += add
					sizing.FullWidth += add
				} else {
					halign := util.Coalesce(sizing.Data.HorizontalAlignment, l.LineHorizontalAlignment)
					sizing.OffsetX += offsetX + halign.Compute(add)
					offsetX += add
				}
			}
		}
	}

	extraHeight := maxHeight - totalHeight
	if l.FullHeight && extraHeight > 0 {
		totalHeightWeight := float32(0)
		lineMaxWeight := make([]float32, len(lines))
		for lineIndex, line := range lines {
			for i := line.start; i < line.endExclusive; i++ {
				sizing := sizings.Sizings[i]
				lineMaxWeight[lineIndex] = util.Max(lineMaxWeight[lineIndex], sizing.Data.HeightWeight)
			}
			totalHeightWeight += lineMaxWeight[lineIndex]
		}
		if totalHeightWeight == 0 {
			equalExtraHeight := extraHeight / float32(len(lines))
			for i := range lines {
				lines[i].height += equalExtraHeight
			}
		} else {
			for i := range lines {
				add := extraHeight * (lineMaxWeight[i] / totalHeightWeight)
				lines[i].height += add
			}
		}
	}

	for _, line := range lines {
		offsetX := float32(0)
		if !l.FullWidth {
			offsetX = util.Max(0, l.HorizontalAlignment.Compute(maxWidth-line.width))
		}
		for i := line.start; i < line.endExclusive; i++ {
			child := layoutable[i]
			sizing := sizings.Sizings[i]

			if util.Coalesce(sizing.Data.FullHeight, l.LineFullHeight) {
				sizing.SetFullHeight(line.height)
			}

			valign := util.Coalesce(sizing.Data.VerticalAlignment, l.LineVerticalAlignment)
			alignY := valign.Compute(line.height - sizing.FullHeight)

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

	width, height := bounds.Dimensions()

	for _, child := range layoutable {
		placement := child.Placement
		keepSize := !util.Coalesce(child.LayoutData.KeepInsideForgetSize, l.KeepInsideForgetSize)

		if util.Coalesce(child.LayoutData.EnforcePreferredSize, l.EnforcePreferredSize) {
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

		if util.Coalesce(child.LayoutData.KeepInside, l.KeepInside) {
			margins := child.Margin.GetBounds(ctx.AmountContext)

			fitInsideWidth := width
			fitInsideHeight := height
			if !util.Coalesce(child.LayoutData.KeepInsideIgnoreMargins, l.KeepInsideIgnoreMargins) {
				fitInsideWidth -= margins.Left + margins.Right
				fitInsideHeight -= margins.Top + margins.Bottom
				placement = placement.Shift(-margins.Left, -margins.Top)
			}
			placement = placement.FitInside(fitInsideWidth, fitInsideHeight, keepSize)
			if !util.Coalesce(child.LayoutData.KeepInsideIgnoreMargins, l.KeepInsideIgnoreMargins) {
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
	Data                  *LayoutData
}

func (ls LayoutSizing) HeightPadding() float32 { return ls.FullHeight - ls.Height }
func (ls LayoutSizing) WidthPadding() float32  { return ls.FullWidth - ls.Width }
func (ls *LayoutSizing) SetFullWidth(fullWidth float32) {
	ls.Width = fullWidth - ls.WidthPadding()
	ls.FullWidth = fullWidth
}
func (ls *LayoutSizing) SetFullHeight(fullHeight float32) {
	ls.Height = fullHeight - ls.HeightPadding()
	ls.FullHeight = fullHeight
}

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
		Data:       &child.LayoutData,
	}
}

type LayoutSizings struct {
	Sizings                 []LayoutSizing
	MaxWidth, MaxHeight     float32
	TotalWidth, TotalHeight float32
	TotalWidthWeight        float32
	TotalHeightWeight       float32
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
		sizings.TotalWidthWeight += child.LayoutData.WidthWeight
		sizings.TotalHeightWeight += child.LayoutData.HeightWeight
	}

	return
}

func (ls LayoutSizings) GetHeightWeight(i int) float32 {
	weight := ls.Sizings[i].Data.HeightWeight
	if ls.TotalHeightWeight > 0 {
		weight /= ls.TotalHeightWeight
	} else {
		weight = 1.0 / float32(len(ls.Sizings))
	}
	return weight
}

func (ls LayoutSizings) GetWidthWeight(i int) float32 {
	weight := ls.Sizings[i].Data.WidthWeight
	if ls.TotalWidthWeight > 0 {
		weight /= ls.TotalWidthWeight
	} else {
		weight = 1.0 / float32(len(ls.Sizings))
	}
	return weight
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
