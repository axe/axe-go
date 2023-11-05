package ui

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/id"
)

type CursorEvent int

const (
	CursorEventHover CursorEvent = iota
	CursorEventDown
	CursorEventDrag
	CursorEventDisabled
)

type Cursors struct {
	ds.EnumMap[CursorEvent, id.Identifier]
}

func NewCursors(m map[CursorEvent]id.Identifier) Cursors {
	return Cursors{
		EnumMap: ds.NewEnumMap(m),
	}
}

func (c Cursors) OverCursor(b *Base, ignoreDragging bool) id.Identifier {
	if b.IsDisabled() {
		cursorDisabled := c.Get(CursorEventDisabled)
		if !cursorDisabled.Empty() {
			return cursorDisabled
		}
	}

	if b.ui.Dragging == b && !ignoreDragging {
		cursorDrag := c.Get(CursorEventDrag)
		if !cursorDrag.Empty() {
			return cursorDrag
		}
	}

	return c.Get(CursorEventHover)
}

func (c Cursors) HandlePointer(ev *PointerEvent, b *Base) {
	cursorDisabled := c.Get(CursorEventDisabled)
	cursorHover := c.Get(CursorEventHover)
	cursorDown := c.Get(CursorEventDown)
	cursorOver := c.OverCursor(b, false)

	switch ev.Type {
	case PointerEventEnter:
		ev.SetCursor(cursorOver, b, false)
	case PointerEventMove:
		ev.SetCursor(cursorOver, b, false)
	case PointerEventDown:
		ev.SetCursor(cursorDown, b, false)
	case PointerEventLeave:
		ev.RemoveCursor(cursorHover)
		ev.RemoveCursor(cursorDisabled)
	case PointerEventUp:
		ev.RemoveCursor(cursorDown)
		ev.SetCursor(cursorOver, b, false)
	}
}

func (c Cursors) HandleDrag(ev *DragEvent, b *Base) {
	dragCursor := c.Get(CursorEventDrag)
	overCursor := c.OverCursor(b, true)

	switch ev.Type {
	case DragEventStart:
		ev.SetCursor(dragCursor, b, false)
	case DragEventEnd:
		ev.RemoveCursor(dragCursor)
		ev.SetCursor(overCursor, b, false)
	case DragEventCancel:
		ev.RemoveCursor(dragCursor)
		ev.SetCursor(overCursor, b, false)
	}
}

type HasCursor struct {
	Cursor        *id.Identifier
	CursorSet     bool
	CursorTrigger *Base
}

func (ev *HasCursor) RemoveCursor(cursor id.Identifier) {
	if *ev.Cursor == 0 || *ev.Cursor == cursor {
		*ev.Cursor = 0
	}
}

func (ev *HasCursor) SetCursor(cursor id.Identifier, trigger *Base, force bool) {
	if cursor.Empty() {
		return
	}
	if (ev.CursorSet == true && trigger != ev.CursorTrigger) && !force {
		return
	}
	*ev.Cursor = cursor
	ev.CursorSet = true
	ev.CursorTrigger = trigger
}

func cursorNil(id id.Identifier) bool {
	return id.Empty()
}
