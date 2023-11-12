package ui

import "time"

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
	// Should be called after ui is set, and if it's going to have a parent - it's parent
	Init()
	Place(ctx *RenderContext, parent Bounds, force bool)
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

type ComponentMap struct {
	set     map[uintptr]*Base
	ordered []*Base
}

func NewComponentMap() ComponentMap {
	return ComponentMap{
		set: make(map[uintptr]*Base),
	}
}

func (cm ComponentMap) Components() []*Base {
	return cm.ordered
}

func (cm *ComponentMap) Add(c *Base) {
	key := toPtr(c)
	if _, exists := cm.set[key]; !exists {
		cm.set[key] = c
		cm.ordered = append(cm.ordered, c)
	}
}

func (cm *ComponentMap) AddMany(c []*Base) {
	for _, m := range c {
		cm.Add(m)
	}
}

func (cm ComponentMap) Has(c *Base) bool {
	_, exists := cm.set[toPtr(c)]
	return exists
}

func (cm *ComponentMap) AddLineage(c *Base) {
	curr := c
	for curr != nil {
		cm.Add(curr)
		curr = curr.Parent()
	}
}

func (old ComponentMap) Compare(new ComponentMap) (inOld []*Base, inBoth []*Base, inNew []*Base) {
	inOld = make([]*Base, 0)
	inBoth = make([]*Base, 0)
	inNew = make([]*Base, 0)

	for _, oldOverAncestor := range old.ordered {
		if !new.Has(oldOverAncestor) {
			inOld = append(inOld, oldOverAncestor)
		} else {
			inBoth = append(inBoth, oldOverAncestor)
		}
	}
	for _, newOverAncestor := range new.ordered {
		if !old.Has(newOverAncestor) {
			inNew = append(inNew, newOverAncestor)
		}
	}
	return
}
