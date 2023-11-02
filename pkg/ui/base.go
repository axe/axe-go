package ui

type Base struct {
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

	dirty  Dirty
	parent *Base
	ui     *UI
}

var _ Component = &Base{}

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
}

func (c *Base) Render(ctx RenderContext, out *VertexBuffers) {
	baseCtx := ctx.WithBounds(c.Bounds)

	if len(c.Layers) > 0 {
		rendered := NewVertexIterator(out)

		for _, layer := range c.Layers {
			if layer.ForStates(c.States) {
				layer.Render(baseCtx, out)
			}
		}

		c.PostProcess(ctx, rendered)
	}

	if len(c.Children) > 0 {
		clipBounds := c.Clip.GetBoundsIn(c.Bounds)

		out.ClipMaybe(clipBounds, func(inner *VertexBuffers) {
			renderedChildren := NewVertexIterator(inner)

			for _, child := range c.Children {
				child.Render(baseCtx, inner)
			}

			c.PostProcess(ctx, renderedChildren)
		})
	}

	c.dirty.Remove(DirtyVisual)
}

func (c *Base) PostProcess(ctx RenderContext, iter VertexIterator) {
	if len(ctx.Theme.StateModifier) == 0 {
		return
	}

	modifier := ctx.Theme.StateModifier[c.States]
	if modifier != nil {
		for iter.HasNext() {
		}
	}
	/*
		states := c.States
		for one := states.Take(); one != 0; one = states.Take() {
			modifier := ctx.Theme.StateModifier[one]
			if modifier != nil {
				for iter.HasNext() {
					modifier(iter.Next())
				}
				iter.Reset()
			}
		}
	*/
}

func (c *Base) Parent() Component {
	if c == nil || c.parent == nil {
		return nil
	}
	return c.parent
}

func (c *Base) At(pt Coord) Component {
	if !c.Bounds.Contains(pt) {
		return nil
	}
	for i := range c.Children {
		at := c.Children[i].At(pt)
		if at != nil {
			return at
		}
	}
	return c
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
