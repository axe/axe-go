package ui

import (
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/util"
)

type Base struct {
	Name                        id.Identifier
	Layers                      []Layer
	Placement                   Placement
	RelativePlacement           Placement
	Bounds                      Bounds
	Children                    []*Base
	ChildrenOrderless           bool
	Focusable                   bool
	Draggable                   bool
	Droppable                   bool
	Events                      Events
	States                      State
	Clip                        Placement
	Transform                   Transform
	TextStyles                  *TextStylesOverride
	Shape                       Shape
	OverShape                   []Coord
	Transparency                float32
	Color                       ColorModify
	Animation                   AnimationState
	Animations                  *Animations
	Cursors                     Cursors
	Hooks                       Hooks
	Colors                      Colors
	Margin                      AmountBounds
	MinSize                     AmountPoint
	MaxSize                     AmountPoint
	Layout                      Layout
	IgnoreLayoutPreferredWidth  bool
	IgnoreLayoutPreferredHeight bool

	visualBounds        Bounds
	dirty               Dirty
	parent              *Base
	renderParent        *Base
	ui                  *UI
	shown               []*Base
	shownDirty          bool
	layerBuffers        *VertexBuffers
	childrenBufferQueue *VertexQueue
}

func (b *Base) UI() *UI {
	return b.ui
}

func (b *Base) GetDirty() Dirty {
	return b.dirty
}

func (b *Base) Dirty(dirty Dirty) {
	if b.dirty.NotAll(dirty) {
		b.dirty.Add(dirty)
		parentDirty := dirty.ParentDirty()
		if b.parent != nil {
			b.parent.Dirty(parentDirty)
		}
		if b.renderParent != nil && b.renderParent != b.parent {
			b.renderParent.Dirty(parentDirty)
		}
	}
}

func (b *Base) SetState(state State) {
	state.Remove(StateDefault)
	if state.None() {
		state.Add(StateDefault)
	}

	if !b.dirty.Is(DirtyVisual) {
		for _, layer := range b.Layers {
			if layer.ForStates(state) != layer.ForStates(b.States) {
				b.Dirty(DirtyVisual)
				break
			}
		}
	}
	if b.ui.Theme.StatePostProcess[state] != nil {
		b.Dirty(DirtyPostProcess)
	}

	b.States = state
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

func (b *Base) SetRelativePlacement(placement Placement) {
	if b.RelativePlacement != placement {
		b.RelativePlacement = placement
		b.Dirty(DirtyPlacement)
	}
}

func (b *Base) SetTransform(transform Transform) {
	if transform.IsEffectivelyIdentity() {
		transform.Identity()
	}
	if b.Transform != transform {
		b.Transform = transform
		b.Dirty(DirtyPostProcess)
	}
}

func (b *Base) SetTransparency(transparency float32) {
	if b.Transparency != transparency {
		b.Transparency = transparency
		b.Dirty(DirtyPostProcess)
	}
}

func (b *Base) SetColor(color ColorModify) {
	if !b.Color.Equals(color) {
		b.Color = color.GetEffective()
		b.Dirty(DirtyPostProcess)
	}
}

func (b *Base) Init() {
	if b.ui != nil && !b.Name.Empty() {
		b.ui.Named.Set(b.Name, b)
	}
	b.shownDirty = true
	b.Placement.Init(Maximized())
	b.RelativePlacement.Init(Maximized())
	b.States &= (StateDisabled | StateFocused)
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
		child.renderParent = b
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
	if !force && !b.dirty.Is(DirtyChildPlacement|DirtyPlacement) {
		return
	}

	doPlacement := force || b.dirty.Is(DirtyPlacement)
	if doPlacement {
		relativeBounds := b.RelativePlacement.GetBoundsIn(parent)
		newBounds := b.Placement.GetBoundsIn(relativeBounds)
		force = b.SetBounds(newBounds)

		b.dirty.Remove(DirtyPlacement)
	}

	baseCtx := ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)

	doChildPlacement := force || b.dirty.Is(DirtyChildPlacement)
	if doChildPlacement {
		shown := b.ShownChildren()

		if b.Layout != nil {
			b.Layout.Layout(b, baseCtx, b.Bounds, shown)
		}

		for _, child := range b.Children {
			if !child.IsHidden() {
				child.Place(baseCtx, child.parent.Bounds, force)
			}
		}

		b.dirty.Remove(DirtyChildPlacement)
	}

	if b.Hooks.OnPlace != nil {
		b.Hooks.OnPlace(b, parent, baseCtx, force)
	}
}

func (b *Base) SetBounds(newBounds Bounds) bool {
	if newBounds != b.Bounds {
		b.Bounds = newBounds
		for i := range b.Layers {
			b.Layers[i].Place(b, b.Bounds)
		}
		dirty := DirtyVisual
		if !b.dirty.Is(DirtyPlacement) {
			dirty.Add(DirtyPlacement)
		}
		b.Dirty(dirty)
		return true
	}
	return false
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
		if child.parent == b {
			child.Update(update)
			if child.parent != b {
				childIndex--
			}
		}
	}

	dirty.Add(b.Animation.Update(b, update))

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

func (b *Base) Relayout() {
	b.Dirty(DirtyChildPlacement)
}

func (b *Base) Rerender() {
	b.Dirty(DirtyVisual | DirtyChildVisual)
}

func (b *Base) CouldRender() bool {
	if b.Transparency == 1 && !b.Animation.IsAnimating() {
		return false
	}
	if b.IsHidden() {
		return false
	}
	return true
}

func (b *Base) HasPostProcess() bool {
	return b.dirty.Is(DirtyPostProcess) ||
		b.ui.Theme.StatePostProcess[b.States] != nil ||
		b.Hooks.OnPostProcess != nil ||
		b.Hooks.OnPostProcessLayers != nil ||
		b.Hooks.OnPostProcessChildren != nil ||
		b.Animation.IsAnimating() ||
		b.Transparency > 0 ||
		b.Color != nil ||
		b.Transform.HasAffect()
}

func (b *Base) Render(ctx *RenderContext, queue *VertexQueue) {
	if !b.CouldRender() {
		return
	}

	hasPostProcess := b.HasPostProcess()
	useCache := b.dirty.Not(DirtyVisual | DirtyChildVisual)

	if !hasPostProcess && useCache {
		if b.layerBuffers != nil {
			queue.Add(b.layerBuffers)
		}
		if b.childrenBufferQueue != nil {
			queue.Add(b.childrenBufferQueue)
		}

		return
	}

	if b.parent != b.renderParent {
		ctx = b.ComputeRenderContext()
	} else {
		ctx = ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)
	}

	queueSizeBeforeLayers := queue.Len()

	baseVertices := NewVertexIterator(queue)
	baseIndices := NewIndexIterator(queue)

	layerVertices := NewVertexIterator(queue)
	layerIndices := NewIndexIterator(queue)

	if b.dirty.Is(DirtyVisual) {
		if b.layerBuffers == nil {
			b.layerBuffers = BufferPool.Get()
		}

		b.layerBuffers.Clear()

		if len(b.Layers) > 0 {
			for _, layer := range b.Layers {
				if layer.ForStates(b.States) {
					layer.Render(b, ctx, b.layerBuffers)
				}
			}
		}

		b.dirty.Remove(DirtyVisual)
	}

	if b.layerBuffers != nil {
		queue.Add(b.layerBuffers)

		if hasPostProcess {
			queue.Clone(queueSizeBeforeLayers, queue.Len())
		}
	}

	layerVertices.Limit()
	layerIndices.Limit()

	queueSizeBeforeChildren := queue.Len()

	childrenVertices := NewVertexIterator(queue)
	childrenIndices := NewIndexIterator(queue)
	clipped := false

	if b.dirty.Is(DirtyChildVisual) {
		if b.childrenBufferQueue == nil {
			b.childrenBufferQueue = BufferQueuePool.Get()
		}

		b.childrenBufferQueue.Clear()

		if len(b.Children) > 0 {
			clipBounds := b.Clip.GetBoundsIn(b.Bounds)

			clipped = b.childrenBufferQueue.Clip(clipBounds, func(innerQueue *VertexQueue, clipping bool) {
				shown := b.ShownChildren()
				for _, child := range shown {
					beforeChild := innerQueue.Position()

					child.Render(ctx, innerQueue)

					if clipping && !child.visualBounds.Intersects(clipBounds) {
						innerQueue.Reset(beforeChild)
					}
				}
			})
		}

		b.dirty.Remove(DirtyChildVisual)
	}

	if b.childrenBufferQueue != nil {
		queue.Add(b.childrenBufferQueue)

		if hasPostProcess && !clipped {
			queue.Clone(queueSizeBeforeChildren, queue.Len())
		}
	}

	if hasPostProcess {
		childrenVertices.Limit()
		childrenIndices.Limit()

		baseVertices.Limit()
		baseIndices.Limit()

		b.Hooks.OnPostProcessLayers.Run(b, ctx, queue, layerIndices, layerVertices)
		b.Hooks.OnPostProcessChildren.Run(b, ctx, queue, childrenIndices, childrenVertices)

		b.PostProcess(ctx, queue, baseIndices, baseVertices)
	}

	if b.Hooks.OnRender != nil {
		b.Hooks.OnRender(b, ctx, queue)
	}

	b.updateVisualBounds(baseVertices)
}

func (b *Base) PostProcess(ctx *RenderContext, out VertexIterable, index IndexIterator, vertex VertexIterator) {

	modifier := b.ui.Theme.StatePostProcess[b.States]
	modifier.Run(b, ctx, out, index, vertex)

	b.Animation.PostProcess(b, ctx, out, index, vertex)

	if alphaMultiplier := 1 - b.Transparency; alphaMultiplier < 0.9999 {
		for vertex.HasNext() {
			v := vertex.Next()
			if !v.HasColor {
				v.Color = ColorWhite
				v.HasColor = true
			}
			v.Color.A *= alphaMultiplier
		}
		vertex.Reset()
	}

	if colorModifier := b.Color; colorModifier != nil {
		for vertex.HasNext() {
			v := vertex.Next()
			if !v.HasColor {
				v.Color = ColorWhite
				v.HasColor = true
			}
			v.Color = colorModifier(v.Color)
		}
		vertex.Reset()
	}

	if transform := &b.Transform; transform.HasAffect() {
		for vertex.HasNext() {
			v := vertex.Next()
			v.X, v.Y = transform.Transform(v.X, v.Y)
		}
	}

	b.Hooks.OnPostProcess.Run(b, ctx, out, index, vertex)

	b.dirty.Remove(DirtyPostProcess)
}

func (b *Base) ComputeRenderContext() *RenderContext {
	if b.parent == nil {
		return b.ui.renderContext.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)
	} else {
		return b.parent.ComputeRenderContext().WithBoundsAndTextStyles(b.Bounds, b.TextStyles)
	}
}

func (b *Base) IsRenderParent(child *Base) bool {
	return child.renderParent == b
}

func (b *Base) ShownChildren() []*Base {
	if b.shownDirty {
		b.shown = util.SliceEnsureSize(b.shown, len(b.Children))
		shownCount := 0
		for _, child := range b.Children {
			if !child.IsHidden() && b.IsRenderParent(child) {
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
	minSizeX, minSizeY := b.MinSize.Get(ctx.AmountContext)
	maxSizeX, maxSizeY := b.MaxSize.Get(ctx.AmountContext)

	if maxSizeX > 0 && maxSizeX < maxWidth {
		maxWidth = maxSizeX
	}
	for _, layer := range b.Layers {
		if layer.ForStates(b.States) {
			size = size.Max(layer.PreferredSize(b, baseCtx, maxWidth))
		}
	}
	size = size.Max(Coord{X: minSizeX, Y: minSizeY})
	maxWidth = util.Max(maxWidth, size.X)
	shown := b.ShownChildren()
	if len(shown) > 0 && (!b.IgnoreLayoutPreferredHeight || !b.IgnoreLayoutPreferredWidth) {
		layout := b.Layout
		if layout == nil {
			layout = LayoutStatic{}
		}
		layoutSize := layout.PreferredSize(b, baseCtx, maxWidth, shown)
		if b.IgnoreLayoutPreferredHeight {
			layoutSize.Y = 0
		}
		if b.IgnoreLayoutPreferredWidth {
			layoutSize.X = 0
		}
		size = size.Max(layoutSize)
	}
	if maxSizeX > 0 {
		size.X = util.Min(maxSizeX, size.X)
	}
	if maxSizeY > 0 {
		size.Y = util.Min(maxSizeY, size.Y)
	}
	return size
}

func (b *Base) Compact() {
	b.CompactFor(b.ComputeRenderContext())
}

func (b *Base) CompactFor(ctx *RenderContext) {
	size := b.PreferredSize(ctx, 0)
	b.SetPlacement(b.Placement.WithSize(size.X, size.Y))
}

func (b *Base) Tighten() {
	b.TightenFor(b.ComputeRenderContext())
}

func (b *Base) TightenFor(ctx *RenderContext) {
	size := b.PreferredSize(ctx, b.Bounds.Width())
	b.SetPlacement(b.Placement.WithSize(size.X, size.Y))
}

func (b *Base) updateVisualBounds(baseVertices VertexIterator) {
	b.visualBounds.Clear()

	for baseVertices.HasNext() {
		v := baseVertices.Next()
		b.visualBounds.Include(v.X, v.Y)
	}
}

func (b *Base) Parent() *Base {
	if b == nil {
		return nil
	}
	return b.parent
}

func (b *Base) IsInside(pt Coord) bool {
	if !b.Bounds.InsideCoord(pt) {
		return false
	}

	if len(b.OverShape) > 0 {
		normalized := Coord{
			X: b.Bounds.Dx(pt.X),
			Y: b.Bounds.Dy(pt.Y),
		}
		if !InPolygon(b.OverShape, normalized) {
			return false
		}
	}

	return true
}

func (b *Base) At(pt Coord) *Base {
	transparency := b.Transparency
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
	b.renderParent.ChildrenUpdated()
}

func (b *Base) Show() {
	if !b.States.Is(StateShowing) && b.States.Is(StateHidden|StateHiding) {
		if b.PlayEvent(AnimationEventShow) {
			b.SetState(b.States.WithAdd(StateShowing).WithRemove(StateHidden | StateHiding))
			b.renderParent.ChildrenUpdated()
		} else {
			b.showNow()
		}
	}
}

func (b *Base) showNow() {
	b.RemoveStates(StateHidden | StateHiding | StateShowing)
	b.renderParent.ChildrenUpdated()
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
	if b.renderParent != nil {
		b.renderParent.ChildrenUpdated()
	}
	b.cleanup()
	b.removeFrom(b.parent)
	b.parent = nil
	b.renderParent = nil
}

func (b *Base) cleanup() {
	if b.parent != b.renderParent {
		b.removeFrom(b.renderParent)
	}
	if b.layerBuffers != nil {
		BufferPool.Free(b.layerBuffers)
		b.layerBuffers = nil
	}
	if b.childrenBufferQueue != nil {
		BufferQueuePool.Free(b.childrenBufferQueue)
		b.childrenBufferQueue = nil
	}
	for _, child := range b.Children {
		if child.parent == b {
			child.cleanup()
		}
	}
}

func (b *Base) removeFrom(parent *Base) {
	if b == nil || parent == nil {
		return
	}
	index := util.SliceIndexOf(parent.Children, b)
	if index != -1 {
		if parent.ChildrenOrderless {
			parent.Children = util.SliceRemoveAtReplace(parent.Children, index)
		} else {
			parent.Children = util.SliceRemoveAt(parent.Children, index)
		}
		parent.ChildrenUpdated()
	}
}

func (b *Base) ChildrenUpdated() {
	if b != nil {
		b.Dirty(DirtyChildPlacement | DirtyChildVisual)
		b.shownDirty = true
	}
}

func (b *Base) AddChildren(children ...*Base) {
	for _, child := range children {
		child.parent = b
		child.renderParent = b
		child.ui = b.ui
		child.Init()
	}
	b.Children = append(b.Children, children...)
	b.ChildrenUpdated()
}

func (b *Base) SetRenderParent(parent *Base) {
	if b.renderParent != b.parent {
		b.removeFrom(b.renderParent)
	}
	if parent == nil || b.parent == parent {
		b.renderParent = b.parent
	} else {
		parent.Children = append(parent.Children, b)
		b.renderParent = parent
	}
	b.renderParent.ChildrenUpdated()
}

func (b *Base) Order() int {
	if b.renderParent == nil {
		return -1
	}
	return util.SliceIndexOf(b.renderParent.Children, b)
}

func (b *Base) SetOrder(order int) {
	currentOrder := b.Order()
	if currentOrder == -1 || order == currentOrder {
		return
	}
	util.SliceMove(b.renderParent.Children, currentOrder, order)
	b.renderParent.ChildrenUpdated()
}

func (b *Base) BringToFront() {
	if b.renderParent != nil {
		b.SetOrder(len(b.renderParent.Children) - 1)
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
	PreLayers                   []Layer
	PostLayers                  []Layer
	Focusable                   bool
	Draggable                   bool
	Droppable                   bool
	PreEvents                   Events
	PostEvents                  Events
	RelativePlacement           Placement
	Clip                        Placement
	TextStyles                  *TextStylesOverride
	Shape                       Shape
	OverShape                   []Coord
	Transparency                float32
	Color                       ColorModify
	Animations                  *Animations
	AnimationsMerge             bool
	Cursors                     Cursors
	CursorsMerge                bool
	Colors                      Colors
	ColorsMerge                 bool
	PreHooks                    Hooks
	PostHooks                   Hooks
	Margin                      AmountBounds
	MinSize                     AmountPoint
	MaxSize                     AmountPoint
	Layout                      Layout
	IgnoreLayoutPreferredWidth  bool
	IgnoreLayoutPreferredHeight bool
}

func (b *Base) ApplyTemplate(t *Template) {
	if t == nil {
		return
	}
	if len(t.PreLayers) > 0 {
		layers := b.Layers
		b.Layers = make([]Layer, 0, len(b.Layers)+len(t.PreLayers)+len(t.PostLayers))
		b.Layers = append(b.Layers, t.PreLayers...)
		b.Layers = append(b.Layers, layers...)
	}
	b.Layers = append(b.Layers, t.PostLayers...)
	b.Focusable = t.Focusable || b.Focusable
	b.Draggable = t.Draggable || b.Draggable
	b.Droppable = t.Droppable || b.Droppable
	b.Events.Add(t.PreEvents, true)
	b.Events.Add(t.PostEvents, false)
	if !b.RelativePlacement.Defined() {
		b.RelativePlacement = t.RelativePlacement
	}
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
	if b.Color == nil {
		b.Color = t.Color
	}
	if b.Transparency == 0 {
		b.Transparency = t.Transparency
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
	if !b.IgnoreLayoutPreferredHeight {
		b.IgnoreLayoutPreferredHeight = t.IgnoreLayoutPreferredHeight
	}
	if !b.IgnoreLayoutPreferredWidth {
		b.IgnoreLayoutPreferredWidth = t.IgnoreLayoutPreferredWidth
	}
}
