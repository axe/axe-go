package ui

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/id"
)

var Area = id.NewArea[uint32, uint16](
	id.WithCapacity(1024),
	id.WithResizeBuffer(24),
)

type UI struct {
	PointerButtons []PointerButtons
	PointerPoint   Coord
	Root           *Base
	PointerOver    *Base
	Focused        *Base
	Dragging       *Base
	DragStart      Coord
	DragCancels    ds.Set[string]
	Theme          *Theme
	Cursor         id.Identifier
	Named          id.DenseMap[*Base, uint16, uint16]

	TransformPointer      bool
	TransparencyThreshold float32
	UpdateHidden          bool
	PointerEntersOnDrag   bool
	PointerMovesOnDrag    bool

	amountContext AmountContext
	renderContext RenderContext
	bounds        Bounds
}

func NewUI() *UI {
	return &UI{
		Theme: &Theme{
			StatePostProcess: make(map[State]PostProcess),
			Fonts: id.NewDenseMap[*Font, uint16, uint8](
				id.WithArea(Area),
				id.WithCapacity(16),
			),
			Cursors: id.NewDenseMap[ExtentTile, uint16, uint8](
				id.WithArea(Area),
				id.WithCapacity(16),
			),
			Colors: NewColors(map[ThemeColor]Colorable{}),
			TextStyles: TextStyles{
				ParagraphStyles: ParagraphStyles{
					LineVerticalAlignment: AlignmentBottom,
					Wrap:                  TextWrapWord,
				},
				ParagraphsStyles: ParagraphsStyles{
					ClipShowX: ClipShowLeft,
					ClipShowY: ClipShowTop,
				},
				Color:    ColorBlack,
				FontSize: Amount{Value: 16},
			},
			Animations: Animations{
				ForEvent: ds.EnumMap[AnimationEvent, AnimationFactory]{},
				Named:    id.NewDenseKeyMap[AnimationFactory, uint16, uint8](),
			},
		},
		PointerButtons: make([]PointerButtons, 3),
		PointerPoint:   Coord{X: -1, Y: -1},
		Named: id.NewDenseMap[*Base, uint16, uint16](
			id.WithArea(Area),
			id.WithCapacity(128),
		),
	}
}

func (ui *UI) Init() {
	ui.Root.ui = ui
	ui.Root.Init()
}

func (ui *UI) Place(newBounds Bounds) {
	force := ui.bounds != newBounds
	if force || ui.Root.GetDirty().Is(DirtyPlacement|DirtyChildPlacement) {
		ui.bounds = newBounds
		ui.Root.Place(&ui.renderContext, newBounds, force)
	}
}

func (ui *UI) SetContext(ctx *AmountContext) {
	if ctx != nil && ui.amountContext != *ctx {
		ui.Root.Dirty(DirtyVisual | DirtyChildVisual)
		ui.amountContext = *ctx
		ui.renderContext = RenderContext{
			AmountContext: ui.amountContext.ForFont(ui.Theme.TextStyles.FontSize),
			Theme:         ui.Theme,
			TextStyles:    &ui.Theme.TextStyles,
		}
	}
}

func (ui *UI) RenderContext() *RenderContext {
	return &ui.renderContext
}

func (ui *UI) NeedsRender() bool {
	return ui.Root.GetDirty().Is(DirtyVisual | DirtyChildVisual | DirtyPostProcess)
}

func (ui *UI) Update(update Update) {
	ui.Root.Update(update)
}

func (ui *UI) Render(queue *VertexQueue) {
	for !ClipMemory.IsEmpty() {
		ClipPool.Free(ClipMemory.Pop())
	}
	ui.Root.Render(&ui.renderContext, queue)
}

func (ui *UI) IsPointerOver() bool {
	return ui.PointerPoint.X != -1 && ui.PointerPoint.Y != -1
}

func (ui *UI) GetCursor() (cursorVertex []Vertex) {
	if !ui.IsPointerOver() {
		return
	}
	cursorName := ui.Cursor
	if cursorName.Empty() {
		cursorName = ui.Theme.DefaultCursor
	}
	if !cursorName.Empty() {
		cursor := ui.Theme.Cursors.Get(cursorName)
		if cursor.Texture != "" {
			e := cursor.Extent
			p := ui.PointerPoint
			cursorVertex = []Vertex{
				{X: e.Left + p.X, Y: e.Top + p.Y, Tex: cursor.Coord(0, 0), HasCoord: true},
				{X: e.Right + p.X, Y: e.Top + p.Y, Tex: cursor.Coord(1, 0), HasCoord: true},
				{X: e.Right + p.X, Y: e.Bottom + p.Y, Tex: cursor.Coord(1, 1), HasCoord: true},
				{X: e.Left + p.X, Y: e.Bottom + p.Y, Tex: cursor.Coord(0, 1), HasCoord: true},
			}
		}
	}
	return cursorVertex
}

func (ui *UI) ProcessKeyEvent(ev KeyEvent) error {
	// If a focused component exists send the key event
	if ui.Focused != nil {
		triggerKeyEvent(getPath(ui.Focused), ev)
	}

	// If dragging and the key event cancels dragging...
	if ui.Dragging != nil && ui.DragCancels.Has(ev.Key) {
		pointer := PointerEvent{
			Event: ev.Event,
			Point: ui.PointerPoint,
			Type:  PointerEventLeave,
		}
		pointer.Target = ui.Focused
		ui.Dragging.OnDrag(ui.dragEvent(pointer, DragEventCancel))
		ui.Dragging = nil
	}

	return nil
}

func (ui *UI) ProcessPointerEvent(ev PointerEvent) error {
	if ui.Root == nil {
		return nil
	}

	// Set to current cursor
	ev.HasCursor = &HasCursor{Cursor: &ui.Cursor}

	// Cached drag event with accurate deltas
	dragEvent := ui.dragEvent(ev, DragEventCancel)

	// If leave event
	if ev.Type == PointerEventLeave {
		// For every component in ui.MouseOver trigger leave
		if ui.PointerOver != nil {
			triggerPointerEvent(getPath(ui.PointerOver), ev.withTarget(ui.PointerOver))
			ui.PointerOver = nil
		}

		// For the component that's being dragged, cancel it
		if ui.Dragging != nil {
			ui.triggerDrag(dragEvent.as(DragEventCancel), true)
		}

		ui.PointerPoint.X = -1
		ui.PointerPoint.Y = -1
		return nil
	}

	// Handle mouse moving & enter/leave/move events
	if !ui.PointerPoint.Equals(ev.Point) {
		over := ui.Root.At(ev.Point)
		ev.Target = over

		if ui.Dragging != nil {
			// Trigger move event
			dragMove := ui.triggerDrag(dragEvent.as(DragEventMove), false)
			// If cancel requested, stop dragging
			if dragMove.Cancel {
				ui.Dragging = nil
			} else if over != nil {
				dragOver := dragEvent.as(DragEventOver)
				triggerDragEvent(getDroppablePath(over), dragOver)
				// If cancel requested, stop dragging
				if dragOver.Cancel {
					ui.triggerDrag(dragEvent.as(DragEventCancel), true)
				}
			}
		}

		ui.PointerPoint = ev.Point

		if over != ui.PointerOver {
			oldOver := NewComponentMap()
			if ui.PointerOver != nil {
				oldOver.AddLineage(ui.PointerOver)
			}
			newOver := NewComponentMap()
			if over != nil {
				newOver.AddLineage(over)
			}

			leavePath, movePath, enterPath := oldOver.Compare(newOver)

			// For every component in ui.MouseOver not in over we need to trigger leave
			triggerPointerEvent(leavePath, ev.as(PointerEventLeave))

			// Only trigger enter & move if we're not dragging
			if ui.Dragging == nil || ui.PointerEntersOnDrag {
				// For every component in over not in ui.MouseOver we need to trigger enter
				triggerPointerEvent(enterPath, ev.as(PointerEventEnter))
			}

			if ui.Dragging == nil || ui.PointerMovesOnDrag {
				// For every component in both and ev.Type = move we need to trigger move
				triggerPointerEvent(movePath, ev.as(PointerEventMove))
			}

			ui.PointerOver = over
		} else if over != nil && (ui.Dragging == nil || ui.PointerMovesOnDrag) {
			// For every component in over's lineage we need to trigger move if we're not dragging
			triggerPointerEvent(getPath(over), ev.as(PointerEventMove))
		}
	}

	// Handle down/up/wheel event
	if (ev.Type == PointerEventDown || ev.Type == PointerEventUp || ev.Type == PointerEventWheel) && ui.PointerOver != nil && ui.Dragging == nil {
		triggerPointerEvent(getPath(ui.PointerOver), ev.withTarget(ui.PointerOver))
	}

	// Handle drag end/drop
	if ui.Dragging != nil && ev.Type == PointerEventUp {
		if ui.PointerOver != nil {
			triggerDragEvent(getDroppablePath(ui.PointerOver), dragEvent.as(DragEventDrop))
		}
		ui.triggerDrag(dragEvent.as(DragEventEnd), true)
	}

	// Change focus on down if not dragging
	if ev.Type == PointerEventDown && ui.Dragging == nil {
		if ui.Focused != ui.PointerOver {
			old := getFocusablePath(ui.Focused)
			new := getFocusablePath(ui.PointerOver)

			oldMap := NewComponentMap()
			oldMap.AddMany(old)
			newMap := NewComponentMap()
			newMap.AddMany(new)

			ui.Focused = ui.PointerOver
			ev := ev.Event.withTarget(ui.Focused)

			inOld, _, inNew := oldMap.Compare(newMap)
			triggerBlurEvent(inOld, ev)
			triggerFocusEvent(inNew, ev)
		}

		// If down and it's draggable, start the drag
		if ui.PointerOver != nil && ui.Dragging == nil {
			draggablePath := getDraggablePath(ui.PointerOver)
			if len(draggablePath) > 0 {
				dragging := draggablePath[0]
				dragStart := dragEvent.as(DragEventStart)
				dragging.OnDrag(dragStart)
				// If not stopped, consider the start accepted
				if !dragStart.Cancel {
					ui.Dragging = dragging
					ui.DragStart = ev.Point
				}
			}
		}
	}

	return nil
}

func (ui UI) dragEvent(ev PointerEvent, dragType DragEventType) *DragEvent {
	return &DragEvent{
		Event:    ev.Event.withTarget(ui.PointerOver),
		Point:    ev.Point,
		Start:    ui.DragStart,
		Type:     dragType,
		Dragging: ui.Dragging,
		DeltaStart: Coord{
			X: ev.Point.X - ui.DragStart.X,
			Y: ev.Point.Y - ui.DragStart.Y,
		},
		DeltaMove: Coord{
			X: ev.Point.X - ui.PointerPoint.X,
			Y: ev.Point.Y - ui.PointerPoint.Y,
		},
		HasCursor: ev.HasCursor,
	}
}

func (ui *UI) triggerDrag(ev *DragEvent, clear bool) *DragEvent {
	ui.Dragging.OnDrag(ev)
	if clear {
		ui.Dragging = nil
	}
	return ev
}

func getPath(c *Base) []*Base {
	return getPathWhere(c, nil)
}

func getFocusablePath(c *Base) []*Base {
	return getPathWhere(c, func(c *Base) bool {
		return c.IsFocusable()
	})
}

func getDroppablePath(c *Base) []*Base {
	return getPathWhere(c, func(c *Base) bool {
		return c.IsDroppable()
	})
}

func getDraggablePath(c *Base) []*Base {
	return getPathWhere(c, func(c *Base) bool {
		return c.IsDraggable()
	})
}

func getPathWhere(c *Base, where func(*Base) bool) []*Base {
	path := make([]*Base, 0, 8)
	curr := c
	for curr != nil {
		if where == nil || where(curr) {
			path = append(path, curr)
		}

		curr = curr.Parent()
	}
	return path
}

func triggerPointerEvent(path []*Base, ev PointerEvent) {
	triggerEvent(path, &ev.Event, func(c *Base) {
		c.OnPointer(&ev)
	})
}

func triggerKeyEvent(path []*Base, ev KeyEvent) {
	triggerEvent(path, &ev.Event, func(c *Base) {
		c.OnKey(&ev)
	})
}

func triggerFocusEvent(path []*Base, ev Event) {
	triggerEvent(path, &ev, func(c *Base) {
		c.OnFocus(&ev)
	})
}

func triggerBlurEvent(path []*Base, ev Event) {
	triggerEvent(path, &ev, func(c *Base) {
		c.OnBlur(&ev)
	})
}

func triggerDragEvent(path []*Base, ev *DragEvent) {
	triggerEvent(path, &ev.Event, func(c *Base) {
		c.OnDrag(ev)
	})
}

func triggerEvent(path []*Base, ev *Event, trigger func(*Base)) {
	ev.Capture = true
	ev.Stop = false
	for i := len(path) - 1; i >= 0; i-- {
		trigger(path[i])
		if ev.Stop {
			return
		}
	}
	ev.Capture = false
	for i := 0; i < len(path); i++ {
		trigger(path[i])
		if ev.Stop {
			return
		}
	}
}
