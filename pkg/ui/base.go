package ui

import (
	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/util"
)

// The Base UI component.
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
	OverShape                   []gfx.Coord
	Transparency                float32
	Color                       color.Modify
	Animation                   AnimationState
	Animations                  *Animations
	Cursors                     Cursors
	Hooks                       Hooks
	Colors                      color.Colors
	Margin                      AmountBounds
	MinSize                     AmountPoint
	MaxSize                     AmountPoint
	Layout                      Layout
	LayoutData                  LayoutData
	IgnoreLayoutPreferredWidth  bool
	IgnoreLayoutPreferredHeight bool

	visualBounds Bounds
	dirty        Dirty
	parent       *Base
	renderParent *Base
	ui           *UI
	// A cache list of the visible children.
	shown []*Base
	// If children changed in some way where the ShownChildren method could return different results.
	// This occurs at initialization, on reordering, showing, hiding, and removal.
	shownDirty bool
	// Any layer vertex data, unmodified by post processing.
	// When layers are dirty this is updated which may trigger post processing even if not marked dirty.
	// If there is no active post processing the buffers are added to the cachedBuffers.
	layerBuffers *VertexBuffers
	// If there is post processing this is non-nil and is used to copy layer data to before
	// doing post processing. If this is non-nil and there's no post processing this is freed and
	// returned back to the pool.
	layerBuffersProcessed *VertexBuffers
	// If there is no clipping and children the child vertex data is stored here, unmodified by post processing.
	// When children are dirty this is updated which may trigger post processing even if not marked dirty.
	childrenBuffers *VertexQueue
	// If there is clipping and children the child vertex data is stored here, unmodified by post processing.
	childrenBuffersClipped *VertexBuffers
	// If there is post processing this is non-nil and is used to copy child data to before doing
	// post processing. If this is non-nil and there's no post processing this is freed and
	// returned back to the pool.
	childrenBuffersProcessed *VertexBuffers
	// All vertex buffers for layers and children. If there's any post processing this will have the processed versions.
	cachedBuffers *VertexQueue
}

// The UI this component belongs to.
func (b *Base) UI() *UI {
	return b.ui
}

// The current dirty state of this component.
func (b *Base) GetDirty() Dirty {
	return b.dirty
}

// Dirties this component and potentially dirties the the parent lineage.
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

// Sets the state of the component. If no state is given the default state is applied.
// If the component is not visually dirty the state change is checked against the layers
// and if any of the layers is shown/hidden based on the state change the component
// is marked visually dirty. If post processing exists for the new state the component
// is marked dirty for post processing.
func (b *Base) SetState(state State) {
	state.Remove(StateDefault)
	if state.None() {
		state.Add(StateDefault)
	}

	if state == b.States {
		return
	}

	dirty := DirtyNone

	if !b.dirty.Is(DirtyVisual) {
		for _, layer := range b.Layers {
			if layer.ForStates(state) != layer.ForStates(b.States) {
				dirty.Add(DirtyVisual)
				break
			}
		}
	}

	if b.ui.Theme.StatePostProcess[state] != nil {
		dirty.Add(DirtyPostProcess)
	}

	b.Dirty(dirty)
	b.States = state
}

// A handy function for editing the state of a component defined inline.
func (b *Base) Edit(editor func(*Base)) *Base {
	editor(b)
	return b
}

// Removes the given states from the component. See SetState.
// This operation has no effect if the state is unchanged.
func (b *Base) RemoveStates(state State) {
	b.SetState(b.States.WithRemove(state))
}

// Adds the given states to the component. See SetState.
// This operation has no effect if the state is unchanged.
func (b *Base) AddStates(state State) {
	b.SetState(b.States.WithAdd(state))
}

// Sets the placement of the component and marks the placement as dirty
// if the placement is different. All parents are notified they have a dirty child placement.
func (b *Base) SetPlacement(placement Placement) {
	if b.Placement != placement {
		b.Placement = placement
		b.Dirty(DirtyPlacement)
	}
}

// Sets the relative placement of the component and marks the placement as dirty
// if the placement is different. All parents are notified they have a dirty child placement.
func (b *Base) SetRelativePlacement(placement Placement) {
	if b.RelativePlacement != placement {
		b.RelativePlacement = placement
		b.Dirty(DirtyPlacement)
	}
}

// Sets the transform of the component and marks the post processing as dirty
// if the transform is different. All parents are notified they have a dirty child visual.
func (b *Base) SetTransform(transform Transform) {
	if transform.IsEffectivelyIdentity() {
		transform.Identity()
	}
	if b.Transform != transform {
		b.Transform = transform
		b.Dirty(DirtyPostProcess)
	}
}

// Sets the transparency of the component and marks the post processing as dirty
// if the transparency is different. All parents are notified they have a dirty child visual.
func (b *Base) SetTransparency(transparency float32) {
	if b.Transparency != transparency {
		b.Transparency = transparency
		b.Dirty(DirtyPostProcess)
	}
}

// Sets the color modification function that's applied to all vertices of this component
// and marks the post processing as dirty if the modification function appears to be different.
// All parents are notified they have a dirty child visual.
func (b *Base) SetColor(color color.Modify) {
	if !b.Color.Equals(color) {
		b.Color = color.GetEffective()
		b.Dirty(DirtyPostProcess)
	}
}

// Sets the clip placement of the children and marks the child placement as dirty
// if the placement is different. All parents are notified they have a dirty child placement.
func (b *Base) SetClip(clip Placement) {
	if b.Clip != clip {
		b.Clip = clip
		b.Dirty(DirtyChildPlacement)
	}
}

// Initializes the state of the component. Before this is done it is expected that the
// component has been set on a UI and/or has a parent component. This should only be
// called when the UI is first initialized or on a child component that is added or removed.
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
	if b.cachedBuffers == nil {
		b.cachedBuffers = BufferQueuePool.Get()
	}
	b.Dirty(DirtyPlacement)
	b.PlayEvent(AnimationEventShow)

	if b.Hooks.OnInit != nil {
		b.Hooks.OnInit(b)
	}
}

// Places this component within the given bounds. If the component does not have any
// dirty placements this may have no effect unless force is passed with true which will
// cause all components in the tree to be placed. The context is expected to have the
// size that matches the dimensions of the parent bounds passed.
func (b *Base) Place(ctx *RenderContext, parent Bounds, force bool) {
	if !force && !b.dirty.Is(DirtyChildPlacement|DirtyPlacement) {
		return
	}

	doPlacement := b.dirty.Removed(DirtyPlacement) || force
	if doPlacement {
		relativeBounds := b.RelativePlacement.GetBoundsIn(parent)
		newBounds := b.Placement.GetBoundsIn(relativeBounds)
		force = b.SetBounds(newBounds)
	}

	baseCtx := ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)

	doChildPlacement := b.dirty.Removed(DirtyChildPlacement) || force
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
	}

	if b.Hooks.OnPlace != nil {
		b.Hooks.OnPlace(b, parent, baseCtx, force)
	}
}

// Sets the bounds of the component. If the bounds are different then each layer is
// placed with the new bounds. If the component does not currently have a dirty placement
// it is added. The component is also marked visually dirty. Returns whether the bounds have changed.
func (b *Base) SetBounds(newBounds Bounds) bool {
	if newBounds != b.Bounds {
		b.Bounds = newBounds
		for i := range b.Layers {
			b.Layers[i].Place(b, b.Bounds)
		}
		b.Dirty(DirtyVisual)
		return true
	}
	return false
}

// Updates the component. This includes it's layers, children, animation, and hooks.
// Each can return a dirty state which is applied to the component. A child is only
// updated by its parent, never by its render parent. If the component is hidden
// and UI.UpdateHidden is false then this has no effect. A component that is playing
// a hide, show, or remove animation needs to be updated to complete the desired action.
func (b *Base) Update(update Update) {
	if b.IsHidden() && !b.ui.UpdateHidden {
		return
	}

	dirty := DirtyNone

	for i := range b.Layers {
		if b.Layers[i].ForStates(b.States) {
			dirty.Add(b.Layers[i].Update(b, update))
		}
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

// Marks the component's layout as dirty so its children are layed out next placement.
func (b *Base) Relayout() {
	b.Dirty(DirtyChildPlacement)
}

// Marks the component as visually dirty (layers, children, & post processing) so it's
// rendered on next render.
func (b *Base) Rerender() {
	b.Dirty(DirtyVisual | DirtyChildVisual | DirtyPostProcess)
}

// Returns whether a component could render based on its invisible, animation, and hidden states.
func (b *Base) CouldRender() bool {
	if b.Transparency == 1 && !b.Animation.IsAnimating() {
		return false
	}
	if b.IsHidden() {
		return false
	}
	return true
}

// Returns whether this component is invisible based on its transparency or color.
func (b *Base) IsInvisible() bool {
	return b.Transparency == 1 || b.Color.Modify(color.White).A == 0
}

// Returns whether a component has a post processing function that needs to be applied.
func (b *Base) HasPostProcess() bool {
	return b.ui.Theme.StatePostProcess[b.States] != nil ||
		b.Hooks.OnPostProcess != nil ||
		b.Hooks.OnPostProcessLayers != nil ||
		b.Hooks.OnPostProcessChildren != nil ||
		b.Animation.IsAnimating() ||
		b.Transparency > 0 ||
		b.Color != nil ||
		b.Transform.HasAffect()
}

// Renders a component and adds all buffers to the given queue. If the component is not
// visible this returns immediately. This will use cached buffers if possible, otherwise
// they may need to be regenerated. If this component has any post processing then the
// vertices applied to the queue by this component are cloned and post processing is
// applied to the clones to preserve the cached state.
func (b *Base) Render(ctx *RenderContext, queue *VertexQueue) {
	if !b.CouldRender() {
		return
	}

	shouldPostProcess := b.dirty.Is(DirtyPostProcess)
	useCache := b.dirty.Not(DirtyVisual | DirtyChildVisual)

	// If we don't have a dirty post process, layer, or child, add the cached buffers.
	if !shouldPostProcess && useCache {
		queue.Add(b.cachedBuffers)

		return
	}

	// We should do post process if its dirty or if a layer or child
	// is dirty and it appears we have post processing operations.
	hasPostProcess := b.HasPostProcess()
	freePostProcess := !(shouldPostProcess || hasPostProcess)
	willPostProcess := shouldPostProcess || (!useCache && hasPostProcess)

	// If the parent is not the render parent, the context should always be from the parent.
	if b.parent != b.renderParent {
		ctx = b.ComputeRenderContext()
	} else {
		ctx = ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)
	}

	// Only update layer buffers if they reported dirty
	if b.dirty.Removed(DirtyVisual) {
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
	}

	// If there are layers get the updated value...
	if b.layerBuffers != nil {
		// Clone the buffers if we will do post processing
		if willPostProcess {
			if b.layerBuffersProcessed == nil {
				b.layerBuffersProcessed = BufferPool.Get()
			}
			b.layerBuffers.CloneTo(b.layerBuffersProcessed)
		} else if freePostProcess && b.layerBuffersProcessed != nil {
			BufferPool.Free(b.layerBuffersProcessed)
			b.layerBuffersProcessed = nil
		}
	}

	// Only update children buffers if they reported dirty
	if b.dirty.Removed(DirtyChildVisual) {
		if b.childrenBuffers == nil {
			b.childrenBuffers = BufferQueuePool.Get()
		}

		b.childrenBuffers.Clear()

		if len(b.Children) > 0 {
			clipBounds := b.Clip.GetBoundsIn(b.Bounds)

			b.childrenBuffers.Clip(
				clipBounds,
				&b.childrenBuffersClipped,
				func(innerQueue *VertexQueue, clipping bool) {
					shown := b.ShownChildren()
					for _, child := range shown {
						beforeChild := innerQueue.Position()

						child.Render(ctx, innerQueue)

						if clipping && !child.visualBounds.Intersects(clipBounds) {
							innerQueue.Reset(beforeChild)
						}
					}
				},
			)
		}
	}

	// If there are children, get the updated value...
	if b.childrenBuffers != nil {
		// Clone the buffers if we will do post processing
		if willPostProcess {
			if b.childrenBuffersProcessed == nil {
				b.childrenBuffersProcessed = ClipPool.Get()
			}
			if b.childrenBuffersClipped != nil {
				b.childrenBuffersClipped.CloneTo(b.childrenBuffersProcessed)
			} else {
				b.childrenBuffers.ToBuffers().CloneTo(b.childrenBuffersProcessed)
			}
		} else if freePostProcess && b.childrenBuffersProcessed != nil {
			BufferPool.Free(b.childrenBuffersProcessed)
			b.childrenBuffersProcessed = nil
		}
	}

	// Rebuild the cached buffers
	b.cachedBuffers.Clear()
	if b.layerBuffersProcessed != nil {
		b.cachedBuffers.Add(b.layerBuffersProcessed)
	} else if b.layerBuffers != nil {
		b.cachedBuffers.Add(b.layerBuffers)
	}
	if b.childrenBuffersProcessed != nil {
		b.cachedBuffers.Add(b.childrenBuffersProcessed)
	} else if b.childrenBuffersClipped != nil {
		b.cachedBuffers.Add(b.childrenBuffersClipped)
	} else if b.childrenBuffers != nil {
		b.cachedBuffers.Add(b.childrenBuffers)
	}

	// If we still need to post process...
	if willPostProcess {
		b.Hooks.OnPostProcessLayers.Run(b, ctx, b.layerBuffersProcessed)
		b.Hooks.OnPostProcessChildren.Run(b, ctx, b.childrenBuffersProcessed)

		b.PostProcess(ctx, b.cachedBuffers.ToBuffers())
	}

	// Pass the cached buffers to the OnRender.
	if b.Hooks.OnRender != nil {
		b.Hooks.OnRender(b, ctx, b.cachedBuffers)
	}

	// Update bounds based on cachedBuffers.
	b.updateVisualBounds()

	// Finally queue cachedBuffers for rendering.
	queue.Add(b.cachedBuffers)
}

// Updates the visual bounds of this component based on all the vertices in the currently cached buffers.
// This may be used by parent components to clip out a component entirely from the rendering process.
func (b *Base) updateVisualBounds() {
	b.visualBounds.Clear()

	vertices := NewVertexIterator(b.cachedBuffers, true)
	for vertices.HasNext() {
		v := vertices.Next()
		b.visualBounds.Include(v.X, v.Y)
	}
}

// Performs the post processing that's defined on this component to the given vertex buffers.
func (b *Base) PostProcess(ctx *RenderContext, out *VertexBuffers) {
	b.postProcessState(ctx, out)
	b.postProcessAnimation(ctx, out)
	b.postProcessTransparency(ctx, out)
	b.postProcessColor(ctx, out)
	b.postProcessTransform(ctx, out)

	b.Hooks.OnPostProcess.Run(b, ctx, out)

	b.dirty.Remove(DirtyPostProcess)
}

// performs post processing based on the component's exact state.
func (b *Base) postProcessState(ctx *RenderContext, out *VertexBuffers) {
	modifier := b.ui.Theme.StatePostProcess[b.States]
	modifier.Run(b, ctx, out)
}

// performs post processing for the current animation.
func (b *Base) postProcessAnimation(ctx *RenderContext, out *VertexBuffers) {
	b.Animation.PostProcess(b, ctx, out)
}

// performs post processing if a non-zero transparency is defined.
func (b *Base) postProcessTransparency(ctx *RenderContext, out *VertexBuffers) {
	if alphaMultiplier := 1 - b.Transparency; alphaMultiplier < 1 {
		vertices := NewVertexIterator(out, true)
		for vertices.HasNext() {
			v := vertices.Next()
			v.InitColor()
			v.Color.A *= alphaMultiplier
		}
	}
}

// performs post processing if a shading technique is defined.
func (b *Base) postProcessColor(ctx *RenderContext, out *VertexBuffers) {
	if colorModifier := b.Color; colorModifier != nil {
		vertices := NewVertexIterator(out, true)
		for vertices.HasNext() {
			v := vertices.Next()
			v.InitColor()
			v.Color = colorModifier(v.Color)
		}
	}
}

// performs post processing if the transform is defined.
func (b *Base) postProcessTransform(ctx *RenderContext, out *VertexBuffers) {
	if transform := &b.Transform; transform.HasAffect() {
		vertices := NewVertexIterator(out, true)
		for vertices.HasNext() {
			v := vertices.Next()
			v.X, v.Y = transform.Transform(v.X, v.Y)
		}
	}
}

// Computes the render context for this component looking up through the parents
// to get the effective text styles and relative to the bounds of this component.
func (b *Base) ComputeRenderContext() *RenderContext {
	if b.parent == nil {
		return b.ui.renderContext.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)
	} else {
		return b.parent.ComputeRenderContext().WithBoundsAndTextStyles(b.Bounds, b.TextStyles)
	}
}

// Returns true if this component is the render parent to the given component.
func (b *Base) IsRenderParent(child *Base) bool {
	return child != nil && child.renderParent == b
}

// Returns the children that are rendered by this component.
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
func (b *Base) PreferredSize(ctx *RenderContext, maxWidth float32) gfx.Coord {
	baseCtx := ctx.WithBoundsAndTextStyles(b.Bounds, b.TextStyles)

	size := gfx.Coord{}
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
	size = size.Max(gfx.Coord{X: minSizeX, Y: minSizeY})
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

// Updates the components placement to be the desired width if possible for the given render context.
func (b *Base) PlaceWidthFor(width float32, ctx *RenderContext) {
	size := b.PreferredSize(ctx, width)
	b.SetPlacement(b.Placement.WithSize(size.X, size.Y))
}

// Updates the components placement to be the desired width if possible.
func (b *Base) PlaceWidth(width float32) {
	b.PlaceWidthFor(width, b.ComputeRenderContext())
}

// Updates the components placement to be the smallest width possible.
func (b *Base) PlaceMinWidth() {
	b.PlaceWidth(0)
}

// Updates the components placement to be the smallest height possible for the
// current width and the given render context.
func (b *Base) PlaceMinHeight() {
	b.PlaceWidth(b.Bounds.Width())
}

// The parent of this component, if any.
func (b *Base) Parent() *Base {
	if b == nil {
		return nil
	}
	return b.parent
}

// If the component is a child to this component.
func (b *Base) IsChild(c *Base) bool {
	return c.IsParent(b)
}

// If the component is a child to this component or is this component.
func (b *Base) IsChildOrSelf(c *Base) bool {
	return c == b || c.IsParent(b)
}

// If the component is a parent to this component.
func (b *Base) IsParent(p *Base) bool {
	if b.parent == p {
		return true
	}
	if b.parent == nil {
		return false
	}
	return b.parent.IsParent(p)
}

// If the component is a parent to this component or is this component.
func (b *Base) IsParentOrSelf(p *Base) bool {
	return b == p || b.IsParent(p)
}

// Returns true if the point is inside the Bounds of this component. The point is expected
// to be transformed to the relative orientation of this component. If an OverShape is
// defined then point-in-polygon logic is used to determine inside-ness.
func (b *Base) IsInside(pt gfx.Coord) bool {
	if !b.Bounds.InsideCoord(pt) {
		return false
	}

	if len(b.OverShape) > 0 {
		normalized := gfx.Coord{
			X: b.Bounds.Dx(pt.X),
			Y: b.Bounds.Dy(pt.Y),
		}
		if !InPolygon(b.OverShape, normalized) {
			return false
		}
	}

	return true
}

// Returns the component at the given point within this component.
// If a TransparencyThreshold is defined on the UI and this component is too
// transparent it is considered invisible to pointer events.
// If the UI is configured to transform the pointer to the relative orientation
// of this component then the coordinate is passed through the inverse transform
// of this component before further inspection is done.
// If the point is not within the Bounds or OverShape of this component then nil
// is returned. Otherwise all ShownChildren are inspected using the same logic
// as above and the first one to return a non-nil value is returned.
func (b *Base) At(pt gfx.Coord) *Base {
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

// Sets the visibility of this component to either show or hide. See Show & Hide.
func (b *Base) SetVisible(show bool) {
	if show {
		b.Show()
	} else {
		b.Hide()
	}
}

// Hides the component if it's not already hiding or hidden. If a hide animation
// is defined the animation will play out before the component is hidden and
// effectively removed from it's render parent's layout.
func (b *Base) Hide() {
	if !b.States.Is(StateHidden | StateHiding) {
		if b.PlayEvent(AnimationEventHide) {
			b.SetState(b.States.WithAdd(StateHiding).WithRemove(StateShowing))
		} else {
			b.hideNow()
		}
	}
}

// Hides the component now, fixing the states and updating the render parent.
func (b *Base) hideNow() {
	b.SetState(b.States.WithAdd(StateHidden).WithRemove(StateHiding | StateShowing))
	b.renderParent.ChildrenUpdated()
	if b.parent != b.renderParent {
		b.parent.ChildrenUpdated()
	}
}

// Shows the component if it's not already shown or showing. If a show
// animation is defined the animation will play out before the component
// is considered shown and can be interacted with. As soon as the showing
// process begins it's effectively added back to its render parent in the same
// position it previously was in it's layout.
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

// Shows the component now, fixing the states and updating the render parent.
func (b *Base) showNow() {
	b.RemoveStates(StateHidden | StateHiding | StateShowing)
	b.renderParent.ChildrenUpdated()
	if b.parent != b.renderParent {
		b.parent.ChildrenUpdated()
	}
}

// Removes all listeners (init, place, update, render, remove, post processes, events) from this component.
func (b *Base) ClearListeners() {
	b.Hooks.Clear()
	b.Events.Clear()
}

// Removes the component if it's not already removed or removing. If a
// remove animation is defined the animation will play out before the component
// is actually removed from the parent & render parent. On removal much of the cached
// state of the component is freed until the component is Init again by a new parent.
func (b *Base) Remove() {
	if b.parent != nil && !b.States.Is(StateRemoving) {
		if b.PlayEvent(AnimationEventRemove) {
			b.SetState(StateRemoving)
		} else {
			b.removeNow()
		}
	}
}

// Removes the component now cleaning up cached data and references.
func (b *Base) removeNow() {
	b.Hooks.OnRemove.Run(b)
	b.Hooks.OnRemove.Clear()
	if b.renderParent != nil {
		b.renderParent.ChildrenUpdated()
	}
	b.cleanup()
	b.removeFrom(b.parent)
	b.parent = nil
	b.renderParent = nil
}

// Removes the component from any render parents, returns all buffers & queues
// back to the pools, and does the same for all descendants.
func (b *Base) cleanup() {
	if b.parent != b.renderParent {
		b.removeFrom(b.renderParent)
	}
	if b.layerBuffers != nil {
		BufferPool.Free(b.layerBuffers)
		b.layerBuffers = nil
	}
	if b.layerBuffersProcessed != nil {
		BufferPool.Free(b.layerBuffersProcessed)
		b.layerBuffersProcessed = nil
	}
	if b.childrenBuffers != nil {
		BufferQueuePool.Free(b.childrenBuffers)
		b.childrenBuffers = nil
	}
	if b.childrenBuffersClipped != nil {
		ClipPool.Free(b.childrenBuffersClipped)
		b.childrenBuffersClipped = nil
	}
	if b.childrenBuffersProcessed != nil {
		ClipPool.Free(b.childrenBuffersProcessed)
		b.childrenBuffersProcessed = nil
	}
	if b.cachedBuffers != nil {
		BufferQueuePool.Free(b.cachedBuffers)
		b.cachedBuffers = nil
	}
	for _, child := range b.Children {
		if child.parent == b {
			child.cleanup()
		}
	}
}

// Removes the component from the given parent.
// If this component does not exist in the given parent then this has no effect.
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

// Marks the children as updated and they required placement and visual updates.
func (b *Base) ChildrenUpdated() {
	if b != nil {
		b.Dirty(DirtyChildPlacement | DirtyChildVisual)
		b.shownDirty = true
	}
}

// Adds the given children to this parent and initializes them. This will
// dirty this component and all parents.
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

// Changes the render parent of this component. You can have a component
// exist in one tree which decides its initial placement but actually
// be considered outside its parent and rendered under another. The
// ordering, final layout, and rendering is done by the render parent.
// You can pass nil to this to reset it back to the parent and remove
// from the current render parent.
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

// Returns index this component is in it's render parent's
// children. Any hidden components in the parent are not
// adjusted for to maintain consistent ordering.
func (b *Base) Order() int {
	if b.renderParent == nil {
		return -1
	}
	return util.SliceIndexOf(b.renderParent.Children, b)
}

// Sets the index this component is in it's render parent's
// children. Any hidden components in the parent are not
// adjusted for to maintain consistent ordering.
func (b *Base) SetOrder(order int) {
	currentOrder := b.Order()
	if currentOrder == -1 || order == currentOrder {
		return
	}
	util.SliceMove(b.renderParent.Children, currentOrder, order)
	b.renderParent.ChildrenUpdated()
}

// Sets this component to be the last one rendered in it's render parent.
func (b *Base) BringToFront() {
	if b.renderParent != nil {
		b.SetOrder(len(b.renderParent.Children) - 1)
	}
}

// Sets this component to be the first one rendered in it's render parent.
func (b *Base) SendToBack() {
	b.SetOrder(0)
}

// Gets the color if any for the themed color.
func (b *Base) GetColorable(t color.Themed) color.Able {
	colorable := b.Colors.Get(t)
	if colorable == nil {
		colorable = b.ui.Theme.Colors.Get(t)
	}
	return colorable
}

// Returns whether this component is focusable. It must be marked focusable
// and it cannot be in a state where it doesn't receive input events.
func (b Base) IsFocusable() bool {
	return b.Focusable && !b.IsDisabled()
}

// Returns whether this component is draggable. It must be marked draggable
// and it cannot be in a state where it doesn't receive input events.
func (b Base) IsDraggable() bool {
	return b.Draggable && !b.IsDisabled()
}

// Returns whether this component is droppable. It must be marked droppable
// and it cannot be in a state where it doesn't receive input events.
func (b Base) IsDroppable() bool {
	return b.Droppable && !b.IsDisabled()
}

// Returns whether this component is any disabled states.
func (b Base) IsDisabled() bool {
	return b.States.Is(StatesDisabled)
}

// Returns if this component is considered shown (not hidden). It may
// be in the process of showing or hiding and still be considered shown.
func (b Base) IsShown() bool {
	return b.States.Not(StateHidden)
}

// Returns if this component is considered hidden. It is only hidden
// when hiding is done and it's not trying to show.
func (b Base) IsHidden() bool {
	return b.States.Is(StateHidden)
}

// Returns if this component is the one currently being drug.
func (b *Base) IsDragging() bool {
	return b.ui.Dragging == b
}

// Sets the disabled state of this component. Any associated disabled
// or enabled animations will be played if there is a change in state.
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

// Handles a pointer event. This is only invoked when the pointer is
// over this component or one of it's children.
// Updates state, plays any necessary animations, and can update the cursor.
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

// Handles a key event. This is only invoked when this component or
// one of its children are focused.
func (b *Base) OnKey(ev *KeyEvent) {
	if b.IsDisabled() {
		return
	}

	b.Events.OnKey.Trigger(ev)
}

// Handles a focus event.
// Updates state and plays any necessary animations.
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

// Handles a blur event.
// Updates state and plays any necessary animations.
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

// Handles a drag or drop event.
// Updates state and plays any necessary animations.
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

// A template for a base which can be applied.
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
	OverShape                   []gfx.Coord
	Transparency                float32
	Color                       color.Modify
	Animations                  *Animations
	AnimationsMerge             bool
	Cursors                     Cursors
	CursorsMerge                bool
	Colors                      color.Colors
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

// Applies the template to this component, anywhere it appears an option
// is defined on the template but not the component.
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

func colorableNil(c color.Able) bool {
	return c == nil
}
