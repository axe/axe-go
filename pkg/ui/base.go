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

	dirty  Dirty
	parent *Base
	root   *Base
}

var _ Component = &Base{}

func (c *Base) GetDirty() Dirty {
	return c.dirty
}

func (c *Base) Dirty(dirty Dirty) {
	if (dirty | c.dirty) != c.dirty {
		c.dirty.Add(dirty)

		if c.root != nil {
			rootDirty := DirtyNone
			if dirty.Is(DirtyPlacement) {
				rootDirty.Add(DirtyDeepPlacement)
			}
			if dirty.Is(DirtyVisual) {
				rootDirty.Add(DirtyVisual)
			}
			c.root.Dirty(rootDirty)
		}
	}
}

func (c *Base) SetState(state State) {
	if !c.dirty.Is(DirtyVisual) {
		for _, layer := range c.Layers {
			if layer.ForStates(state) != layer.ForStates(c.States) {
				c.Dirty(DirtyVisual)
				break
			}
		}
	}

	c.States = state
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
	if c.root == nil {
		c.root = c
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
		child.root = c.root
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

func (c *Base) Render(ctx AmountContext, out *VertexBuffer) {
	baseCtx := ctx.WithParent(c.Bounds.Width(), c.Bounds.Height())

	for _, layer := range c.Layers {
		if layer.ForStates(c.States) {
			layer.Render(baseCtx, out)
		}
	}

	for _, child := range c.Children {
		child.Render(baseCtx, out)
	}

	c.dirty.Remove(DirtyVisual)
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

func (c *Base) OnFocus(ev *ComponentEvent) {
	if c.IsDisabled() {
		return
	}

	c.AddStates(StateFocused)

	if c.Events.OnFocus != nil {
		c.Events.OnFocus(ev)
	}
}

func (c Base) OnBlur(ev *ComponentEvent) {
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
