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

	parent *Base
}

var _ Component = &Base{}

func (c *Base) Init(init Init) {
	c.Placement.Init(Maximized())
	if c.States == 0 {
		c.States = StateDefault
	}
	for i := range c.Layers {
		c.Layers[i].Init(init)
	}
	for i := range c.Children {
		c.Children[i].Init(init)
	}
}

func (c *Base) Place(parent Bounds) {
	c.Bounds = c.Placement.GetBoundsIn(parent)
	for i := range c.Layers {
		c.Layers[i].Place(c.Bounds)
	}
	for k := range c.Children {
		c.Children[k].Place(c.Bounds)
	}
}

func (c *Base) Update(update Update) {
	for i := range c.Layers {
		c.Layers[i].Update(update)
	}
	for i := range c.Children {
		c.Children[i].Update(update)
	}
}

func (c *Base) Render(out *UIVertexBuffer) {
	for _, layer := range c.Layers {
		if layer.ForStates(c.States) {
			layer.Render(out)
		}
	}
	for _, child := range c.Children {
		child.Render(out)
	}
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
	return c.States.Has(StateDisabled)
}

func (c *Base) OnPointer(ev *PointerEvent) {
	if c.IsDisabled() {
		return
	}
	switch ev.Type {
	case PointerEventEnter:
		c.States |= StateHover
	case PointerEventDown:
		c.States |= StatePressed
	case PointerEventLeave:
		c.States &= ^(StateHover | StateDragOver)
	case PointerEventUp:
		c.States &= ^(StatePressed | StateDragOver)
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
	c.States |= StateFocused
	if c.Events.OnFocus != nil {
		c.Events.OnFocus(ev)
	}
}

func (c Base) OnBlur(ev *ComponentEvent) {
	if c.IsDisabled() {
		return
	}
	c.States &= ^StateFocused
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
		c.States |= StateDragging
	case DragEventEnd:
		c.States &= ^StateDragging
	case DragEventOver:
		c.States |= StateDragOver
	}
	if c.Events.OnDrag != nil {
		c.Events.OnDrag(ev)
	}
}
