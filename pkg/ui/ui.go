package ui

import "github.com/axe/axe-go/pkg/ds"

type Component interface {
	Init(init Init)
	Place(parent Bounds)
	Update(update Update)
	Render(ctx AmountContext, out *UIVertexBuffer)
	Parent() Component
	At(pt Coord) Component

	IsFocusable() bool
	IsDraggable() bool
	IsDroppable() bool

	OnPointer(ev *PointerEvent)
	OnKey(ev *KeyEvent)
	OnFocus(ev *ComponentEvent)
	OnBlur(ev *ComponentEvent)
	OnDrag(ev *DragEvent)
}

type UI struct {
	PointerButtons []PointerButtons
	PointerPoint   Coord
	Root           Component
	PointerOver    Component
	Focused        Component
	Dragging       Component
	DragStart      Coord
	DragCancels    ds.Set[string]
	Theme          *Theme
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
		ui.Dragging.OnDrag(ui.dragEvent(pointer, DragEventCancel))
		ui.Dragging = nil
	}

	return nil
}

func (ui *UI) ProcessPointerEvent(ev PointerEvent) error {
	if ui.Root == nil {
		return nil
	}

	// Cached drag event with accurate deltas
	dragEvent := ui.dragEvent(ev, DragEventCancel)

	// If leave event
	if ev.Type == PointerEventLeave {
		// For every component in ui.MouseOver trigger leave
		if ui.PointerOver != nil {
			triggerPointerEvent(getPath(ui.PointerOver), ev)
			ui.PointerOver = nil
		}

		// For the component that's being dragged, cancel it
		if ui.Dragging != nil {
			ui.Dragging.OnDrag(dragEvent.as(DragEventCancel))
			ui.Dragging = nil
		}
		return nil
	}

	// Handle mouse moving & enter/leave/move events
	if !ui.PointerPoint.Equals(ev.Point) {
		over := ui.Root.At(ev.Point)

		if ui.Dragging != nil {
			// Trigger move event
			dragMove := dragEvent.as(DragEventMove)
			ui.Dragging.OnDrag(dragMove)
			// If cancel requested, stop dragging
			if dragMove.Cancel {
				ui.Dragging = nil
			} else if over != nil {
				dragOver := dragEvent.as(DragEventOver)
				triggerDragEvent(getDroppablePath(over), dragOver)
				// If cancel requested, stop dragging
				if dragOver.Cancel {
					ui.Dragging.OnDrag(ui.dragEvent(ev, DragEventCancel))
					ui.Dragging = nil
				}
			}
		}

		ui.PointerPoint = ev.Point

		if over != ui.PointerOver {
			oldOver := ComponentMap{}
			if ui.PointerOver != nil {
				oldOver.AddLineage(ui.PointerOver)
			}
			newOver := ComponentMap{}
			if over != nil {
				newOver.AddLineage(over)
			}

			leavePath, movePath, enterPath := oldOver.Compare(newOver)

			// For every component in ui.MouseOver not in over we need to trigger leave
			triggerPointerEvent(leavePath, ev.as(PointerEventLeave))

			// For every component in over not in ui.MouseOver we need to trigger enter
			triggerPointerEvent(enterPath, ev.as(PointerEventEnter))

			// For every component in both and ev.Type = move we need to trigger move
			triggerPointerEvent(movePath, ev.as(PointerEventMove))

			ui.PointerOver = over
		}
	}

	// Handle down/up/wheel event
	if (ev.Type == PointerEventDown || ev.Type == PointerEventUp || ev.Type == PointerEventWheel) && ui.PointerOver != nil {
		triggerPointerEvent(getPath(ui.PointerOver), ev)
	}

	// Handle drag end/drop
	if ui.Dragging != nil && ev.Type == PointerEventUp {
		if ui.PointerOver != nil {
			triggerDragEvent(getDroppablePath(ui.PointerOver), dragEvent.as(DragEventDrop))
		}
		ui.Dragging.OnDrag(dragEvent.as(DragEventEnd))
		ui.Dragging = nil
	}

	// Change focus on down
	if ev.Type == PointerEventDown {
		if ui.Focused != ui.PointerOver {
			old := getFocusablePath(ui.Focused)
			new := getFocusablePath(ui.PointerOver)

			oldMap := ComponentMap{}
			oldMap.AddMany(old)
			newMap := ComponentMap{}
			newMap.AddMany(new)

			ui.Focused = ui.PointerOver
			ev := ComponentEvent{Target: ui.Focused}

			inOld, _, inNew := oldMap.Compare(newMap)
			triggerBlurEvent(inOld, ev)
			triggerFocusEvent(inNew, ev)
		}

		// If down and it's draggable, start the drag
		if ui.PointerOver != nil && ui.PointerOver.IsDraggable() && ui.Dragging == nil {
			dragStart := dragEvent.as(DragEventStart)
			ui.PointerOver.OnDrag(dragStart)
			// If not stopped, consider the start accepted
			if !dragStart.Cancel {
				ui.Dragging = ui.PointerOver
				ui.DragStart = ev.Point
			}
		}
	}

	return nil
}

func (ui UI) dragEvent(ev PointerEvent, dragType DragEventType) *DragEvent {
	return &DragEvent{
		Event:    ev.Event,
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
	}
}

func getPath(c Component) []Component {
	return getPathWhere(c, nil)
}

func getFocusablePath(c Component) []Component {
	return getPathWhere(c, func(c Component) bool {
		return c.IsFocusable()
	})
}

func getDroppablePath(c Component) []Component {
	return getPathWhere(c, func(c Component) bool {
		return c.IsDroppable()
	})
}

func getPathWhere(c Component, where func(Component) bool) []Component {
	path := make([]Component, 0)
	curr := c
	for curr != nil {
		if where == nil || where(curr) {
			path = append(path, curr)
		}

		curr = curr.Parent()
	}
	return path
}

func triggerPointerEvent(path []Component, ev PointerEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnPointer(&ev)
	})
}

func triggerKeyEvent(path []Component, ev KeyEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnKey(&ev)
	})
}

func triggerFocusEvent(path []Component, ev ComponentEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnFocus(&ev)
	})
}

func triggerBlurEvent(path []Component, ev ComponentEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnBlur(&ev)
	})
}

func triggerDragEvent(path []Component, ev *DragEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnDrag(ev)
	})
}

func triggerEvent(path []Component, ev *Event, trigger func(Component)) {
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
