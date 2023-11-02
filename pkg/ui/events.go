package ui

import (
	"time"
)

type Events struct {
	OnPointer func(ev *PointerEvent)
	OnKey     func(ev *KeyEvent)
	OnFocus   func(ev *Event)
	OnBlur    func(ev *Event)
	OnDrag    func(ev *DragEvent)
}

type Event struct {
	Time    time.Time
	Stop    bool
	Capture bool
	Cancel  bool
	Target  Component
}

func (e Event) withTarget(target Component) Event {
	e.Target = target
	return e
}

type DragEventType string

const (
	DragEventStart  DragEventType = "start"
	DragEventMove   DragEventType = "move"
	DragEventEnd    DragEventType = "end"
	DragEventCancel DragEventType = "cancel"
	DragEventOver   DragEventType = "over"
	DragEventDrop   DragEventType = "drop"
)

type DragEvent struct {
	Event
	Point      Coord
	Start      Coord
	DeltaStart Coord
	DeltaMove  Coord
	Type       DragEventType
	Dragging   Component
}

func (ev DragEvent) as(dragType DragEventType) *DragEvent {
	ev.Event = Event{Capture: false}
	ev.Type = dragType
	return &ev
}

type PointerEventType string

const (
	PointerEventDown  PointerEventType = "down"
	PointerEventUp    PointerEventType = "up"
	PointerEventLeave PointerEventType = "leave"
	PointerEventEnter PointerEventType = "enter"
	PointerEventWheel PointerEventType = "wheel"
	PointerEventMove  PointerEventType = "move"
)

type PointerEvent struct {
	Event
	Point  Coord
	Button int
	Amount int
	Type   PointerEventType
}

func (ev PointerEvent) as(eventType PointerEventType) PointerEvent {
	ev.Event = Event{Capture: true}
	ev.Type = eventType
	return ev
}

func (e PointerEvent) withTarget(target Component) PointerEvent {
	e.Target = target
	return e
}

type PointerButtons struct {
	Down     bool
	DownTime int64
}

type KeyEventType string

const (
	KeyEventDown  KeyEventType = "down"
	KeyEventUp    KeyEventType = "up"
	KeyEventPress KeyEventType = "press"
)

type KeyEvent struct {
	Event
	Key  string
	Char rune
	Type KeyEventType
}
