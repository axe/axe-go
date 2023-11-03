package ui

import "time"

type Init struct {
	Theme *Theme
}

type Update struct {
	DeltaTime time.Duration
}

type RenderContext struct {
	AmountContext *AmountContext
	Theme         *Theme
	TextStyles    *TextStyles
}

func (ctx *RenderContext) WithAmountContext(amt *AmountContext) *RenderContext {
	copy := *ctx
	copy.AmountContext = amt
	return &copy
}

func (ctx *RenderContext) WithTextStyles(styles *TextStyles) *RenderContext {
	copy := *ctx
	copy.TextStyles = styles
	return &copy
}

func (ctx *RenderContext) WithParent(width, height float32) *RenderContext {
	if ctx.AmountContext.IsSameSize(width, height, ctx.TextStyles.FontSize) {
		return ctx
	}
	return ctx.WithAmountContext(ctx.AmountContext.Resize(width, height, ctx.TextStyles.FontSize))
}

func (ctx *RenderContext) WithBounds(parent Bounds) *RenderContext {
	return ctx.WithParent(parent.Width(), parent.Height())
}

func (ctx *RenderContext) WithBoundsAndTextStyles(parent Bounds, styles *TextStylesOverride) *RenderContext {
	width, height := parent.Dimensions()
	fontSize := ctx.TextStyles.FontSize
	if styles != nil && styles.FontSize != nil {
		fontSize = *styles.FontSize
	}
	if ctx.AmountContext.IsSameSize(width, height, fontSize) && !styles.HasOverride() {
		return ctx
	}
	copy := *ctx
	copy.TextStyles = ctx.TextStyles.Override(styles)
	copy.AmountContext = ctx.AmountContext.Resize(width, height, copy.TextStyles.FontSize)
	return &copy
}

type Component interface {
	Init(init Init)
	Place(parent Bounds, force bool)
	Update(update Update)
	Render(ctx *RenderContext, out *VertexBuffers)
	GetDirty() Dirty
	Dirty(dirty Dirty)
	Parent() Component
	At(pt Coord) Component

	IsFocusable() bool
	IsDraggable() bool
	IsDroppable() bool

	OnPointer(ev *PointerEvent)
	OnKey(ev *KeyEvent)
	OnFocus(ev *Event)
	OnBlur(ev *Event)
	OnDrag(ev *DragEvent)
}

type ComponentMap map[uintptr]Component

func (cm ComponentMap) Add(c Component) {
	cm[toPtr(c)] = c
}

func (cm ComponentMap) AddMany(c []Component) {
	for _, m := range c {
		cm.Add(m)
	}
}

func (cm ComponentMap) Has(c Component) bool {
	_, exists := cm[toPtr(c)]
	return exists
}

func (cm ComponentMap) AddLineage(c Component) {
	curr := c
	for curr != nil {
		cm.Add(curr)
		curr = curr.Parent()
	}
}

func (old ComponentMap) Compare(new ComponentMap) (inOld []Component, inBoth []Component, inNew []Component) {
	inOld = make([]Component, 0)
	inBoth = make([]Component, 0)
	inNew = make([]Component, 0)

	for _, oldOverAncestor := range old {
		if !new.Has(oldOverAncestor) {
			inOld = append(inOld, oldOverAncestor)
		} else {
			inBoth = append(inBoth, oldOverAncestor)
		}
	}
	for _, newOverAncestor := range new {
		if !old.Has(newOverAncestor) {
			inNew = append(inNew, newOverAncestor)
		}
	}
	return
}
