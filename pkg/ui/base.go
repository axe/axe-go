package ui

import "github.com/axe/axe-go/pkg/id"

type Base struct {
	Name      id.Identifier
	Layers    []Layer
	Placement Placement
	Bounds    Bounds
	Children  []*Base
	Focusable bool
	Draggable bool
	Droppable bool
	Events    Events
	States    State
	Clip      Placement

	TextStyles *TextStylesOverride

	OverShape    []Coord
	Transparency Watch[float32]

	dirty  Dirty
	parent *Base
	ui     *UI
}

var _ Component = &Base{}

func (c *Base) UI() *UI {
	return c.ui
}

func (c *Base) GetDirty() Dirty {
	return c.dirty
}

func (c *Base) Dirty(dirty Dirty) {
	if (dirty | c.dirty) != c.dirty {
		c.dirty.Add(dirty)

		if c.ui != nil {
			rootDirty := DirtyNone
			if dirty.Is(DirtyPlacement) {
				rootDirty.Add(DirtyDeepPlacement)
			}
			if dirty.Is(DirtyVisual) {
				rootDirty.Add(DirtyVisual)
			}
			c.ui.Root.Dirty(rootDirty)
		}
	}
}

func (c *Base) SetState(state State) {
	state.Remove(StateDefault)
	if state.None() {
		state.Add(StateDefault)
	}

	if c.isDirtyForState(state) {
		c.Dirty(DirtyVisual)
	}

	c.States = state
}

func (c *Base) isDirtyForState(state State) bool {
	if !c.dirty.Is(DirtyVisual) {
		for _, layer := range c.Layers {
			if layer.ForStates(state) != layer.ForStates(c.States) {
				return true
			}
		}
		if c.ui != nil && c.ui.Theme.StateModifier[state] != nil {
			return true
		}
	}
	return false
}

func (c *Base) RemoveStates(state State) {
	c.SetState(c.States.WithRemove(state))
}

func (c *Base) AddStates(state State) {
	c.SetState(c.States.WithAdd(state))
}

func (c *Base) SetPlacement(placement Placement) {
	if c.Placement != placement {
		c.Placement = placement
		c.Dirty(DirtyPlacement)
	}
}

func (c *Base) Init(init Init) {
	if c.ui != nil && !c.Name.Empty() {
		c.ui.Named.Set(c.Name, c)
	}
	c.Placement.Init(Maximized())
	if c.States == 0 {
		c.States = StateDefault
	}
	for i := range c.Layers {
		c.Layers[i].Init(init)
	}
	for _, child := range c.Children {
		child.parent = c
		child.ui = c.ui
		child.Init(init)
	}
	c.Dirty(DirtyPlacement)
}

func (c *Base) Place(parent Bounds, force bool) {
	doPlacement := force || c.dirty.Is(DirtyPlacement)
	if doPlacement {
		newBounds := c.Placement.GetBoundsIn(parent)
		if newBounds != c.Bounds {
			c.Bounds = newBounds
			for i := range c.Layers {
				c.Layers[i].Place(c.Bounds)
			}
			c.Dirty(DirtyVisual)
			force = true
		}
		c.dirty.Remove(DirtyPlacement)
	}
	for _, child := range c.Children {
		child.Place(c.Bounds, force)
	}
}

func (c *Base) Update(update Update) {
	for i := range c.Layers {
		c.Dirty(c.Layers[i].Update(update))
	}
	for _, child := range c.Children {
		child.Update(update)
	}

	c.CheckForChanges()
}

func (c *Base) CheckForChanges() {
	if c.Transparency.Cleaned() {
		c.Dirty(DirtyVisual)
	}
}

func (c *Base) Render(ctx *RenderContext, out *VertexBuffers) {
	if c.Transparency.Get() == 1 {
		return
	}

	baseCtx := ctx.WithBoundsAndTextStyles(c.Bounds, c.TextStyles)

	if len(c.Layers) > 0 {
		rendered := NewVertexIterator(out)

		for _, layer := range c.Layers {
			if layer.ForStates(c.States) {
				layer.Render(baseCtx, out)
			}
		}

		c.PostProcess(rendered)
	}

	if len(c.Children) > 0 {
		clipBounds := c.Clip.GetBoundsIn(c.Bounds)

		out.ClipMaybe(clipBounds, func(inner *VertexBuffers) {
			renderedChildren := NewVertexIterator(inner)

			for _, child := range c.Children {
				child.Render(baseCtx, inner)
			}

			c.PostProcess(renderedChildren)
		})
	}

	c.dirty.Remove(DirtyVisual)
}

func (c *Base) PostProcess(iter VertexIterator) {
	modifier := c.ui.Theme.StateModifier[c.States]
	if modifier != nil {
		for iter.HasNext() {
			modifier(iter.Next())
		}
	}

	if c.Transparency.Get() > 0 {
		iter.Reset()
		alphaMultiplier := 1 - c.Transparency.Get()
		for iter.HasNext() {
			v := iter.Next()
			v.Color.A *= alphaMultiplier
		}
	}
}

func (c *Base) Parent() Component {
	if c == nil || c.parent == nil {
		return nil
	}
	return c.parent
}

func (c *Base) At(pt Coord) Component {
	if !c.Bounds.Inside(pt) {
		return nil
	}

	if len(c.OverShape) > 0 {
		normalized := Coord{
			X: c.Bounds.Dx(pt.X),
			Y: c.Bounds.Dy(pt.Y),
		}
		if !inPolygon(c.OverShape, normalized) {
			return nil
		}
	}

	last := len(c.Children) - 1
	for i := last; i >= 0; i-- {
		at := c.Children[i].At(pt)
		if at != nil {
			return at
		}
	}

	return c
}

func (c *Base) Order() int {
	if c.parent == nil {
		return -1
	}
	return sliceIndexOf(c.parent.Children, c)
}

func (c *Base) SetOrder(order int) {
	currentOrder := c.Order()
	if currentOrder == -1 || order == currentOrder {
		return
	}
	sliceMove(c.parent.Children, currentOrder, order)
	c.Dirty(DirtyVisual)
}

func (c *Base) BringToFront() {
	if c.parent != nil {
		c.SetOrder(len(c.parent.Children) - 1)
	}
}

func (c *Base) SendToBack() {
	if c.parent != nil {
		c.SetOrder(0)
	}
}

func (c Base) IsFocusable() bool {
	return c.Focusable && !c.IsDisabled()
}

func (c Base) IsDraggable() bool {
	return c.Draggable && !c.IsDisabled()
}

func (c Base) IsDroppable() bool {
	return c.Droppable && !c.IsDisabled()
}

func (c Base) IsDisabled() bool {
	return c.States.Is(StateDisabled)
}

func (c *Base) SetDisabled(disabled bool) {
	if disabled {
		c.AddStates(StateDisabled)
	} else {
		c.RemoveStates(StateDisabled)
	}
}

func (c *Base) OnPointer(ev *PointerEvent) {
	if c.IsDisabled() {
		return
	}

	switch ev.Type {
	case PointerEventEnter:
		c.AddStates(StateHover)
	case PointerEventDown:
		c.AddStates(StatePressed)
	case PointerEventLeave:
		c.RemoveStates(StateHover | StateDragOver)
	case PointerEventUp:
		c.RemoveStates(StatePressed | StateDragOver)
	}

	if c.Events.OnPointer != nil {
		c.Events.OnPointer(ev)
	}
}

func (c *Base) OnKey(ev *KeyEvent) {
	if c.IsDisabled() {
		return
	}

	if c.Events.OnKey != nil {
		c.Events.OnKey(ev)
	}
}

func (c *Base) OnFocus(ev *Event) {
	if c.IsDisabled() {
		return
	}

	c.AddStates(StateFocused)

	if c.Events.OnFocus != nil {
		c.Events.OnFocus(ev)
	}
}

func (c Base) OnBlur(ev *Event) {
	if c.IsDisabled() {
		return
	}

	c.RemoveStates(StateFocused)

	if c.Events.OnBlur != nil {
		c.Events.OnBlur(ev)
	}
}

func (c Base) OnDrag(ev *DragEvent) {
	if c.IsDisabled() {
		return
	}

	switch ev.Type {
	case DragEventStart:
		c.AddStates(StateDragging)
	case DragEventEnd:
		c.RemoveStates(StateDragging)
	case DragEventOver:
		c.AddStates(StateDragOver)
	}

	if c.Events.OnDrag != nil {
		c.Events.OnDrag(ev)
	}
}
