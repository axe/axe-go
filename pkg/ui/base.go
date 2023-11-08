package ui

import (
	"github.com/axe/axe-go/pkg/id"
)

type Base struct {
	Name         id.Identifier
	Layers       []Layer
	Placement    Placement
	Bounds       Bounds
	Children     []*Base
	Focusable    bool
	Draggable    bool
	Droppable    bool
	Events       Events
	States       State
	Clip         Placement
	Transform    Transform
	TextStyles   *TextStylesOverride
	Shape        Shape
	OverShape    []Coord
	Transparency Watch[float32]
	Animation    AnimationState
	Animations   *Animations
	Cursors      Cursors
	Hooks        Hooks

	Margin  AmountBounds
	MinSize Coord
	MaxSize Coord
	Layout  Layout

	visualBounds  Bounds
	dirty         Dirty
	parent        *Base
	ui            *UI
	lastTransform Transform
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
		c.Dirty(DirtyVisual | DirtyPlacement)
	}

	c.States = state
}

func (c *Base) isDirtyForState(state State) bool {
	if !c.dirty.Is(DirtyVisual | DirtyPlacement) {
		for _, layer := range c.Layers {
			if layer.ForStates(state) != layer.ForStates(c.States) {
				return true
			}
		}
		if c.ui != nil && c.ui.Theme.StatePostProcess[state] != nil {
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

func (c *Base) GetPlacement() Placement {
	return c.Placement
}

func (b *Base) Init(init Init) {
	if b.ui != nil && !b.Name.Empty() {
		b.ui.Named.Set(b.Name, b)
	}
	b.Placement.Init(Maximized())
	if b.States == 0 {
		b.States = StateDefault
	}
	if b.Shape != nil {
		b.Shape.Init(init)
	}
	for i := range b.Layers {
		b.Layers[i].Init(b, init)
	}
	if b.Layout != nil {
		b.Layout.Init(b, init)
	}
	for _, child := range b.Children {
		child.parent = b
		child.ui = b.ui
		child.Init(init)
	}
	b.Dirty(DirtyPlacement)
	b.PlayEvent(AnimationEventShow)

	if b.Hooks.OnInit != nil {
		b.Hooks.OnInit(b, init)
	}
}

func (b *Base) Place(ctx *RenderContext, parent Bounds, force bool) {
	doPlacement := force || b.dirty.Is(DirtyPlacement)
	if doPlacement {
		newBounds := b.Placement.GetBoundsIn(parent)
		if newBounds != b.Bounds {
			b.Bounds = newBounds
			for i := range b.Layers {
				b.Layers[i].Place(b, b.Bounds)
			}
			b.Dirty(DirtyVisual)
			force = true
		}
		b.dirty.Remove(DirtyPlacement)
	}

	baseCtx := ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)

	if b.Layout != nil {
		b.Layout.Layout(b, baseCtx, b.Bounds, b.Children)
	}

	for _, child := range b.Children {
		child.Place(baseCtx, b.Bounds, force)
	}

	if b.Hooks.OnPlace != nil {
		b.Hooks.OnPlace(b, parent, baseCtx, force)
	}
}

func (b *Base) PreferredSize(ctx *RenderContext, ignorePlacement bool, maxWidth float32) Coord {
	size := Coord{}
	for _, layer := range b.Layers {
		if layer.ForStates(b.States) {
			size = size.Max(layer.PreferredSize(b, ctx, maxWidth))
		}
	}
	if !ignorePlacement {
		size = size.Max(Coord{
			X: b.Placement.PreferredWidth(),
			Y: b.Placement.PreferredHeight(),
		})
	}
	size = size.Max(b.MinSize)
	return size
}

func (c *Base) Update(update Update) {
	dirty := DirtyNone

	for i := range c.Layers {
		dirty.Add(c.Layers[i].Update(c, update))
	}
	for _, child := range c.Children {
		child.Update(update)
	}

	dirty.Add(c.Animation.Update(c, update))
	dirty.Add(c.CheckForChanges())

	if c.Hooks.OnUpdate != nil {
		dirty.Add(c.Hooks.OnUpdate(c, update))
	}

	c.Dirty(dirty)
}

func (c *Base) CheckForChanges() Dirty {
	dirty := DirtyNone

	if c.Transparency.Cleaned() {
		dirty.Add(DirtyVisual)
	}

	if c.lastTransform != c.Transform {
		dirty.Add(DirtyVisual)
		c.lastTransform = c.Transform
	}

	return dirty
}

func (c *Base) ComputeRenderContext() *RenderContext {
	if c.parent == nil {
		return c.ui.renderContext.WithBoundsAndTextStyles(c.Bounds, c.TextStyles)
	} else {
		return c.parent.ComputeRenderContext().WithBoundsAndTextStyles(c.Bounds, c.TextStyles)
	}
}

func (c *Base) Render(ctx *RenderContext, out *VertexBuffers) {
	if c.Transparency.Get() == 1 {
		return
	}

	c.visualBounds.Clear()

	baseCtx := ctx.WithBoundsAndTextStyles(c.Bounds, c.TextStyles)
	rendered := NewVertexIterator(out)
	renderedIndex := NewIndexIterator(out)

	if len(c.Layers) > 0 {
		for _, layer := range c.Layers {
			if layer.ForStates(c.States) {
				layer.Render(c, baseCtx, out)
			}
		}

		c.PostProcess(baseCtx, out, renderedIndex, rendered)
		c.Hooks.OnPostProcessLayers.Run(c, baseCtx, out, renderedIndex, rendered)
	}

	if len(c.Children) > 0 {
		clipBounds := c.Clip.GetBoundsIn(c.Bounds)

		out.ClipMaybe(clipBounds, func(inner *VertexBuffers) {
			renderedChildren := NewVertexIterator(inner)
			renderedIndex := NewIndexIterator(inner)

			for _, child := range c.Children {
				child.Render(baseCtx, inner)
			}

			c.PostProcess(baseCtx, inner, renderedIndex, renderedChildren)
			c.Hooks.OnPostProcessChildren.Run(c, baseCtx, out, renderedIndex, renderedChildren)
		})
	}

	c.dirty.Remove(DirtyVisual)

	if c.Hooks.OnRender != nil {
		c.Hooks.OnRender(c, baseCtx, out)
	}

	c.updateVisualBounds(rendered)
}

func (c *Base) PostProcess(ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator) {
	modifier := c.ui.Theme.StatePostProcess[c.States]
	modifier.Run(c, ctx, out, index, vertex)

	c.Animation.PostProcess(c, ctx, out, index, vertex)

	if c.Transparency.Get() > 0 {
		alphaMultiplier := 1 - c.Transparency.Get()
		for vertex.HasNext() {
			v := vertex.Next()
			v.Color.A *= alphaMultiplier
		}
		vertex.Reset()
	}

	transform := &c.Transform
	if transform.HasAffect() {
		for vertex.HasNext() {
			v := vertex.Next()
			v.X, v.Y = transform.Transform(v.X, v.Y)
		}
	}

	c.Hooks.OnPostProcess.Run(c, ctx, out, index, vertex)
}

func (c *Base) updateVisualBounds(vertex VertexIterator) {
	for vertex.HasNext() {
		v := vertex.Next()
		c.visualBounds.Include(v.X, v.Y)
	}
}

func (c *Base) Parent() Component {
	if c == nil || c.parent == nil {
		return nil
	}
	return c.parent
}

func (c *Base) IsInside(pt Coord) bool {
	if !c.Bounds.InsideCoord(pt) {
		return false
	}

	if len(c.OverShape) > 0 {
		normalized := Coord{
			X: c.Bounds.Dx(pt.X),
			Y: c.Bounds.Dy(pt.Y),
		}
		if !InPolygon(c.OverShape, normalized) {
			return false
		}
	}

	return true
}

func (c *Base) At(pt Coord) Component {
	if !c.IsInside(pt) {
		return nil
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
	if disabled == c.IsDisabled() {
		return
	}
	if disabled {
		c.AddStates(StateDisabled)
		c.PlayEvent(AnimationEventDisabled)
	} else {
		c.RemoveStates(StateDisabled)
		c.PlayEvent(AnimationEventEnabled)
	}
}

func (c *Base) OnPointer(ev *PointerEvent) {
	if c.IsDisabled() {
		return
	}

	c.Events.OnPointer.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	switch ev.Type {
	case PointerEventEnter:
		c.AddStates(StateHover)
		c.PlayEvent(AnimationEventPointerEnter)
	case PointerEventDown:
		c.AddStates(StatePressed)
		c.PlayEvent(AnimationEventPointerDown)
	case PointerEventLeave:
		c.RemoveStates(StateHover | StateDragOver)
		c.PlayEvent(AnimationEventPointerLeave)
	case PointerEventUp:
		c.RemoveStates(StatePressed | StateDragOver)
		c.PlayEvent(AnimationEventPointerUp)
	}

	c.Cursors.HandlePointer(ev, c)
}

func (c *Base) OnKey(ev *KeyEvent) {
	if c.IsDisabled() {
		return
	}

	c.Events.OnKey.Trigger(ev)
}

func (c *Base) OnFocus(ev *Event) {
	if c.IsDisabled() {
		return
	}

	c.Events.OnFocus.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	c.AddStates(StateFocused)
	c.PlayEvent(AnimationEventFocus)
}

func (c *Base) OnBlur(ev *Event) {
	if c.IsDisabled() {
		return
	}

	c.Events.OnBlur.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	c.RemoveStates(StateFocused)
	c.PlayEvent(AnimationEventBlur)
}

func (c *Base) OnDrag(ev *DragEvent) {
	if c.IsDisabled() {
		return
	}

	c.Events.OnDrag.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	switch ev.Type {
	case DragEventStart:
		c.AddStates(StateDragging)
		c.PlayEvent(AnimationEventDragStart)
	case DragEventEnd:
		c.RemoveStates(StateDragging)
		c.PlayEvent(AnimationEventDragStop)
	case DragEventOver:
		if !c.States.Is(StateDragOver) {
			c.PlayEvent(AnimationEventDragEnter)
		}
		c.AddStates(StateDragOver)
	case DragEventCancel:
		c.PlayEvent(AnimationEventDragCancel)
	case DragEventDrop:
		c.PlayEvent(AnimationEventDrop)
	}

	c.Cursors.HandleDrag(ev, c)
}

type Template struct {
	PreLayers       []Layer
	PostLayers      []Layer
	Focusable       bool
	Draggable       bool
	Droppable       bool
	PreEvents       Events
	PostEvents      Events
	Clip            Placement
	TextStyles      *TextStylesOverride
	Shape           Shape
	OverShape       []Coord
	Animations      *Animations
	AnimationsMerge bool
	Cursors         Cursors
	CursorsMerge    bool
	PreHooks        Hooks
	PostHooks       Hooks
}

func (b *Base) ApplyTemplate(t *Template) {
	if t == nil {
		return
	}
	if len(t.PreLayers) > 0 {
		layers := b.Layers[:]
		b.Layers = append(make([]Layer, 0, len(b.Layers)+len(t.PreLayers)+len(t.PostLayers)))
		b.Layers = append(b.Layers, t.PreLayers...)
		b.Layers = append(b.Layers, layers...)
	}
	b.Layers = append(b.Layers, t.PostLayers...)
	b.Focusable = t.Focusable || b.Focusable
	b.Draggable = t.Draggable || b.Draggable
	b.Droppable = t.Droppable || b.Droppable
	b.Events.Add(t.PreEvents, true)
	b.Events.Add(t.PostEvents, false)
	if !b.Clip.Defined() {
		b.Clip = t.Clip
	}
	if b.TextStyles == nil {
		b.TextStyles = t.TextStyles
	}
	if b.Shape == nil {
		b.Shape = t.Shape
	}
	if b.OverShape == nil {
		b.OverShape = t.OverShape
	}
	if t.Animations != nil && (t.AnimationsMerge || b.Animations == nil) {
		if b.Animations == nil {
			b.Animations = &Animations{}
		}
		b.Animations.Merge(t.Animations, false)
	}
	if !t.Cursors.Empty() && (t.CursorsMerge || b.Cursors.Empty()) {
		b.Cursors.Merge(t.Cursors.EnumMap, false, cursorNil)
	}
	b.Hooks.Add(t.PreHooks, true)
	b.Hooks.Add(t.PostHooks, false)
}
