package ui

import (
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/util"
)

type Base struct {
	Name              id.Identifier
	Layers            []Layer
	Placement         Placement
	Bounds            Bounds
	Children          []*Base
	ChildrenOrderless bool
	Focusable         bool
	Draggable         bool
	Droppable         bool
	Events            Events
	States            State
	Clip              Placement
	Transform         Transform
	TextStyles        *TextStylesOverride
	Shape             Shape
	OverShape         []Coord
	Transparency      Watch[float32]
	Animation         AnimationState
	Animations        *Animations
	Cursors           Cursors
	Hooks             Hooks
	Colors            Colors
	Margin            AmountBounds
	MarginBounds      Bounds
	MinSize           Coord
	MaxSize           Coord
	Layout            Layout

	visualBounds  Bounds
	dirty         Dirty
	parent        *Base
	ui            *UI
	lastTransform Transform
	shown         []*Base
	shownDirty    bool
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

func (b *Base) Edit(editor func(*Base)) *Base {
	editor(b)
	return b
}

func (b *Base) RemoveStates(state State) {
	b.SetState(b.States.WithRemove(state))
}

func (b *Base) AddStates(state State) {
	b.SetState(b.States.WithAdd(state))
}

func (b *Base) SetPlacement(placement Placement) {
	if b.Placement != placement {
		b.Placement = placement
		b.Dirty(DirtyPlacement)
	}
}

func (b *Base) GetPlacement() Placement {
	return b.Placement
}

func (b *Base) Init() {
	if b.ui != nil && !b.Name.Empty() {
		b.ui.Named.Set(b.Name, b)
	}
	b.shownDirty = true
	b.Placement.Init(Maximized())
	if b.States == 0 {
		b.States = StateDefault
	}
	if b.Shape != nil {
		b.Shape.Init()
	}
	for i := range b.Layers {
		b.Layers[i].Init(b)
	}
	if b.Layout != nil {
		b.Layout.Init(b)
	}
	for _, child := range b.Children {
		child.parent = b
		child.ui = b.ui
		child.Init()
	}
	b.Dirty(DirtyPlacement)
	b.PlayEvent(AnimationEventShow)

	if b.Hooks.OnInit != nil {
		b.Hooks.OnInit(b)
	}
}

func (b *Base) Place(ctx *RenderContext, parent Bounds, force bool) {
	baseCtx := ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)

	doPlacement := force || b.dirty.Is(DirtyPlacement)
	if doPlacement {
		newBounds := b.Placement.GetBoundsIn(parent)
		if newBounds != b.Bounds {
			b.Bounds = newBounds

			margin := b.Margin.GetBounds(baseCtx.AmountContext)
			b.MarginBounds = b.Bounds.Expand(margin)

			for i := range b.Layers {
				b.Layers[i].Place(b, b.Bounds)
			}
			b.Dirty(DirtyVisual)
			force = true
		}
		b.dirty.Remove(DirtyPlacement)
	}

	shown := b.ShownChildren()

	if b.Layout != nil {
		b.Layout.Layout(b, baseCtx, b.Bounds, shown)
	}

	for _, child := range shown {
		child.Place(baseCtx, b.Bounds, force)
	}

	if b.Hooks.OnPlace != nil {
		b.Hooks.OnPlace(b, parent, baseCtx, force)
	}
}

func (b *Base) ShownChildren() []*Base {
	if b.shownDirty {
		b.shown = util.SliceEnsureSize(b.shown, len(b.Children))
		shownCount := 0
		for _, child := range b.Children {
			if !child.IsHidden() {
				b.shown[shownCount] = child
				shownCount++
			}
		}
		b.shown = util.SliceResize(b.shown, shownCount)
		b.shownDirty = false
	}

	return b.shown
}

// PreferredSize computes the preferred size possibly given a width we are fitting this component into.
// If maxWidth = 0 then the preferred size return will present one with a minimum possible width.
// ctx should be relative to this component
// maxWidth is the width we're aiming for for this component
func (b *Base) PreferredSize(ctx *RenderContext, maxWidth float32) Coord {
	baseCtx := ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)

	size := Coord{}
	if b.MaxSize.X > 0 && b.MaxSize.X < maxWidth {
		maxWidth = b.MaxSize.X
	}
	for _, layer := range b.Layers {
		if layer.ForStates(b.States) {
			size = size.Max(layer.PreferredSize(b, baseCtx, maxWidth))
		}
	}
	size = size.Max(b.MinSize)
	maxWidth = max(maxWidth, size.X)
	shown := b.ShownChildren()
	if len(shown) > 0 {
		layout := b.Layout
		if layout == nil {
			layout = LayoutStatic{}
		}
		size = size.Max(layout.PreferredSize(b, baseCtx, maxWidth, shown))
	}
	if b.MaxSize.X > 0 {
		size.X = min(b.MaxSize.X, size.X)
	}
	if b.MaxSize.Y > 0 {
		size.Y = min(b.MaxSize.Y, size.Y)
	}
	return size
}

func (b *Base) Compact() {
	b.CompactFor(b.ComputeRenderContext())
}

func (b *Base) CompactFor(ctx *RenderContext) {
	size := b.PreferredSize(ctx, 0)
	placement := b.Placement
	placement.Right.Base = placement.Left.Base + size.X
	placement.Bottom.Base = placement.Top.Base + size.Y
	b.SetPlacement(placement)
}

func (b *Base) Tighten() {
	b.TightenFor(b.ComputeRenderContext())
}

func (b *Base) TightenFor(ctx *RenderContext) {
	size := b.PreferredSize(ctx, b.Bounds.Width())
	placement := b.Placement
	placement.Right.Base = placement.Left.Base + size.X
	placement.Bottom.Base = placement.Top.Base + size.Y
	b.SetPlacement(placement)
}

func (b *Base) Update(update Update) {
	dirty := DirtyNone

	if b.IsHidden() && !b.ui.UpdateHidden {
		return
	}

	for i := range b.Layers {
		dirty.Add(b.Layers[i].Update(b, update))
	}
	for childIndex := 0; childIndex < len(b.Children); childIndex++ {
		child := b.Children[childIndex]
		child.Update(update)
		if child.parent != b {
			childIndex--
		}
	}

	dirty.Add(b.Animation.Update(b, update))
	dirty.Add(b.CheckForChanges())

	if b.Hooks.OnUpdate != nil {
		dirty.Add(b.Hooks.OnUpdate(b, update))
	}

	b.Dirty(dirty)

	if !b.Animation.IsAnimating() {
		if b.States.Is(StateRemoving) {
			b.removeNow()
		}
		if b.States.Is(StateHiding) {
			b.hideNow()
		}
		if b.States.Is(StateShowing) {
			b.showNow()
		}
	}
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

func (b *Base) WillRender() bool {
	if b.Transparency.Get() == 1 && !b.Animation.IsAnimating() {
		return false
	}
	if b.IsHidden() {
		return false
	}
	return true
}

func (b *Base) Render(ctx *RenderContext, out *VertexBuffers) {
	if !b.WillRender() {
		return
	}

	b.visualBounds.Clear()

	baseCtx := ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)
	rendered := NewVertexIterator(out)
	renderedIndex := NewIndexIterator(out)

	if len(b.Layers) > 0 {
		for _, layer := range b.Layers {
			if layer.ForStates(b.States) {
				layer.Render(b, baseCtx, out)
			}
		}

		b.PostProcess(baseCtx, out, renderedIndex, rendered)
		b.Hooks.OnPostProcessLayers.Run(b, baseCtx, out, renderedIndex, rendered)
	}

	if len(b.Children) > 0 {
		clipBounds := b.Clip.GetBoundsIn(b.Bounds)

		out.ClipMaybe(clipBounds, func(inner *VertexBuffers) {
			renderedChildren := NewVertexIterator(inner)
			renderedIndex := NewIndexIterator(inner)

			shown := b.ShownChildren()
			for _, child := range shown {
				child.Render(baseCtx, inner)
			}

			b.PostProcess(baseCtx, inner, renderedIndex, renderedChildren)
			b.Hooks.OnPostProcessChildren.Run(b, baseCtx, out, renderedIndex, renderedChildren)
		})
	}

	b.dirty.Remove(DirtyVisual)

	if b.Hooks.OnRender != nil {
		b.Hooks.OnRender(b, baseCtx, out)
	}

	b.updateVisualBounds(rendered)
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

func (b *Base) At(pt Coord) Component {
	transparency := b.Transparency.Get()
	if transparency > 0 && b.ui.TransparencyThreshold > 0 && transparency >= b.ui.TransparencyThreshold {
		return nil
	}
	if b.ui.TransformPointer {
		if b.Transform.HasAffect() {
			inv := b.Transform.GetInvert()
			pt = inv.TransformCoord(pt)
		}
	}

	if !b.IsInside(pt) {
		return nil
	}

	shown := b.ShownChildren()
	last := len(shown) - 1
	for i := last; i >= 0; i-- {
		at := shown[i].At(pt)
		if at != nil {
			return at
		}
	}

	return b
}

func (b *Base) SetVisible(show bool) {
	if show {
		b.Show()
	} else {
		b.Hide()
	}
}

func (b *Base) Hide() {
	if !b.States.Is(StateHidden | StateHiding) {
		if b.PlayEvent(AnimationEventHide) {
			b.SetState(b.States.WithAdd(StateHiding).WithRemove(StateShowing))
		} else {
			b.hideNow()
		}
	}
}

func (b *Base) hideNow() {
	b.SetState(b.States.WithAdd(StateHidden).WithRemove(StateHiding | StateShowing))
	b.parent.ChildrenUpdated()
}

func (b *Base) Show() {
	if !b.States.Is(StateShowing) && b.States.Is(StateHidden|StateHiding) {
		if b.PlayEvent(AnimationEventShow) {
			b.SetState(b.States.WithAdd(StateShowing).WithRemove(StateHidden | StateHiding))
			b.parent.ChildrenUpdated()
		} else {
			b.showNow()
		}
	}
}

func (b *Base) showNow() {
	b.RemoveStates(StateHidden | StateHiding | StateShowing)
	b.parent.ChildrenUpdated()
}

func (b *Base) Remove() {
	if b.parent != nil && !b.States.Is(StateRemoving) {
		if b.PlayEvent(AnimationEventRemove) {
			b.SetState(StateRemoving)
		} else {
			b.removeNow()
		}
	}
}

func (b *Base) removeNow() {
	index := b.Order()
	if index != -1 {
		if b.parent.ChildrenOrderless {
			b.parent.Children = util.SliceRemoveAtReplace(b.parent.Children, index)
		} else {
			b.parent.Children = util.SliceRemoveAt(b.parent.Children, index)
		}
		b.parent.ChildrenUpdated()
		b.parent = nil
	}
}

func (b *Base) Order() int {
	if b.parent == nil {
		return -1
	}
	return util.SliceIndexOf(b.parent.Children, b)
}

func (b *Base) SetOrder(order int) {
	currentOrder := b.Order()
	if currentOrder == -1 || order == currentOrder {
		return
	}
	util.SliceMove(b.parent.Children, currentOrder, order)
	b.parent.ChildrenUpdated()
}

func (b *Base) ChildrenUpdated() {
	if b != nil {
		b.Dirty(DirtyPlacement)
		b.shownDirty = true
	}
}

func (b *Base) AddChildren(children ...*Base) {
	for _, child := range children {
		child.parent = b
		child.ui = b.ui
		child.Init()
	}
	b.Children = append(b.Children, children...)
	b.ChildrenUpdated()
}

func (b *Base) BringToFront() {
	if b.parent != nil {
		b.SetOrder(len(b.parent.Children) - 1)
	}
}

func (b *Base) SendToBack() {
	b.SetOrder(0)
}

func (b Base) IsFocusable() bool {
	return b.Focusable && !b.IsDisabled()
}

func (b Base) IsDraggable() bool {
	return b.Draggable && !b.IsDisabled()
}

func (b Base) IsDroppable() bool {
	return b.Droppable && !b.IsDisabled()
}

func (b Base) IsDisabled() bool {
	return b.States.Is(StatesDisabled)
}

func (b Base) IsShown() bool {
	return b.States.Not(StateHidden)
}

func (b Base) IsHidden() bool {
	return b.States.Is(StateHidden)
}

func (b *Base) SetDisabled(disabled bool) {
	if disabled == b.IsDisabled() {
		return
	}
	if disabled {
		b.AddStates(StateDisabled)
		b.PlayEvent(AnimationEventDisabled)
	} else {
		b.RemoveStates(StateDisabled)
		b.PlayEvent(AnimationEventEnabled)
	}
}

func (b *Base) OnPointer(ev *PointerEvent) {
	if b.IsDisabled() {
		return
	}

	b.Events.OnPointer.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	switch ev.Type {
	case PointerEventEnter:
		b.AddStates(StateHover)
		b.PlayEvent(AnimationEventPointerEnter)
	case PointerEventDown:
		b.AddStates(StatePressed)
		b.PlayEvent(AnimationEventPointerDown)
	case PointerEventLeave:
		b.RemoveStates(StateHover | StateDragOver)
		b.PlayEvent(AnimationEventPointerLeave)
	case PointerEventUp:
		b.RemoveStates(StatePressed | StateDragOver)
		b.PlayEvent(AnimationEventPointerUp)
	}

	b.Cursors.HandlePointer(ev, b)
}

func (b *Base) OnKey(ev *KeyEvent) {
	if b.IsDisabled() {
		return
	}

	b.Events.OnKey.Trigger(ev)
}

func (b *Base) OnFocus(ev *Event) {
	if b.IsDisabled() || b.States.Is(StateFocused) {
		return
	}

	b.Events.OnFocus.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	b.AddStates(StateFocused)
	b.PlayEvent(AnimationEventFocus)
}

func (b *Base) OnBlur(ev *Event) {
	if b.IsDisabled() || b.States.Not(StateFocused) {
		return
	}

	b.Events.OnBlur.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	b.RemoveStates(StateFocused)
	b.PlayEvent(AnimationEventBlur)
}

func (b *Base) IsDragging() bool {
	return b.ui.Dragging == b
}

func (b *Base) OnDrag(ev *DragEvent) {
	if b.IsDisabled() {
		return
	}

	b.Events.OnDrag.Trigger(ev)

	if ev.Cancel || ev.Capture {
		return
	}

	switch ev.Type {
	case DragEventStart:
		b.AddStates(StateDragging)
		b.PlayEvent(AnimationEventDragStart)
	case DragEventEnd:
		b.RemoveStates(StateDragging)
		b.PlayEvent(AnimationEventDragStop)
	case DragEventOver:
		if !b.States.Is(StateDragOver) {
			b.PlayEvent(AnimationEventDragEnter)
		}
		b.AddStates(StateDragOver)
	case DragEventCancel:
		b.PlayEvent(AnimationEventDragCancel)
	case DragEventDrop:
		b.PlayEvent(AnimationEventDrop)
	}

	b.Cursors.HandleDrag(ev, b)
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
	Colors          Colors
	ColorsMerge     bool
	PreHooks        Hooks
	PostHooks       Hooks
	Margin          AmountBounds
	MinSize         Coord
	MaxSize         Coord
	Layout          Layout
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
	if !t.Colors.Empty() && (t.ColorsMerge || b.Colors.Empty()) {
		b.Colors.Merge(t.Colors.EnumMap, false, colorableNil)
	}
	b.Hooks.Add(t.PreHooks, true)
	b.Hooks.Add(t.PostHooks, false)
	if b.Margin.IsZero() {
		b.Margin = t.Margin
	}
	if b.MinSize.IsZero() {
		b.MinSize = t.MinSize
	}
	if b.MaxSize.IsZero() {
		b.MaxSize = t.MaxSize
	}
	if b.Layout == nil {
		b.Layout = t.Layout
	}
}
